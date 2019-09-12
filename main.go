package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	URL      string
	APIToken string
	client   *http.Client
}

// IPAddress struct for requests to an IP address service
type IPAddress struct {
	URL string
}

func main() {
	api := Cloudflare("https://api.cloudflare.com")
	fmt.Println(api.URL)
}

// Cloudflare returns a CloudflareAPI struct
// populated with the required data to build requests
func Cloudflare(url string) *CloudflareAPI {
	return &CloudflareAPI{
		URL:      url,
		APIToken: "abcd1234",
		client:   &http.Client{Timeout: time.Second * 10},
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

// ListDNSRecords retrieves a list of the current DNS records
func ListDNSRecords(c *CloudflareAPI) []byte {
	resp, err := http.Get(c.URL)
	if err != nil {
		log.Fatal(err)
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
	reqBody, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	req := c.cloudflareRequest(reqBody)
	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return respBody
}

func (c *CloudflareAPI) cloudflareRequest(body []byte) *http.Request {
	req, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIToken)

	return req
}
