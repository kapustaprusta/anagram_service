package simplestore

import (
	"sync"

	"github.com/kapustaprusta/anagram_service/internal/app/store"
)

// Store implement the interface of store.Store
type Store struct {
	dictionary *Dictionary // instance of dictionary
}

// NewStore create and return instance of Store
func NewStore() *Store {
	return &Store{}
}

// Dictionary return instance of store.Dictionary in case of store.Dictionary is not exist will create it
func (s *Store) Dictionary() store.Dictionary {
	if s.dictionary == nil {
		s.dictionary = &Dictionary{
			mutex:    &sync.Mutex{},
			anagrams: make(map[int][]string),
		}
	}

	return s.dictionary
}
