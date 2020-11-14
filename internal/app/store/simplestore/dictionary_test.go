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
	expectedAnagrams := map[int][]string{
		390:  {"Abba", "BaBa"},
		633:  {"foobar", "barfoo", "boofar"},
		1508: {"живу", "вижу"},
	}

	s := NewStore()
	d := s.Dictionary().(*Dictionary)
	d.SetWords(words)

	assert.Equal(t, expectedAnagrams, d.anagrams)
}

func TestFindAnagrams(t *testing.T) {
	expectedAnagrams := map[string][]string{
		"test":   nil,
		"Abba":   {"Abba", "BaBa"},
		"abba":   {"Abba", "BaBa"},
		"BaBa":   {"Abba", "BaBa"},
		"живу":   {"живу", "вижу"},
		"вижу":   {"живу", "вижу"},
		"foobar": {"foobar", "barfoo", "boofar"},
		"barfoo": {"foobar", "barfoo", "boofar"},
		"OOFBAR": {"foobar", "barfoo", "boofar"},
		"boofar": {"foobar", "barfoo", "boofar"},
	}

	s := NewStore()
	d := s.Dictionary().(*Dictionary)
	d.SetWords(words)

	for word, anagrams := range expectedAnagrams {
		assert.Equal(t, anagrams, d.FindAnagrams(word))
	}
}
