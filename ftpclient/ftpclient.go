package ftpclient

import (
	"errors"
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
