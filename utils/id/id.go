package id

import (
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Join(meta metav1.ObjectMeta) string {
	return meta.Namespace + "/" + meta.Name
}

func Split(id string) (string, string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("unexpected ID format (%q), expected: namespace/name", id)
	}

	return parts[0], parts[1], nil
}
