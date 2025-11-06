package ftpclient

func (c *FtpClient) GetDirectoryList() (string, error) {
	err := c.connectData()
	if err != nil {
		return "", err
	}

	_, err = c.sendCommand("LIST")
	if err != nil {
		return "", err
	}

	return string(<-c.dataChan), nil
}
