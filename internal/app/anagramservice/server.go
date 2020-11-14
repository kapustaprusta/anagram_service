package anagramservice

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kapustaprusta/anagram_service/internal/app/store"
)

type server struct {
	store store.Store
}

func newServer(store store.Store) *server {
	return &server{
		store: store,
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("from: %s request: %s method: %s\n", r.RemoteAddr, r.RequestURI, r.Method)
	switch r.URL.Path {
	case "/get":
		s.handleGet(w, r)
	case "/load":
		s.handleLoad(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s *server) handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	wordQuery := r.URL.Query()["word"]
	if len(wordQuery) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	anagrams := s.store.Dictionary().FindAnagrams(wordQuery[0])
	anagramsRaw, err := json.Marshal(&anagrams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(anagramsRaw)
}

func (s *server) handleLoad(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	wordsRaw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var wordsJSON []interface{}
	err = json.Unmarshal(wordsRaw, &wordsJSON)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var words []string
	for _, word := range wordsJSON {
		words = append(words, word.(string))
	}

	s.store.Dictionary().SetWords(words)
	w.WriteHeader(http.StatusOK)
}
