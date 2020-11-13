package anagramservice

import (
	"net/http"

	"github.com/kapustaprusta/anagram_service/internal/app/store/simplestore"
)

// Start ...
func Start(config *Config) error {
	return http.ListenAndServe(config.BindAddr, newServer(simplestore.NewStore()))
}
