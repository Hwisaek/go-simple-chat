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

	buffer := make([]byte, common.BuffSize)
	go func() {
		for {
			n, err := client.Read(buffer)
			if err != nil {
				return
			}

			log.Println(string(buffer[:n]))
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
