package simplestore

import (
	"strings"
	"sync"

	"github.com/kapustaprusta/anagram_service/internal/app/store"
)

// Dictionary ...
type Dictionary struct {
	mutex    *sync.Mutex
	anagrams map[int][]string // key: sum of word bytes, val: words with such sum of bytes
}

// SetWords ...
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

// FindAnagrams ...
func (d *Dictionary) FindAnagrams(word string) []string {
	sumOfBytes := store.CalcSumOfBytes(strings.ToLower(word))
	d.mutex.Lock()
	anagrams := d.anagrams[sumOfBytes]
	d.mutex.Unlock()

	return anagrams
}
