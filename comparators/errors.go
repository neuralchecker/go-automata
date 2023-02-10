package comparators

import "errors"

var (
	ErrNoCounterexampleFound = errors.New("no counterexample found")
	ErrAlphabetsDiffer       = errors.New("alphabets differ")
)
