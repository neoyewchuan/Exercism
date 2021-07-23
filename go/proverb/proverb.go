// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverbs := []string{}
	// loop through the given words
	for i := 0; i < len(rhyme); i++ {
		// the current word must not be empty
		if rhyme[i] != "" {
			// check if there is more word in the list, and the word must not be empty
			if i+1 < len(rhyme) && rhyme[i+1] != "" {
				proverbs = append(proverbs, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
			} else if i+1 == len(rhyme) {
				// this is the last word in the list
				proverbs = append(proverbs, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
			}
		}
	}
	return proverbs
}
