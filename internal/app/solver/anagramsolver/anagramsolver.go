package anagramsolver

import (
	"sort"
	"strings"
	"unicode/utf8"
)

// RuneSlice declares slice of runes
type RuneSlice []rune

func (r RuneSlice) Len() int {
	return len(r)
}

func (r RuneSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RuneSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func stringToRuneSlice(s string) RuneSlice {
	idx := 0
	r := make(RuneSlice, utf8.RuneCountInString(s))
	for _, val := range s {
		r[idx] = val
		idx++
	}

	return r
}

// AnagramSolver implements the
// interface of solver.Solver
// with caching of processed words
type AnagramSolver struct {
	wordsCache map[string]RuneSlice
}

// NewSolver returns pointer
// to the AnagramSolver
func NewSolver() *AnagramSolver {
	return &AnagramSolver{
		wordsCache: make(map[string]RuneSlice),
	}
}

// Solve defines that strings are anagrams
func (s *AnagramSolver) Solve(s1 string, s2 string) bool {
	// Check length
	if len(s1) != len(s2) {
		return false
	}

	// Convert strings to slices of rune
	var isExist bool
	var r1 RuneSlice
	var r2 RuneSlice

	if r1, isExist = s.wordsCache[s1]; !isExist {
		r1 = stringToRuneSlice(strings.ToLower(s1))
		sort.Sort(r1)
		s.wordsCache[s1] = r1
	}

	if r2, isExist = s.wordsCache[s2]; !isExist {
		r2 = stringToRuneSlice(strings.ToLower(s2))
		sort.Sort(r2)
		s.wordsCache[s2] = r2
	}

	//Find difference between runes
	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}
