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

	for {
		msg, err := client.Read()
		if err != nil {
			return
		}

		log.Println(msg.Body)
		err = client.Write(msg.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}
