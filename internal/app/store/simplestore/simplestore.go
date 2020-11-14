package simplestore

import (
	"sync"

	"github.com/kapustaprusta/anagram_service/internal/app/solver"
)

// Store ...
type Store struct {
	mutex      *sync.Mutex
	solver     solver.Solver
	dictionary []string
}

// NewStore ...
func NewStore(solver solver.Solver) *Store {
	return &Store{
		mutex:      &sync.Mutex{},
		solver:     solver,
		dictionary: make([]string, 0),
	}
}

// SetDictionary ...
func (s *Store) SetDictionary(vocabulary []string) {
	// Lock
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Set new dictionary
	s.dictionary = vocabulary
}

// Anagrams ...
func (s *Store) Anagrams(word string) []string {
	// Lock
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Find anagrams in dictionary
	var anagrams []string
	for _, wordInDictionary := range s.dictionary {
		if s.solver.Solve(wordInDictionary, word) {
			anagrams = append(anagrams, wordInDictionary)
		}
	}

	return anagrams
}
