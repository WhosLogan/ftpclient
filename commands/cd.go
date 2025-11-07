package commands

import (
	"fmt"
	"ftpclient/ftpclient"
)

var cdCommand = &Cmd{
	name:               "cd",
	description:        "Changes the current working directory",
	argName:            "directory",
	requiresConnection: true,
	execute: func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd) {
		if !client.IsConnected() {
			fmt.Println("You are not connected to a server")
		}

		err := client.ChangeDirectory(arg)
		if err != nil {
			fmt.Println("Unable to change directory")
		}
	},
}
