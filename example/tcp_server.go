package main

import (
	"io"
	"log"
	"net"
	"time"
)

// https://zenn.dev/satumahayato010/articles/abdd898c80386c

func echoHandler(conn *net.TCPConn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, "hello...hogehoge")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}

func recieveTCPConn(ln *net.TCPListener) {
	for {
		err := ln.SetDeadline(time.Now().Add(time.Second * 60))
		if err != nil {
			log.Fatal(err)
		}
		conn, err := ln.AcceptTCP()
		if err != nil {
			log.Fatal(conn)
		}
		go echoHandler(conn)
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatal("error")
	}
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal()
	}

	recieveTCPConn(ln)
}
