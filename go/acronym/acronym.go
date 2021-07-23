// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"bytes"
	"strings"
)

// Abbreviate scovert a long sting/name to its acronym.
func Abbreviate(s string) string {
	var buf bytes.Buffer
	// Split the string into words, separated by spaces, dash and underscore
	split := strings.FieldsFunc(strings.ToUpper(s), func(r rune) bool {
		return r == '-' || r == ' ' || r == '_'
	})
	for _, v := range split {
		if len(v) > 0 && int(v[0]) >= 65 && int(v[0]) <= 90 {
			// write to bytes buffer only if the word is not empty
			// and the first letter falls between A..Z
			buf.WriteByte(v[0])
		}
	}
	return buf.String()
}
