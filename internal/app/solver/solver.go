package solver

// Solver declares the interface
// which conatains Solve method
type Solver interface {
	Solve(string, string) bool
}
