package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateDNSRecord(t *testing.T) {
	t.Run("updates a DNS record", func(t *testing.T) {
		data := mockDNSRecordResponse()
		var record DNSRecord
		json.Unmarshal(data, &record)

		srv := makeServer(http.StatusOK, data)
		c := Cloudflare(srv.URL)

		got := UpdateDNSRecord(c, &record)

		if string(got) != string(data) {
			t.Errorf("got %q, want %q", got, data)
		}
	})
}

func TestListDNSRecords(t *testing.T) {
	t.Run("gets the current DNS records", func(t *testing.T) {
		data := mockDNSListResponse()
		srv := makeServer(http.StatusOK, data)
		dns := &CloudflareAPI{URL: srv.URL}

		got := ListDNSRecords(dns)

		if string(got) != string(data) {
			t.Errorf("got %q, want %q", got, data)
		}
	})
}

func TestCurrentIP(t *testing.T) {
	t.Run("gets the current IP address", func(t *testing.T) {
		data := []byte("127.0.0.1\n")
		srv := makeServer(http.StatusOK, data)
		defer srv.Close()

		rip := &IPAddress{URL: srv.URL}

		got := CurrentIP(rip)

		if string(got) != string(data) {
			t.Errorf("got %q, want %q", got, data)
		}
	})
}

// Build a mock DNSRecord
func mockDNSRecord() *DNSRecord {
	return &DNSRecord{
		ID:      "asdlvinasoildvasdvlin",
		Name:    "example.com",
		Type:    "A",
		Context: "",
		Proxied: true,
	}
}

// Build a mock DNSRecordResponse
func mockDNSRecordResponse() []byte {
	record := mockDNSRecord()
	data, _ := json.Marshal(record)
	return data
}

// Build a mock DNSListResponse
func mockDNSListResponse() []byte {
	record := mockDNSRecord()
	records := make([]DNSRecord, 1)
	errors := make([]ResponseInfo, 0)
	messages := make([]ResponseInfo, 0)
	resultInfo := new(ResultInfo)
	records[0] = *record

	response := &DNSListResponse{
		Result: records,
		Response: Response{
			Success:  true,
			Errors:   errors,
			Messages: messages,
		},
		ResultInfo: *resultInfo,
	}

	data, _ := json.Marshal(response)
	return data
}

// makeServer returns an httptest server for testing
// api requests
func makeServer(status int, data []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write(data)
	}))
}
