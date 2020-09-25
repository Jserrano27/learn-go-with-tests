package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("returns the fastest url", func(t *testing.T) {
		slowServer := createDelayedServer(100 * time.Millisecond)
		fastServer := createDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("Did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("Got %q, but expected was %q", got, want)
		}
	})

	t.Run("returns an error if servers doesn't respond within the specified timeout", func(t *testing.T) {
		server := createDelayedServer(5 * time.Millisecond)

		defer server.Close()

		url := server.URL

		_, err := ConfigurableRacer(url, url, 1*time.Millisecond)

		if err == nil {
			t.Error("Expecting an error but got none")
		}
	})
}

func createDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		response.WriteHeader(http.StatusOK)
	}))
}
