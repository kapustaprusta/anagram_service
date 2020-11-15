package anagramservice

import (
	"net/http"

	"github.com/kapustaprusta/anagram_service/internal/app/solver/anagramsolver"
	"github.com/kapustaprusta/anagram_service/internal/app/store/simplestore"
)

// Start launches the server
// with corresponding configuration
func Start(config *Config) error {
	solver := anagramsolver.NewSolver()
	store := simplestore.NewStore(solver)

	return http.ListenAndServe(config.BindAddr, newServer(store))
}
