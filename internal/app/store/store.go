package store

// Store ...
type Store interface {
	SetDictionary([]string)
	Anagrams(string) []string
}
