package simplestore

import (
	"fmt"
	"testing"

	"github.com/kapustaprusta/anagram_service/internal/app/solver/anagramsolver"
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

func TestAnagrams(t *testing.T) {
	solver := anagramsolver.NewSolver()
	store := NewStore(solver)
	store.SetDictionary(dictionary)
	fmt.Println(store.Anagrams(dictionary[0]))
}

func BenchmarkAnagrams(b *testing.B) {
	solver := anagramsolver.NewSolver()
	store := NewStore(solver)
	store.SetDictionary(dictionary)

	for i := 0; i < b.N; i++ {
		store.Anagrams(dictionary[2])
	}
}
