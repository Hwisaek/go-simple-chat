package main

import (
	"fmt"
	"github.com/Hwisaek/go-chat/common"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", common.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	client := common.NewClient(conn)

	for {
		msg, netErr := client.Read()
		if netErr != nil {
			return
		}

		log.Println(msg.Body)
		netErr = client.Write(msg.Body)
		if netErr != nil {
			log.Fatal(netErr)
		}
	}
}
