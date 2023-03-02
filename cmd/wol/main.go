package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/insomniacslk/wol"
)

var (
	flagMAC = flag.String("m", "", "Mac address to send Wake-on-LAN packet to")
)

func main() {
	flag.Parse()
	if *flagMAC == "" {
		log.Fatalf("MAC address not specified")
	}
	if err := wol.Send(*flagMAC); err != nil {
		log.Fatalf("Failed to send WOL packet to %s: %v", *flagMAC, err)
	}
	fmt.Printf("Magic packet sent to %s\n", *flagMAC)
}
