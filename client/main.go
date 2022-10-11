package main

import (
	"bufio"
	"fmt"
	"github.com/Hwisaek/go-chat/common"
	"log"
	"net"
	"os"
)

var (
	s = bufio.NewScanner(os.Stdin)
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", common.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	client := common.NewClient(conn)

	go func() {
		for {
			msg, err := client.Read()
			if err != nil {
				return
			}

			log.Println(msg.Body)
		}
	}()
	for s.Scan() {
		msg := s.Text()

		err = client.Write(msg)
		if err != nil {
			return
		}
	}
}
