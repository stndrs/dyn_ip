package ip

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	assertEqual(t, client.URL, clientURL)
}

func TestCurrent(t *testing.T) {
	data := []byte("127.0.0.1\n")
	srv := makeServer(http.StatusOK, data)
	defer srv.Close()

	client := &Client{URL: srv.URL}
	got := client.GetAddress()
	want := "127.0.0.1"
	assertEqual(t, got, want)
}

// makeServer returns an httptest server for testing
// api requests
func makeServer(status int, data []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write(data)
	}))
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
