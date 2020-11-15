package simplestore

import (
	"testing"

	"github.com/kapustaprusta/anagram_service/internal/app/solver/anagramsolver"
	"github.com/stretchr/testify/assert"
)

var (
	dictionary = []string{
		"живу",
		"вижу",
		"foobar",
		"boofar",
		"barfoo",
		"AbbA",
		"aBBa",
	}
)

func TestSetDictionry(t *testing.T) {
	store := NewStore(anagramsolver.NewSolver())
	store.SetDictionary(dictionary)
	assert.Equal(t, []string{"живу", "вижу"}, store.GetAnagrams("жуви"))

	var expected []string
	store.SetDictionary(dictionary[2:])
	assert.Equal(t, expected, store.GetAnagrams("жуви"))
}

func TestGetAnagrams(t *testing.T) {
	store := NewStore(anagramsolver.NewSolver())
	store.SetDictionary(dictionary)
	tests := []struct {
		word     string
		expected []string
	}{
		{
			word: "foobar",
			expected: []string{
				"foobar",
				"boofar",
				"barfoo",
			},
		},
		{
			word: "FOObar",
			expected: []string{
				"foobar",
				"boofar",
				"barfoo",
			},
		},
		{
			word: "rabfoo",
			expected: []string{
				"foobar",
				"boofar",
				"barfoo",
			},
		},
		{
			word: "abba",
			expected: []string{
				"AbbA",
				"aBBa",
			},
		},
		{
			word: "ABBA",
			expected: []string{
				"AbbA",
				"aBBa",
			},
		},
		{
			word: "baba",
			expected: []string{
				"AbbA",
				"aBBa",
			},
		},
		{
			word: "живу",
			expected: []string{
				"живу",
				"вижу",
			},
		},
		{
			word: "жИВу",
			expected: []string{
				"живу",
				"вижу",
			},
		},
		{
			word: "жуви",
			expected: []string{
				"живу",
				"вижу",
			},
		},
		{
			word: "yoda",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, store.GetAnagrams(test.word))
	}
}

func BenchmarkGetAnagrams(b *testing.B) {
	solver := anagramsolver.NewSolver()
	store := NewStore(solver)
	store.SetDictionary(dictionary)

	for i := 0; i < b.N; i++ {
		store.GetAnagrams("foobar")
	}
}
