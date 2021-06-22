package apigroups

// Flatten structured object into unstructured.
func Flatten(in []string) []interface{} {
	flattened := make([]interface{}, len(in))

	for _, value := range in {
		flattened = append(flattened, value)
	}

	return flattened
}
