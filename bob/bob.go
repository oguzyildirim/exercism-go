// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

type Remark struct {
	remark string
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	r := Remark{strings.TrimSpace(remark)}
	switch {
	case r.isSilence():
		return "Fine. Be that way!"
	case r.isExasperated():
		return "Calm down, I know what I'm doing!"
	case r.isShouting():
		return "Whoa, chill out!"
	case r.isQuestion():
		return "Sure."
	default:
		return "Whatever."
	}
}

func (r Remark) isSilence() bool {
	return r.remark == ""
}

func (r Remark) isShouting() bool {
	hasLetters := strings.IndexFunc(r.remark, unicode.IsLetter) >= 0
	isUpcased := strings.ToUpper(r.remark) == r.remark

	return hasLetters && isUpcased
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}

func (r Remark) isExasperated() bool {
	return r.isShouting() && r.isQuestion()
}
