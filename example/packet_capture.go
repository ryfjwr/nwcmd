package main

import (
	"fmt"
	"io"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	pcapFile string = "http.cap"
	handle   *pcap.Handle
	err      error
)

func main() {
	// network interface を指定できる？
	log.SetFlags(log.Llongfile)
	//log.setFlags(log.Llongfile)
	handle, err := pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	// ipv4, ipv6, tcp, udp は区別できる？
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for {
		packet, err := packetSource.NextPacket()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("error...", err)
			continue
		}

		fmt.Println(packet)
	}
}
