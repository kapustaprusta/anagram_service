package simplestore

import (
	"sync"

	"github.com/kapustaprusta/anagram_service/internal/app/store"
)

// Store ...
type Store struct {
	dictionary *Dictionary
}

// NewStore ...
func NewStore() *Store {
	return &Store{}
}

// Dictionary ...
func (s *Store) Dictionary() store.Dictionary {
	if s.dictionary == nil {
		s.dictionary = &Dictionary{
			mutex:    &sync.Mutex{},
			anagrams: make(map[string][]string),
		}
	}

	return s.dictionary
}
