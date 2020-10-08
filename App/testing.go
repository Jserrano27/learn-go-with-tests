package poker

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	leage    League
}

func (s *StubPlayerStore) getPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) recordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) getLeague() League {
	return s.leage
}

// Helpers

func NewGetScoreRequest(name string) *http.Request {
	route := fmt.Sprintf("/players/%v", name)
	request, _ := http.NewRequest(http.MethodGet, route, nil)

	return request
}

func NewPostWinRequest(name string) *http.Request {
	route := fmt.Sprintf("/players/%v", name)
	req, _ := http.NewRequest(http.MethodPost, route, nil)

	return req
}

func NewGetLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func GetLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	t.Helper()
	league, err := NewLeague(body)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return league
}

func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q but want %q", got, want)
	}
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, name string) {
	if len(store.winCalls) != 1 {
		t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != name {
		t.Errorf("Did not store correct winner. Got %q, want %q", store.winCalls[0], name)
	}
}

func AssertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got status code %d but want %d", got, want)
	}
}

func AssertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	contentType := response.Result().Header.Get("content-type")
	if contentType != want {
		t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
	}
}

func CreateTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("Could not create temporary file: %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
