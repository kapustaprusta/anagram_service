package store

// Dictionary ...
type Dictionary interface {
	SetWords([]string)
	FindAnagrams(string) []string
}
