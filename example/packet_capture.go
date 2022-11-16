package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	handle *pcap.Handle
	err    error
)

func main() {
	// https://mrtc0.hateblo.jp/entry/2016/03/19/232252
	// network interface を指定できる？

	// https://qiita.com/kkyouhei/items/846e74c6a9653b069e5f
	snapshotLen := 1024
	timeOut := 600 * time.Second
	promiscous := false
	log.SetFlags(log.Llongfile)
	//log.setFlags(log.Llongfile)
	handle, err := pcap.OpenLive("eth0", int32(snapshotLen), promiscous, timeOut)
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
