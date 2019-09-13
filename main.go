package main

import (
	cloudflare "github.com/stndrs/dyn_ip/cloudflare"
	ip "github.com/stndrs/dyn_ip/ip"
)

func main() {
	ipClient := ip.NewClient()
	cfClient := cloudflare.NewClient("https://api.cloudflare.com")

	list := cfClient.DNSRecords()
	cfClient.UpdateRecords(list, ipClient.GetAddress())
}
