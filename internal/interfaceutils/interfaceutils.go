package interfaceutils

// ExpandSlice converts an interface slice to a string slice.
func ExpandSlice(s []interface{}) []string {
	result := make([]string, len(s), len(s))

	for k, v := range s {
		result[k] = v.(string)
	}

	return result
}

// ExpandMap will return a structured object.
func ExpandMap(in map[string]interface{}) map[string]string {
	labels := make(map[string]string)

	for k, v := range in {
		labels[k] = v.(string)
	}

	return labels
}
