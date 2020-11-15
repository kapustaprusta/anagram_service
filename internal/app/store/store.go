package store

// Store declares the interface
// which contains SetDictionary
// and GetAnagrams methods
type Store interface {
	SetDictionary([]string)
	GetAnagrams(string) []string
}
