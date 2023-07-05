package apigroups

// Expand will return a structured object.
func Expand(s []interface{}) []string {
	result := make([]string, len(s))

	for k, v := range s {
		if v == nil {
			// This empty string = "core api" for K8s roles.
			result[k] = ""
		} else {
			result[k] = v.(string)
		}
	}

	return result
}
