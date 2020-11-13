package store

// Dictionary ...
type Dictionary interface {
	AddWords([]string)
	FindAnagrams(string) []string
}
