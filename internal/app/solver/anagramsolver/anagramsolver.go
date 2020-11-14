package anagramsolver

import (
	"sort"
	"strings"
	"unicode/utf8"
)

// RuneSlice ...
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

// AnagramSolver ...
type AnagramSolver struct {
	wordsCache map[string]RuneSlice
}

// NewSolver ...
func NewSolver() *AnagramSolver {
	return &AnagramSolver{
		wordsCache: make(map[string]RuneSlice),
	}
}

// ClearCache ...
func (s *AnagramSolver) ClearCache() {
	// Clear words cache
	s.wordsCache = make(map[string]RuneSlice)
}

// Solve ...
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
		s.wordsCache[s1] = r1
	}

	if r2, isExist = s.wordsCache[s2]; !isExist {
		r2 = stringToRuneSlice(strings.ToLower(s2))
		s.wordsCache[s2] = r2
	}

	// Sort slices
	sort.Sort(r1)
	sort.Sort(r2)

	// Find difference between runes
	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}
