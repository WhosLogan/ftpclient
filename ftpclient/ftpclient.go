package ftpclient

type FtpClient struct {
	Username string
	Password string
	Server   string
}

func NewClient() *FtpClient {
	return &FtpClient{
		Username: "anonymous",
	}
}
