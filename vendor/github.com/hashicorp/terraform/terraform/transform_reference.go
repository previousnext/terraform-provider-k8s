package terraform

import (
	"fmt"
	"log"
	"sort"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/configs/configschema"
	"github.com/hashicorp/terraform/dag"
	"github.com/hashicorp/terraform/lang"
)

// GraphNodeReferenceable must be implemented by any node that represents
// a Terraform thing that can be referenced (resource, module, etc.).
//
// Even if the thing has no name, this should return an empty list. By
// implementing this and returning a non-nil result, you say that this CAN
// be referenced and other methods of referencing may still be possible (such
// as by path!)
type GraphNodeReferenceable interface {
	GraphNodeModulePath

	// ReferenceableAddrs returns a list of addresses through which this can be
	// referenced.
	ReferenceableAddrs() []addrs.Referenceable
}

// GraphNodeReferencer must be implemented by nodes that reference other
// Terraform items and therefore depend on them.
type GraphNodeReferencer interface {
	GraphNodeModulePath

	// References returns a list of references made by this node, which
	// include both a referenced address and source location information for
	// the reference.
	References() []*addrs.Reference
}

type GraphNodeAttachDependencies interface {
	GraphNodeConfigResource
	AttachDependencies([]addrs.ConfigResource)
}

// graphNodeAttachResourceDependencies records all resources that are transitively
// referenced through depends_on in the configuration. This is used by data
// resources to determine if they can be read during the plan, or if they need
// to be further delayed until apply.
// We can only use an addrs.ConfigResource address here, because modules are
// not yet expended in the graph. While this will cause some extra data
// resources to show in the plan when their depends_on references may be in
// unrelated module instances, the fact that it only happens when there are any
// resource updates pending means we can still avoid the problem of the
// "perpetual diff"
type graphNodeAttachResourceDependencies interface {
	GraphNodeConfigResource
	AttachResourceDependencies([]addrs.ConfigResource)
	DependsOn() []*addrs.Reference
}

// GraphNodeReferenceOutside is an interface that can optionally be implemented.
// A node that implements it can specify that its own referenceable addresses
// and/or the addresses it references are in a different module than the
// node itself.
//
// Any referenceable addresses returned by ReferenceableAddrs are interpreted
// relative to the returned selfPath.
//
// Any references returned by References are interpreted relative to the
// returned referencePath.
//
// It is valid but not required for either of these paths to match what is
// returned by method Path, though if both match the main Path then there
// is no reason to implement this method.
//
// The primary use-case for this is the nodes representing module input
// variables, since their expressions are resolved in terms of their calling
// module, but they are still referenced from their own module.
type GraphNodeReferenceOutside interface {
	// ReferenceOutside returns a path in which any references from this node
	// are resolved.
	ReferenceOutside() (selfPath, referencePath addrs.Module)
}

// ReferenceTransformer is a GraphTransformer that connects all the
// nodes that reference each other in order to form the proper ordering.
type ReferenceTransformer struct{}

func (t *ReferenceTransformer) Transform(g *Graph) error {
	// Build a reference map so we can efficiently look up the references
	vs := g.Vertices()
	m := NewReferenceMap(vs)

	// Find the things that reference things and connect them
	for _, v := range vs {
		if _, ok := v.(GraphNodeDestroyer); ok {
			// destroy nodes references are not connected, since they can only
			// use their own state.
			continue
		}
		parents := m.References(v)
		parentsDbg := make([]string, len(parents))
		for i, v := range parents {
			parentsDbg[i] = dag.VertexName(v)
		}
		log.Printf(
			"[DEBUG] ReferenceTransformer: %q references: %v",
			dag.VertexName(v), parentsDbg)

		for _, parent := range parents {
			g.Connect(dag.BasicEdge(v, parent))
		}

		if len(parents) > 0 {
			continue
		}
	}

	return nil
}

type depMap map[string]addrs.ConfigResource

// addDep adds the vertex if it represents a resource in the
// graph.
func (m depMap) add(v dag.Vertex) {
	// we're only concerned with resources which may have changes that
	// need to be applied.
	switch v := v.(type) {
	case GraphNodeResourceInstance:
		instAddr := v.ResourceInstanceAddr()
		addr := instAddr.ContainingResource().Config()
		m[addr.String()] = addr
	case GraphNodeConfigResource:
		addr := v.ResourceAddr()
		m[addr.String()] = addr
	}
}

// attachDataResourceDependenciesTransformer records all resources transitively referenced
// through a configuration depends_on.
type attachDataResourceDependenciesTransformer struct {
}

func (t attachDataResourceDependenciesTransformer) Transform(g *Graph) error {
	// First we need to make a map of referenceable addresses to their vertices.
	// This is very similar to what's done in ReferenceTransformer, but we keep
	// implementation separate as they may need to change independently.
	vertices := g.Vertices()
	refMap := NewReferenceMap(vertices)

	for _, v := range vertices {
		depender, ok := v.(graphNodeAttachResourceDependencies)
		if !ok {
			continue
		}
		selfAddr := depender.ResourceAddr()

		// Only data need to attach depends_on, so they can determine if they
		// are eligible to be read during plan.
		if selfAddr.Resource.Mode != addrs.DataResourceMode {
			continue
		}

		// depMap will only add resource references and dedupe
		m := make(depMap)

		for _, dep := range refMap.DependsOn(v) {
			// any the dependency
			m.add(dep)
			// and check any ancestors
			ans, _ := g.Ancestors(dep)
			for _, v := range ans {
				m.add(v)
			}
		}

		deps := make([]addrs.ConfigResource, 0, len(m))
		for _, d := range m {
			deps = append(deps, d)
		}

		log.Printf("[TRACE] AttachDependsOnTransformer: %s depends on %s", depender.ResourceAddr(), deps)
		depender.AttachResourceDependencies(deps)
	}

	return nil
}

// AttachDependenciesTransformer records all resource dependencies for each
// instance, and attaches the addresses to the node itself. Managed resource
// will record these in the state for proper ordering of destroy operations.
type AttachDependenciesTransformer struct {
}

func (t AttachDependenciesTransformer) Transform(g *Graph) error {
	for _, v := range g.Vertices() {
		attacher, ok := v.(GraphNodeAttachDependencies)
		if !ok {
			continue
		}
		selfAddr := attacher.ResourceAddr()

		// Data sources don't need to track destroy dependencies
		if selfAddr.Resource.Mode == addrs.DataResourceMode {
			continue
		}

		ans, err := g.Ancestors(v)
		if err != nil {
			return err
		}

		// dedupe addrs when there's multiple instances involved, or
		// multiple paths in the un-reduced graph
		depMap := map[string]addrs.ConfigResource{}
		for _, d := range ans {
			var addr addrs.ConfigResource

			switch d := d.(type) {
			case GraphNodeResourceInstance:
				instAddr := d.ResourceInstanceAddr()
				addr = instAddr.ContainingResource().Config()
			case GraphNodeConfigResource:
				addr = d.ResourceAddr()
			default:
				continue
			}

			// Data sources don't need to track destroy dependencies
			if addr.Resource.Mode == addrs.DataResourceMode {
				continue
			}

			if addr.Equal(selfAddr) {
				continue
			}
			depMap[addr.String()] = addr
		}

		deps := make([]addrs.ConfigResource, 0, len(depMap))
		for _, d := range depMap {
			deps = append(deps, d)
		}
		sort.Slice(deps, func(i, j int) bool {
			return deps[i].String() < deps[j].String()
		})

		log.Printf("[TRACE] AttachDependenciesTransformer: %s depends on %s", attacher.ResourceAddr(), deps)
		attacher.AttachDependencies(deps)
	}

	return nil
}

// ReferenceMap is a structure that can be used to efficiently check
// for references on a graph, mapping internal reference keys (as produced by
// the mapKey method) to one or more vertices that are identified by each key.
type ReferenceMap map[string][]dag.Vertex

// References returns the set of vertices that the given vertex refers to,
// and any referenced addresses that do not have corresponding vertices.
func (m ReferenceMap) References(v dag.Vertex) []dag.Vertex {
	rn, ok := v.(GraphNodeReferencer)
	if !ok {
		return nil
	}

	var matches []dag.Vertex

	for _, ref := range rn.References() {
		subject := ref.Subject

		key := m.referenceMapKey(v, subject)
		if _, exists := m[key]; !exists {
			// If what we were looking for was a ResourceInstance then we
			// might be in a resource-oriented graph rather than an
			// instance-oriented graph, and so we'll see if we have the
			// resource itself instead.
			switch ri := subject.(type) {
			case addrs.ResourceInstance:
				subject = ri.ContainingResource()
			case addrs.ResourceInstancePhase:
				subject = ri.ContainingResource()
			case addrs.AbsModuleCallOutput:
				subject = ri.ModuleCallOutput()
			default:
				log.Printf("[WARN] ReferenceTransformer: reference not found: %q", subject)
				continue
			}
			key = m.referenceMapKey(v, subject)
		}
		vertices := m[key]
		for _, rv := range vertices {
			// don't include self-references
			if rv == v {
				continue
			}
			matches = append(matches, rv)
		}
	}

	return matches
}

// DependsOn returns the set of vertices that the given vertex refers to from
// the configured depends_on.
func (m ReferenceMap) DependsOn(v dag.Vertex) []dag.Vertex {
	depender, ok := v.(graphNodeAttachResourceDependencies)
	if !ok {
		return nil
	}

	var matches []dag.Vertex

	for _, ref := range depender.DependsOn() {
		subject := ref.Subject

		key := m.referenceMapKey(v, subject)
		vertices, ok := m[key]
		if !ok {
			log.Printf("[WARN] DependOn: reference not found: %q", subject)
			continue
		}
		for _, rv := range vertices {
			// don't include self-references
			if rv == v {
				continue
			}
			matches = append(matches, rv)
		}
	}

	return matches
}

func (m *ReferenceMap) mapKey(path addrs.Module, addr addrs.Referenceable) string {
	return fmt.Sprintf("%s|%s", path.String(), addr.String())
}

// vertexReferenceablePath returns the path in which the given vertex can be
// referenced. This is the path that its results from ReferenceableAddrs
// are considered to be relative to.
//
// Only GraphNodeModulePath implementations can be referenced, so this method will
// panic if the given vertex does not implement that interface.
func vertexReferenceablePath(v dag.Vertex) addrs.Module {
	sp, ok := v.(GraphNodeModulePath)
	if !ok {
		// Only nodes with paths can participate in a reference map.
		panic(fmt.Errorf("vertexMapKey on vertex type %T which doesn't implement GraphNodeModulePath", sp))
	}

	if outside, ok := v.(GraphNodeReferenceOutside); ok {
		// Vertex is referenced from a different module than where it was
		// declared.
		path, _ := outside.ReferenceOutside()
		return path
	}

	// Vertex is referenced from the same module as where it was declared.
	return sp.ModulePath()
}

// vertexReferencePath returns the path in which references _from_ the given
// vertex must be interpreted.
//
// Only GraphNodeModulePath implementations can have references, so this method
// will panic if the given vertex does not implement that interface.
func vertexReferencePath(v dag.Vertex) addrs.Module {
	sp, ok := v.(GraphNodeModulePath)
	if !ok {
		// Only nodes with paths can participate in a reference map.
		panic(fmt.Errorf("vertexReferencePath on vertex type %T which doesn't implement GraphNodeModulePath", v))
	}

	if outside, ok := v.(GraphNodeReferenceOutside); ok {
		// Vertex makes references to objects in a different module than where
		// it was declared.
		_, path := outside.ReferenceOutside()
		return path
	}

	// Vertex makes references to objects in the same module as where it
	// was declared.
	return sp.ModulePath()
}

// referenceMapKey produces keys for the "edges" map. "referrer" is the vertex
// that the reference is from, and "addr" is the address of the object being
// referenced.
//
// The result is an opaque string that includes both the address of the given
// object and the address of the module instance that object belongs to.
//
// Only GraphNodeModulePath implementations can be referrers, so this method will
// panic if the given vertex does not implement that interface.
func (m *ReferenceMap) referenceMapKey(referrer dag.Vertex, addr addrs.Referenceable) string {
	path := vertexReferencePath(referrer)
	return m.mapKey(path, addr)
}

// NewReferenceMap is used to create a new reference map for the
// given set of vertices.
func NewReferenceMap(vs []dag.Vertex) ReferenceMap {
	// Build the lookup table
	m := make(ReferenceMap)
	for _, v := range vs {
		// We're only looking for referenceable nodes
		rn, ok := v.(GraphNodeReferenceable)
		if !ok {
			continue
		}

		path := vertexReferenceablePath(v)

		// Go through and cache them
		for _, addr := range rn.ReferenceableAddrs() {
			key := m.mapKey(path, addr)
			m[key] = append(m[key], v)
		}
	}

	return m
}

// ReferencesFromConfig returns the references that a configuration has
// based on the interpolated variables in a configuration.
func ReferencesFromConfig(body hcl.Body, schema *configschema.Block) []*addrs.Reference {
	if body == nil {
		return nil
	}
	refs, _ := lang.ReferencesInBlock(body, schema)
	return refs
}

// appendResourceDestroyReferences identifies resource and resource instance
// references in the given slice and appends to it the "destroy-phase"
// equivalents of those references, returning the result.
//
// This can be used in the References implementation for a node which must also
// depend on the destruction of anything it references.
func appendResourceDestroyReferences(refs []*addrs.Reference) []*addrs.Reference {
	given := refs
	for _, ref := range given {
		switch tr := ref.Subject.(type) {
		case addrs.Resource:
			newRef := *ref // shallow copy
			newRef.Subject = tr.Phase(addrs.ResourceInstancePhaseDestroy)
			refs = append(refs, &newRef)
		case addrs.ResourceInstance:
			newRef := *ref // shallow copy
			newRef.Subject = tr.Phase(addrs.ResourceInstancePhaseDestroy)
			refs = append(refs, &newRef)
		}
		// FIXME: Using this method in module expansion references,
		// May want to refactor this method beyond resources
	}
	return refs
}

func modulePrefixStr(p addrs.ModuleInstance) string {
	return p.String()
}

func modulePrefixList(result []string, prefix string) []string {
	if prefix != "" {
		for i, v := range result {
			result[i] = fmt.Sprintf("%s.%s", prefix, v)
		}
	}

	return result
}
