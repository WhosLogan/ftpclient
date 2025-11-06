package ftpclient

import (
	"errors"
	"fmt"
	"net"
)

type FtpClient struct {
	Username string
	Password string
	Server   string
	conn     net.Conn
}

func NewClient() *FtpClient {
	return &FtpClient{
		Username: "anonymous",
	}
}

func (c *FtpClient) Connect() error {
	if c.conn != nil {
		return errors.New("please close the active connection before opening a new one")
	}

	conn, err := net.Dial("tcp", c.Server+":21")
	if err != nil {
		return err
	}

	c.conn = conn

	return nil
}

func (c *FtpClient) Close() {
	if c.conn != nil {
		_ = c.conn.Close()
		c.conn = nil
	}
}

func (c *FtpClient) sendCommand(cmd string) ([]byte, error) {
	_, err := fmt.Fprintf(c.conn, cmd+"\r\n")
	if err != nil {
		return nil, err
	}

	received := 1024
	var response []byte

	for received == 1024 {
		reply := make([]byte, 1024)
		received, err = c.conn.Read(reply)
		if err != nil {
			if received != 1024 {
				break
			}

			return nil, err
		}

		response = append(response, reply[:received]...)
	}

	return response, nil
}
