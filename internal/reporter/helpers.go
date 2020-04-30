package reporter

import (
	"strings"
)

func processFormValue(data string) []string {
	var res []string

	parts := strings.Split(data, "\r\n")
	// remove empty strings.
	for _, a := range parts {
		if a != "" {
			res = append(res, a)
		}
	}

	return res
}
