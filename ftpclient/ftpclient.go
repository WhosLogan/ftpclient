package ftpclient

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

type FtpClient struct {
	Username string
	Password string
	Server   string
	conn     net.Conn
	dataConn net.Conn
	awaiting bool
	dataChan chan []byte
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

	data, err := c.sendCommand("USER " + c.Username)
	if err != nil {
		return err
	}

	status, err := strconv.Atoi(string(data[0:3]))
	if err != nil {
		c.conn = nil
		return errors.New("unable to read status code on authentication")
	}

	if status == statusReady {
		return nil
	}

	if status == statusNeedPassword {
		if c.Password == "" {
			c.conn = nil
			return errors.New("please set a password using the 'pass' command")
		}

		data, err := c.sendCommand("PASS " + c.Password)
		if err != nil {
			c.conn = nil
			return err
		}

		status, err := strconv.Atoi(string(data[0:3]))
		if err != nil {
			c.conn = nil
			return errors.New("unable to read status code on authentication")
		}

		if status == statusReady {
			return nil
		}

		c.conn = nil
		return translateErrorStatusCode(status)
	}

	c.conn = nil
	return translateErrorStatusCode(status)
}

func (c *FtpClient) Close() {
	if c.conn != nil {
		_, _ = c.sendCommand("QUIT")
		_ = c.conn.Close()
		c.conn = nil
	}

	if c.dataConn != nil {
		_ = c.dataConn.Close()
		c.dataConn = nil
	}

	c.awaiting = false
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
