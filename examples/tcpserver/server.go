package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	service := ":8999"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	request := make([]byte, 8192)
	readLen, err := conn.Read(request)
	// defer conn.Close()
	checkError(err)
	log.Println(string(request))
	log.Println(readLen)

	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
