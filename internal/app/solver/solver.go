package solver

// Solver ...
type Solver interface {
	ClearCache()
	Solve(string, string) bool
}
