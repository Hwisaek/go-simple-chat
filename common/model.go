package common

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
)

type Msg struct {
	Id   string
	Body string
}

type Client struct {
	Id      string
	Conn    *net.Conn
	writer  *bufio.Writer
	reader  *bufio.Reader
	buffer  []byte
	network *bytes.Buffer
	enc     *gob.Encoder
	dec     *gob.Decoder
}

func NewClient(conn net.Conn) (client Client) {
	var buffer bytes.Buffer
	client = Client{
		Id:      fmt.Sprint(rand.Int()),
		Conn:    &conn,
		writer:  bufio.NewWriter(conn),
		reader:  bufio.NewReader(conn),
		buffer:  make([]byte, 0, BuffSize),
		network: &buffer,
		enc:     gob.NewEncoder(&buffer),
		dec:     gob.NewDecoder(&buffer),
	}
	return
}

func (c Client) Write(s string) (err error) {
	msg := Msg{
		Id:   c.Id,
		Body: s,
	}
	err = c.enc.Encode(msg)
	if err != nil {
		return err
	}

	_, err = c.writer.Write(c.network.Bytes())
	if err != nil {
		return
	}

	err = c.writer.Flush()
	if err != nil {
		return
	}
	return
}

func (c Client) Read() (msg Msg, err error) {
	_, err = c.reader.Read(c.buffer)
	if err != nil {
		return
	}

	c.network.Write(c.buffer)

	err = c.dec.Decode(&msg)
	if err != nil {
		return
	}

	return
}
