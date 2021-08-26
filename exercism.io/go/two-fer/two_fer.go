// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith returns a string message in the format of One for X, one for me.
func ShareWith(name string) string {
	message := "One for %s, one for me."
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf(message, name)
}
