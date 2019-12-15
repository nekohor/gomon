package gomon

import (
	"encoding/json"
	"log"
	"net"
)

func RunTcpServer(app *Application) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", app.Ctx.Cfg.GetPort())
	CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err)

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		handleClient(app, conn)
		conn.Close()
	}

}

func handleClient(app *Application, conn net.Conn) {

	requestBuffNum := 8192
	request := make([]byte, requestBuffNum)

	readLen, err := conn.Read(request)
	CheckError(err)

	if requestBuffNum > readLen {
		log.Println(string(request))
		log.Println(readLen)
	} else {
		panic("readLen is not smaller than requestBuffNum")
	}

	coils := ""
	b, err := json.MarshalIndent(coils, "", "  ")
	CheckError(err)

	//log.Println(coils)
	//defer conn.Close()

	conn.Write(b)
}
