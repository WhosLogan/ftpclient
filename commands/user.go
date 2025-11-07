package commands

import (
	"ftpclient/ftpclient"
)

var userCommand = &Cmd{
	name:        "user",
	description: "Sets a username for the FTP client",
	argName:     "username",
	execute: func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd) {
		client.Username = arg
	},
}
