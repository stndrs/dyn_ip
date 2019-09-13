package cloudflare

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// DNSRecord holds a Cloudflare DNS record
type DNSRecord struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Proxied bool   `json:"proxied"`
}

// DNSRecordResponse contains a DNS record from Cloudflare
type DNSRecordResponse struct {
	Result  DNSRecord `json:"result"`
	Success bool      `json:"success"`
}

// DNSListResponse contains an array of DNS records from Cloudflare
type DNSListResponse struct {
	Result  []DNSRecord `json:"result"`
	Success bool        `json:"success"`
}

// Client struct for requests to Cloudflare
type Client struct {
	URL      string
	APIToken string
	client   *http.Client
}

// NewClient returns a Client struct
// populated with the required data to build requests
func NewClient(url string) *Client {
	return &Client{
		URL:      url,
		APIToken: "abcd1234",
		client:   &http.Client{Timeout: time.Second * 10},
	}
}

// DNSRecords retrieves a list of the current DNS records
func (c *Client) DNSRecords() *DNSListResponse {
	list := c.getDNSRecords()
	var records DNSListResponse
	json.Unmarshal(list, &records)

	return &records
}

// UpdateRecords loops over the DNS List and update the relevant records
func (c *Client) UpdateRecords(r *DNSListResponse, ip string) {
}

func (c *Client) getDNSRecords() []byte {
	req := c.request(http.MethodGet, nil)
	resp, err := c.client.Do(req)
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

// updateDNS makes a request to the API to update a DNS record
func (c *Client) updateDNS(d *DNSRecord) []byte {
	reqBody, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	req := c.request(http.MethodPost, reqBody)
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

func (c *Client) request(method string, body []byte) *http.Request {
	req, err := http.NewRequest(method, c.URL, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIToken)

	return req
}
