package common

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
)

type Msg struct {
	Id   string
	Body string
}

type Client struct {
	Id     string
	Conn   *net.Conn
	writer *bufio.Writer
	reader *bufio.Reader
}

func NewClient(conn net.Conn) (client Client) {
	client = Client{
		Id:     fmt.Sprint(rand.Int()),
		Conn:   &conn,
		writer: bufio.NewWriter(conn),
		reader: bufio.NewReader(conn),
	}
	return
}

func (c Client) Write(s string) (err error) {
	_, err = c.writer.WriteString(c.Id + ": " + s)
	if err != nil {
		return
	}

	err = c.writer.Flush()
	if err != nil {
		return
	}
	return
}

func (c Client) Read(buffer []byte) (n int, err error) {
	n, err = c.reader.Read(buffer)
	if err != nil {
		return
	}

	return
}
