package ftpclient

import (
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

func (c *FtpClient) GetDirectoryList() (string, error) {
	c.discardPending()
	err := c.readDataConnection()
	if err != nil {
		return "", err
	}

	_, err = c.sendCommand("LIST")
	if err != nil {
		_ = c.dataConn.Close()
		return "", err
	}

	dir := <-c.dataChan

	_, _ = c.readData() // Discard second response

	return string(dir), nil
}

func (c *FtpClient) ChangeDirectory(path string) error {
	c.discardPending()
	data, err := c.sendCommand("CWD " + path)
	if err != nil {
		return err
	}

	status, err := strconv.Atoi(string(data[0:3]))
	if err != nil {
		return errors.New("unable to read status code")
	}

	if status == statusDirectoryChanged {
		return nil
	}

	return errors.New("unable to change directory")
}

func (c *FtpClient) GetFile(path string) ([]byte, error) {
	err := c.readDataConnection()
	if err != nil {
		return nil, err
	}

	data, err := c.sendCommand("RETR " + path)
	if err != nil {
		return nil, err
	}

	status, err := strconv.Atoi(string(data[0:3]))
	if err != nil {
		_ = c.dataConn.Close()
		return nil, errors.New("unable to read status code")
	}

	if status == statusNoFileOrDirectory {
		_ = c.dataConn.Close()
		return nil, errors.New("no such file or directory")
	}

	file := <-c.dataChan

	_, _ = c.readData()

	return file, nil
}

func (c *FtpClient) SendFile(path string) (int, error) {
	if strings.Contains(path, "/") || strings.Contains(path, "\\") {
		return 0, errors.New("file must be a file name (not a path)")
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return 0, errors.New("unable to open file")
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	dataConn, err := c.getDataConnection()
	if err != nil {
		return 0, errors.New("unable to open data connection with server")
	}

	_, err = c.sendCommand("TYPE I")
	if err != nil {
		return 0, errors.New("unable to convert the data connection to binary")
	}

	data, err := c.sendCommand("STOR " + path)
	if err != nil {
		_ = dataConn.Close()
		return 0, errors.New("unable to initiate file storing")
	}

	status, err := strconv.Atoi(string(data[0:3]))
	if err != nil {
		_ = c.dataConn.Close()
		return 0, errors.New("unable to read status code")
	}

	if status != statusStartingTransfer {
		_ = dataConn.Close()
		return 0, errors.New("unable to start transfer")
	}

	// Check the status code here
	written, err := io.Copy(dataConn, file)
	_ = dataConn.Close()
	if err != nil {
		return 0, errors.New("unable to copy data to server")
	}

	_, _ = c.readData()
	return int(written), nil
}
