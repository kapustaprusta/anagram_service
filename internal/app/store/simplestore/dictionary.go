package simplestore

import (
	"strings"
	"sync"

	"github.com/kapustaprusta/anagram_service/internal/app/store"
)

// Dictionary implement the interface of store.Dictionary
type Dictionary struct {
	mutex    *sync.Mutex      // read-write mutex
	anagrams map[int][]string // key: sum of word bytes, val: slice of words with such sum of bytes
}

// SetWords save slice of words to Dictionary
func (d *Dictionary) SetWords(words []string) {
	d.mutex.Lock()
	d.anagrams = make(map[int][]string)
	d.mutex.Unlock()
	for _, word := range words {
		sumOfBytes := store.CalcSumOfBytes(strings.ToLower(word))
		d.mutex.Lock()
		d.anagrams[sumOfBytes] = append(d.anagrams[sumOfBytes], word)
		d.mutex.Unlock()
	}
}

// FindAnagrams retrieve anagrams of word from Dictionary
func (d *Dictionary) FindAnagrams(word string) []string {
	sumOfBytes := store.CalcSumOfBytes(strings.ToLower(word))
	d.mutex.Lock()
	anagrams := d.anagrams[sumOfBytes]
	d.mutex.Unlock()

	return anagrams
}
