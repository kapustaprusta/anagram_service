package simplestore

import (
	"strings"

	"github.com/kapustaprusta/anagram_service/internal/app/store"
)

// Dictionary ...
type Dictionary struct {
	anagrams map[string][]string
}

// Clear ...
func (d *Dictionary) Clear() {
	d.anagrams = make(map[string][]string)
}

// AddWords ...
func (d *Dictionary) AddWords(words []string) {
	for i := 0; i < len(words); i++ {
		d.anagrams[words[i]] = append(d.anagrams[words[i]], words[i])
		for j := i + 1; j < len(words); j++ {
			if store.CalcSumOfBytes(strings.ToLower(words[i])) == store.CalcSumOfBytes(strings.ToLower(words[j])) {
				d.anagrams[words[i]] = append(d.anagrams[words[i]], words[j])
				d.anagrams[words[j]] = append(d.anagrams[words[j]], words[i])
			}
		}
	}
}

// FindAnagrams ...
func (d *Dictionary) FindAnagrams(word string) []string {
	return d.anagrams[word]
}
