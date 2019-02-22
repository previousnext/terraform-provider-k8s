package data

// Expand will return a structured object.
func Expand(in map[string]interface{}) map[string][]byte {
	labels := make(map[string][]byte)

	for k, v := range in {
		labels[k] = []byte(v.(string))
	}

	return labels
}
