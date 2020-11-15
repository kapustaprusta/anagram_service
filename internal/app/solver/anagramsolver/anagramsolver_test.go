package anagramsolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToRuneSlice(t *testing.T) {
	tests := []struct {
		s        string
		expected RuneSlice
	}{
		{
			s:        "foobar",
			expected: RuneSlice{'f', 'o', 'o', 'b', 'a', 'r'},
		},
		{
			s:        "живу",
			expected: RuneSlice{'ж', 'и', 'в', 'у'},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, stringToRuneSlice(test.s))
	}
}

func TestSolve(t *testing.T) {
	solver := NewSolver()
	tests := []struct {
		words    [2]string
		expected bool
	}{
		{
			words:    [2]string{"foobar", "qwerty"},
			expected: false,
		},
		{
			words:    [2]string{"foobar", "yoda"},
			expected: false,
		},
		{
			words:    [2]string{"foobar", "foobar"},
			expected: true,
		},
		{
			words:    [2]string{"fOoBaR", "foobar"},
			expected: true,
		},
		{
			words:    [2]string{"BaRfoo", "foobar"},
			expected: true,
		},
		{
			words:    [2]string{"живу", "живу"},
			expected: true,
		},
		{
			words:    [2]string{"живу", "вижу"},
			expected: true,
		},
		{
			words:    [2]string{"жИвУ", "вижу"},
			expected: true,
		},
		{
			words:    [2]string{"ex", "no"},
			expected: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, solver.Solve(test.words[0], test.words[1]))
	}
}

func BenchmarkSolve(b *testing.B) {
	solver := NewSolver()
	for i := 0; i < b.N; i++ {
		solver.Solve("foobar", "boofar")
	}
}
