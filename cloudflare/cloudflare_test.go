package cloudflare

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateRecords(t *testing.T) {
	t.Run("update a list of records", func(t *testing.T) {
		data := dnsListData()
		resp := dnsListResponse()

		srv := makeServer(http.StatusOK, data)
		defer srv.Close()

		client := NewClient(srv.URL)
		client.UpdateRecords(resp, "192.168.0.1")
	})

	t.Run("skip an update if the IP address has not changed", func(t *testing.T) {
		c := NewClient("")
		resp := dnsListResponse()

		c.UpdateRecords(resp, "127.0.0.1")
	})
}

func TestDNSRecords(t *testing.T) {
	t.Run("get the current DNS records", func(t *testing.T) {
		data := dnsListData()
		resp := dnsListResponse()
		srv := makeServer(http.StatusOK, data)
		defer srv.Close()

		client := NewClient(srv.URL)
		result := client.DNSRecords().Result[0]

		got := result.ID
		want := resp.Result[0].ID
		assertEqual(t, got, want)
	})
}

func dnsRecordResponse() *DNSRecordResponse {
	data := dnsRecordData()
	var resp DNSRecordResponse
	json.Unmarshal(data, &resp)

	return &resp
}

func dnsListResponse() *DNSListResponse {
	data := dnsListData()
	var resp DNSListResponse
	json.Unmarshal(data, &resp)

	fmt.Printf("resp: %+v", resp)
	return &resp
}

// A Mock DNS Record as a byte array
func dnsRecordData() []byte {
	bytes := []byte(`{
		"result": {
			"id": "id_string",
			"name": "example.com",
			"type": "A",
			"content": "127.0.0.1",
			"proxied": true
		},
		"success": true
	}`)
	return bytes
}

// A mock list of DNS records
func dnsListData() []byte {
	bytes := []byte(`{
		"result": [{
			"id": "id_string",
			"name": "example.com",
			"type": "A",
			"content": "127.0.0.1",
			"proxied": true
		}],
		"success": true
	}`)
	return bytes
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
