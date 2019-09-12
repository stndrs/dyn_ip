package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// ResponseInfo can be used for Cloudflare Messages and Errors
type ResponseInfo struct {
	Code       int            `json:"code"`
	Message    string         `json:"message"`
	ErrorChain []ResponseInfo `json:"error_chain"`
}

// Response contains response information from the Cloudflare API
type Response struct {
	Success  bool           `json:"success"`
	Errors   []ResponseInfo `json:"errors"`
	Messages []ResponseInfo `json:"messages"`
}

// ResultInfo contains information from a Cloudflare API response
type ResultInfo struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalPages int `json:"total_pages"`
	Count      int `json:"count"`
	Total      int `json:"total_count"`
}

// DNSRecord holds a Cloudflare DNS record
type DNSRecord struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Context string `json:"context"`
	Proxied bool   `json:"proxied"`
}

// DNSRecordResponse contains a DNS record from Cloudflare
type DNSRecordResponse struct {
	Result DNSRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// DNSListResponse contains an array of DNS records from Cloudflare
type DNSListResponse struct {
	Result []DNSRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// CloudflareAPI struct for requests to Cloudflare
type CloudflareAPI struct {
	URL         string
	ContentType string
	Email       string
	APIKey      string
	client      *http.Client
}

// IPAddress struct for requests to an IP address service
type IPAddress struct {
	URL string
}

func main() {
	// client := &IPAddress{URL: "ipv4.icanhazip.com"}
	// resp := CurrentIP(client)
	// ip := cleanIp(resp)
	// fmt.Println(ip)
	// req := Cloudflare()
	// resp := CurrentDNSRecords(req)
	// var records []DNSRecord
	// json.Unmarshal(resp, &records)
	// updated := make([]DNSRecord, 10)

	// for _, record := range records {
	// 	u := UpdateDNSRecord(req, &record)
	// 	fmt.Println(string(u))
	// append(updated, u)
	// }
}

// Cloudflare returns a CloudflareAPI struct
// populated with the required data to build requests
func Cloudflare() *CloudflareAPI {
	return &CloudflareAPI{
		URL:         "https://api.cloudflare.com",
		ContentType: "application/json",
		Email:       "email@example.com",
		APIKey:      "abcd1234",
		client:      &http.Client{Timeout: time.Second * 10},
	}
}

// CurrentIP retrieves the current external IP address
func CurrentIP(ip *IPAddress) []byte {
	resp, err := http.Get(ip.URL)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func cleanIP(body []byte) string {
	str := string(body)
	return strings.Trim(str, "\n")
}

// CurrentDNSRecords retrieves a list of the current DNS records
func CurrentDNSRecords(c *CloudflareAPI) []byte {
	resp, err := http.Get(c.URL)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return body
}

// UpdateDNSRecord makes a request to the API to update a DNS record
func UpdateDNSRecord(c *CloudflareAPI, d *DNSRecord) []byte {
	// resp, err := http.Post(r.URL, record)
	return nil
}
