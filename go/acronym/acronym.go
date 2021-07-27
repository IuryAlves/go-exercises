// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate: Returns the abbreviation of a name.
func Abbreviate(s string) string {
	onlyLetters := regexp.MustCompile(`[^a-zA-Z]`)
	splitted := onlyLetters.Split(s, -1)
	var acronym = ""
	for _, ss := range splitted {
		if len(ss) > 0 {
			acronym += string(ss[0])
		}
	}

	return strings.ToUpper(acronym)
}
