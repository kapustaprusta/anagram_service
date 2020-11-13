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

// AddWords ...
func (d *Dictionary) AddWords(words []string) {
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
	defer d.mutex.Unlock()

	return d.anagrams[word]
}
