package store

// Dictionary is an object that can store words and find anagrams
type Dictionary interface {
	SetWords([]string)
	FindAnagrams(string) []string
}
