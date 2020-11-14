package anagramsolver

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	s1 := "no"
	s2 := "ex"
	solver := NewSolver()
	fmt.Println(solver.Solve(s1, s2))
}

func BenchmarkSolve(b *testing.B) {
	solver := NewSolver()
	for i := 0; i < b.N; i++ {
		solver.Solve("FOObar", "boofar")
	}
}
