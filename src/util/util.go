package util

import (
	"bytes"
	"fmt"
	"strings"
)

func CreateKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=%s ", key, value)
	}
	return strings.Trim(b.String(), "\n")
}
