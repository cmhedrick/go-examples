package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	device  string = "wlp1s0"
	snaplen int32  = 65535
	promisc bool   = false
	err     error
	timeout time.Duration = -1 * time.Second
	handle  *pcap.Handle
)

func main() {
	fmt.Println("[+] Running")

	handle, err = pcap.OpenLive(device, snaplen, promisc, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	fmt.Println("[-] Setting filter")

	var filter string = "src host 192.168.0.8 and icmp"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[+] Filter set")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	fmt.Println("[+] Waiting ping")

	for packet := range packetSource.Packets() {
		fmt.Println("Someone Pinged me!!")
		fmt.Println(packet)
		fmt.Println("--------------------")
	}
}
