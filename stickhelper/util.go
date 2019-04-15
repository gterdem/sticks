// Package stickhelper is used in solution for string utility functions such as splitting a string after or before the given value and struct enums
// Credit to: https://www.dotnetperls.com/between-before-after-go
package stickhelper

import "strings"

// After used to get substring after a string.
func After(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

// Before used to get substring before a string.
func Before(value string, a string) string {
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}
