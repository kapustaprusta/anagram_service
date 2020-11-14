package store

// Store is an object that can store and return a dictionary
type Store interface {
	Dictionary() Dictionary
}
