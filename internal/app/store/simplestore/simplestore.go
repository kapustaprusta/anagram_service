package simplestore

import (
	"sync"

	"github.com/kapustaprusta/anagram_service/internal/app/solver"
)

// Store implements the store.Store
// interface with caching of
// processed words
type Store struct {
	mutex         *sync.Mutex
	solver        solver.Solver
	dictionary    []string
	anagramsCache map[string][]string
}

// NewStore accepts a struct which
// implements the interface of solver.Solver
// and returns pointer to the Store
func NewStore(solver solver.Solver) *Store {
	return &Store{
		mutex:         &sync.Mutex{},
		solver:        solver,
		dictionary:    make([]string, 0),
		anagramsCache: make(map[string][]string),
	}
}

// SetDictionary updates the dictionary
// according to the input argument and
// clears the inner cache
func (s *Store) SetDictionary(dictionary []string) {
	// Lock
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Set new dictionary
	s.dictionary = dictionary

	// Clear cache
	s.anagramsCache = make(map[string][]string)
}

// GetAnagrams returns the anagrams
// of input argument found in the
// dictionary and updates the
// inner cache
func (s *Store) GetAnagrams(word string) []string {
	// Lock
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Find anagrams in dictionary
	var isExist bool
	var anagrams []string

	if anagrams, isExist = s.anagramsCache[word]; !isExist {
		for _, wordInDictionary := range s.dictionary {
			if s.solver.Solve(wordInDictionary, word) {
				anagrams = append(anagrams, wordInDictionary)
			}
		}

		s.anagramsCache[word] = anagrams
	}

	return anagrams
}
