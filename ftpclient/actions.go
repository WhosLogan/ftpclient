package ftpclient

import (
	"errors"
	"strconv"
)

func (c *FtpClient) GetDirectoryList() (string, error) {
	err := c.connectData()
	if err != nil {
		return "", err
	}

	_, err = c.sendCommand("LIST")
	if err != nil {
		return "", err
	}

	dir := <-c.dataChan

	_, _ = c.readData() // Discard second response

	return string(dir), nil
}

func (c *FtpClient) ChangeDirectory(path string) error {
	data, err := c.sendCommand("CWD " + path)
	if err != nil {
		return err
	}

	status, err := strconv.Atoi(string(data[0:3]))
	if err != nil {
		c.conn = nil
		return errors.New("unable to read status code")
	}

	if status == statusDirectoryChanged {
		return nil
	}

	return errors.New("unable to change directory")
}
