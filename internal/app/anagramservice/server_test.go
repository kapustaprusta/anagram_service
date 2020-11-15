package anagramservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/kapustaprusta/anagram_service/internal/app/solver/anagramsolver"
	"github.com/kapustaprusta/anagram_service/internal/app/store/simplestore"
	"github.com/stretchr/testify/assert"
)

var (
	client = &http.Client{
		Timeout: time.Second,
	}
)

const (
	APIGet  = "/get"
	APILoad = "/load"
)

type TestCase struct {
	Method string
	Path   string
	Query  string
	Status int
	Result []string
}

func TestAPI(t *testing.T) {
	solver := anagramsolver.NewSolver()
	store := simplestore.NewStore(solver)
	testServer := httptest.NewServer(newServer(store))
	testCases := []TestCase{
		{
			Method: http.MethodPost,
			Path:   APILoad,
			Query:  `["foobar", "boofar", "живу", "вижу", "Abba", "bABa"]`,
			Status: http.StatusOK,
		},
		{
			Method: http.MethodGet,
			Path:   APIGet,
			Query:  "word=foobar",
			Status: http.StatusOK,
			Result: []string{"foobar", "boofar"},
		},
		{
			Method: http.MethodGet,
			Path:   "/notfound",
			Status: http.StatusNotFound,
		},
		{
			Method: http.MethodPost,
			Path:   APIGet,
			Status: http.StatusMethodNotAllowed,
		},
		{
			Method: http.MethodGet,
			Path:   APILoad,
			Status: http.StatusMethodNotAllowed,
		},
		{
			Method: http.MethodGet,
			Path:   APIGet,
			Status: http.StatusBadRequest,
		},
		{
			Method: http.MethodPost,
			Path:   APILoad,
			Status: http.StatusBadRequest,
		},
	}

	runTests(t, testServer, testCases)
}

func runTests(t *testing.T, ts *httptest.Server, testCases []TestCase) {
	var err error
	var res []string
	var req *http.Request
	var reqBody *strings.Reader

	for idx, testCase := range testCases {
		caseName := fmt.Sprintf("case %d: [%s] %s %s", idx, testCase.Method, testCase.Path, testCase.Query)
		if testCase.Method == http.MethodPost {
			reqBody = strings.NewReader(testCase.Query)
			req, err = http.NewRequest(testCase.Method, ts.URL+testCase.Path, reqBody)
		} else {
			req, err = http.NewRequest(testCase.Method, ts.URL+testCase.Path+"?"+testCase.Query, nil)
		}

		if err != nil {
			t.Errorf("[%s] create new request error: %v", caseName, err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("[%s] request error: %v", caseName, err)
			continue
		}

		if resp.StatusCode != testCase.Status {
			t.Errorf("[%s] expected http status %v, got %v", caseName, testCase.Status, resp.StatusCode)
			continue
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if testCase.Path != APIGet || testCase.Status != http.StatusOK {
			if len(body) != 0 {
				t.Errorf("[%s] not empty response body", caseName)
			}

			continue
		}

		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Errorf("[%s] cant unpack json: %v", caseName, err)
			continue
		}

		assert.Equal(t, testCase.Result, res)
	}
}
