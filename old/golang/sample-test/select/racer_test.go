package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDeplayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDeplayedServer(20 * time.Millisecond)
		fastServer := makeDeplayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		faseURL := fastServer.URL

		want := faseURL
		got, err := Racer(slowURL, faseURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDeplayedServer(25 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
