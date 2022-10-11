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
	defer func() {
		err = listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	client := common.NewClient(conn)

	buffer := make([]byte, common.BuffSize)
	for {
		n, err := client.Read(buffer)
		if err != nil {
			return
		}

		log.Println(string(buffer[:n]))
		err = client.Write(string(buffer[:n]))
		if err != nil {
			log.Fatal(err)
		}
	}
}
