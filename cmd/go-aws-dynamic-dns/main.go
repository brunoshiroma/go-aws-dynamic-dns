package main

import (
	"flag"
	"log"

	client "github.com/brunoshiroma/go-aws-dynamic-dns/internal"
)

func main() {

	var oldIp string

	hostedZoneId := flag.String("z", "", "ZoneId")
	recordName := flag.String("r", "", "RecordName")
	ttl := flag.Int64("t", 300, "ttl")
	ipDiscoveryServiceURL := flag.String("i", "", "URL for ip discovery service ex. https://tiny-credit-7e1b.benchtool.workers.dev")
	flag.Parse()

	if *hostedZoneId == "" || *recordName == "" || *ipDiscoveryServiceURL == "" {
		flag.Usage()
		return
	}

	ip, err := client.GetIp(*ipDiscoveryServiceURL)
	if err != nil {
		log.Fatal(err)
	}

	oldIp, err = client.GetDNSIP(hostedZoneId, recordName)
	if err != nil {
		log.Fatal(err)
	}

	if oldIp != ip {
		log.Printf("Changing ip to %s\n", ip)
		err = client.SetDNSIP(hostedZoneId, recordName, ttl, ip)
		log.Printf("Changed ip to %s failure!!!\n\n", ip)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Changed ip to %s success!!!\n", ip)
		return
	}

	log.Printf("ip not changed from %s\n", ip)
}
