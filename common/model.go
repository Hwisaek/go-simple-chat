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
	Id          string
	Conn        *net.Conn
	writer      *bufio.Writer
	reader      *bufio.Reader
	codecBuffer *bytes.Buffer
	enc         *gob.Encoder
	dec         *gob.Decoder
}

func NewClient(conn net.Conn) (client Client) {
	var codecBuffer bytes.Buffer
	client = Client{
		Id:          fmt.Sprint(rand.Int()),
		Conn:        &conn,
		writer:      bufio.NewWriter(conn),
		reader:      bufio.NewReader(conn),
		codecBuffer: &codecBuffer,
		enc:         gob.NewEncoder(&codecBuffer),
		dec:         gob.NewDecoder(&codecBuffer),
	}
	return
}

func (c Client) Write(s string) (err error) {
	msg := Msg{
		Id:   c.Id,
		Body: s,
	}

	if err = c.enc.Encode(msg); err != nil {
		return err
	}

	if _, err = c.writer.Write(c.codecBuffer.Bytes()); err != nil {
		return
	}

	if err = c.writer.Flush(); err != nil {
		return
	}
	c.codecBuffer.Reset()
	return
}

func (c Client) Read() (msg Msg, err error) {
	buffer := make([]byte, 4096) // receive buffer: 4kB
	n, err := c.reader.Read(buffer)
	if err != nil {
		return
	}

	if n > 0 {
		data := buffer[:n]
		c.codecBuffer.Write(data)

		msg = Msg{}
		if err = c.dec.Decode(&msg); err != nil {
			return
		}
	}

	return msg, nil
}
