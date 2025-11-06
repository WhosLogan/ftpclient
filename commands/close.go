package commands

import (
	"ftpclient/ftpclient"
)

var closeCommand = &Cmd{
	name:        "close",
	description: "Close any active connection to the FTP server",
	execute: func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd) {
		client.Close()
	},
}
