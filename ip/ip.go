package ip

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const clientURL = "https://ipv4.icanhazip.com"

// Client struct for requests to an IP address service
type Client struct {
	URL string
}

// NewClient returns a client for requests to icanhazip.com
func NewClient() *Client {
	return &Client{URL: clientURL}
}

// GetAddress retrieves the current external IP address
func (c *Client) GetAddress() string {
	return cleanIP(c.getIP())
}

func (c *Client) getIP() []byte {
	resp, err := http.Get(c.URL)
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
