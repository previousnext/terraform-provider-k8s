package resources

// Expand will return a structured object.
func Expand(s []interface{}) []string {
	result := make([]string, len(s), len(s))

	for k, v := range s {
		result[k] = v.(string)
	}

	return result
}
