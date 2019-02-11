package labels

// Expand will return a structured object.
func Expand(in map[string]interface{}) map[string]string {
	labels := make(map[string]string)

	for k, v := range in {
		labels[k] = v.(string)
	}

	return labels
}
