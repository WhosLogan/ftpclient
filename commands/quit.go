package commands

import (
	"ftpclient/ftpclient"
	"os"
)

var quitCommand = &Cmd{
	name:        "quit",
	description: "Closes all connections and exits the FTP client",
	execute: func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd) {
		client.Close()
		os.Exit(0)
	},
}
