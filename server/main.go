package main

import (
	"fmt"
	"github.com/Hwisaek/go-chat/common"
	"log"
	"net"
)

var (
	clients = make([]common.Client, 0)
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

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := common.NewClient(conn)
		clients = append(clients, client)
		go func() {
			for {
				msg, err := client.Read()
				if err != nil {
					return
				}

				log.Println(msg)
				for i := len(clients) - 1; i >= 0; i-- {
					if err := clients[i].Write(msg); err != nil {
						clients = append(clients[:i], clients[i+1:]...)
					}
				}
			}
		}()
	}

}
