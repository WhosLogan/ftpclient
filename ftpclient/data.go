package ftpclient

import (
	"errors"
	"io"
	"net"
	"strconv"
	"strings"
)

func (c *FtpClient) connectData() error {
	data, err := c.sendCommand("PASV")
	if err != nil {
		return err
	}

	status, err := strconv.Atoi(string(data[0:3]))
	if err != nil {
		return errors.New("unable to read status code from server")
	}

	if status != statusEnteringPassive {
		return errors.New("unable to enter passive mode")
	}

	ipAndPorts := strings.Split(strings.Split(strings.Split(string(data), ")")[0], "(")[1], ",")
	ip := ipAndPorts[0] + "." + ipAndPorts[1] + "." + ipAndPorts[2] + "." + ipAndPorts[3]
	portPart1, _ := strconv.Atoi(ipAndPorts[4])
	portPart2, _ := strconv.Atoi(ipAndPorts[5])
	port := portPart1*256 + portPart2

	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	c.dataConn = conn
	c.dataChan = make(chan []byte)
	c.awaiting = true

	go c.handleConnection()

	return nil
}

func (c *FtpClient) handleConnection() {
	data, err := io.ReadAll(c.dataConn)
	if err != nil {
		return
	}

	c.awaiting = false
	c.dataChan <- data
}
