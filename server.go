package gomon


import (
	"net"
	"log"
	"encoding/json"
	"strconv"
)

func RunServer(app *Monitor) {

	port := ":" + strconv.Itoa(app.Ctx.Cfg.Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
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

func handleClient(app *Monitor, conn net.Conn) {
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

	coils := app.RespondCoils(string(request))
	b, err := json.MarshalIndent(coils, "", "  ")
	CheckError(err)

	//log.Println(coils)

	//defer conn.Close()
	conn.Write(b)
}