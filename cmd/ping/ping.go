package main

import (
	"fmt"
	"syscall"
)

type ICMP struct {
	Type           []byte
	Code           []byte
	CheckSum       []byte
	Identification []byte
	SequenceNumber []byte
	Data           []byte
}

/**
	8.8.8.8 に投げて、ping echo reply を受け取り表示する
	ICMP クライアント
	https://atmarkit.itmedia.co.jp/ait/articles/0306/13/news002.html


**/

func main() {
	// ethernet request
	ifIndex := 0

	// create icmp packet
	icmpPacket := ICMP{
		// https://www.infraexpert.com/study/tcpip4.html
		Type: []byte{0x08}, // icmp echo request -> 0x08
		Code: []byte{0x00}, // echo request の 要求 -> 0x00

	}

	// add ip header
	//
	addr := syscall.SockaddrLinklayer{
		Protocol: syscall.ETH_P_IP,
		Ifindex:  ifIndex,
	}

	//
	sendFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, int())
	if err != nil {
		fmt.Printf("socket error")
	}

	// 無限ループ
	for {

	}
}
