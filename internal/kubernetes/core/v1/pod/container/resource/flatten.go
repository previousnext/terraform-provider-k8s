package resource

import corev1 "k8s.io/api/core/v1"

// Flatten structured object into unstructured.
func Flatten(in corev1.ResourceList) map[string]interface{} {
	flattened := map[string]interface{}{}

	if cpu, ok := in[FieldCPU]; ok {
		flattened[FieldCPU] = cpu.String()
	}

	if memory, ok := in[FieldMemory]; ok {
		flattened[FieldMemory] = memory.String()
	}

	return flattened
}
