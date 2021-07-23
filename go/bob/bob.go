// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {

	if strings.Trim(strings.Trim(strings.Trim(strings.Trim(remark, "\n"), "\t"), "\r"), " ") == "" {
		return "Fine. Be that way!"
	}
	if remark[len(strings.Trim(remark, " "))-1] == '?' {
		if remark == strings.ToUpper(remark) {
			for _, c := range remark {
				if int(c) >= 65 && int(c) <= 90 {
					return "Calm down, I know what I'm doing!"
				}
			}
		}
		return "Sure."
	}
	if remark == strings.ToUpper(remark) {
		for _, c := range remark {
			if int(c) >= 65 && int(c) <= 90 {
				return "Whoa, chill out!"
			}
		}
		return "Whatever."
	}

	return "Whatever."
}
