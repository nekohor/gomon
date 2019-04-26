package main

import (
	"encoding/json"
	"github.com/nekohor/gomon"
	"log"
	"net"
)

func main() {
	service := ":8999"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	gomon.CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	gomon.CheckError(err)
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}
		handleClient(conn)
		conn.Close()
	}
}

func handleClient(conn net.Conn) {
	requestBuffNum := 8192
	request := make([]byte, requestBuffNum)

	readLen, err := conn.Read(request)
	gomon.CheckError(err)

	if requestBuffNum > readLen {
		log.Println(string(request))
		log.Println(readLen)
	} else {
		panic("readLen is not smaller than requestBuffNum")
	}

	app := gomon.NewGoMonitor()

	coils := app.RespondCoils(string(request))
	b, err := json.MarshalIndent(coils, "", "  ")
	gomon.CheckError(err)

	//defer conn.Close()
	conn.Write(b)
}
