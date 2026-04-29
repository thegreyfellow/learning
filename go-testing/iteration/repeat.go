// Package iterations
package iterations

import "strings"

func Repeat(s string, count int) string {
	if count == 0 {
		return s
	}

	var res strings.Builder
	for range count {
		res.WriteString(s)
	}
	return res.String()
}
