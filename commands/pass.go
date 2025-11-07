package commands

import (
	"ftpclient/ftpclient"
)

var passCommand = &Cmd{
	name:        "pass",
	description: "Sets a password for the FTP client",
	argName:     "password",
	execute: func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd) {
		client.Password = arg
	},
}
