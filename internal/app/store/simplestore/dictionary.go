package simplestore

import (
	"strings"
	"sync"

	"github.com/kapustaprusta/anagram_service/internal/app/store"
)

// Dictionary ...
type Dictionary struct {
	mutex    *sync.Mutex
	anagrams map[string][]string
}

// SetWords ...
func (d *Dictionary) SetWords(words []string) {
	d.anagrams = make(map[string][]string)
	for i := 0; i < len(words); i++ {
		d.mutex.Lock()
		d.anagrams[words[i]] = append(d.anagrams[words[i]], words[i])
		d.mutex.Unlock()
		for j := i + 1; j < len(words); j++ {
			if store.CalcSumOfBytes(strings.ToLower(words[i])) == store.CalcSumOfBytes(strings.ToLower(words[j])) {
				d.mutex.Lock()
				d.anagrams[words[i]] = append(d.anagrams[words[i]], words[j])
				d.anagrams[words[j]] = append(d.anagrams[words[j]], words[i])
				d.mutex.Unlock()
			}
		}
	}
}

// FindAnagrams ...
func (d *Dictionary) FindAnagrams(word string) []string {
	d.mutex.Lock()
	anagrams := d.anagrams[word]
	d.mutex.Unlock()

	return anagrams
}
