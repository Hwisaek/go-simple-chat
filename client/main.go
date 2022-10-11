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
	defer conn.Close()

	client := common.NewClient(conn)

	go func() {
		for {
			msg, readErr := client.Read()
			if readErr != nil {
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
