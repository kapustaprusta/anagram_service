package simplestore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	words = []string{
		"foobar",
		"barfoo",
		"boofar",
		"живу",
		"вижу",
		"Abba",
		"BaBa",
	}
)

func TestAddWords(t *testing.T) {
	expectedAnagrams := map[string][]string{
		"Abba":   {"Abba", "BaBa"},
		"BaBa":   {"Abba", "BaBa"},
		"живу":   {"живу", "вижу"},
		"вижу":   {"живу", "вижу"},
		"foobar": {"foobar", "barfoo", "boofar"},
		"barfoo": {"foobar", "barfoo", "boofar"},
		"boofar": {"foobar", "barfoo", "boofar"},
	}

	s := NewStore()
	d := s.Dictionary().(*Dictionary)
	d.AddWords(words)

	assert.Equal(t, expectedAnagrams, d.anagrams)
}

func TestFindAnagrams(t *testing.T) {
	expectedAnagrams := map[string][]string{
		"test":   nil,
		"Abba":   {"Abba", "BaBa"},
		"BaBa":   {"Abba", "BaBa"},
		"живу":   {"живу", "вижу"},
		"вижу":   {"живу", "вижу"},
		"foobar": {"foobar", "barfoo", "boofar"},
		"barfoo": {"foobar", "barfoo", "boofar"},
		"boofar": {"foobar", "barfoo", "boofar"},
	}

	s := NewStore()
	d := s.Dictionary().(*Dictionary)
	d.AddWords(words)

	for word, anagrams := range expectedAnagrams {
		assert.Equal(t, anagrams, d.FindAnagrams(word))
	}
}
