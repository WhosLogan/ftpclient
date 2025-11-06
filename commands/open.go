package commands

import (
	"fmt"
	"ftpclient/ftpclient"
)

var openCommand = &Cmd{
	name:        "open",
	description: "Opens a connection to an ftp server",
	argName:     "server",
	execute: func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd) {
		if arg == "" {
			fmt.Println("Please specify a target server")
			return
		}

		client.Server = arg
		err := client.Connect()
		if err != nil {
			fmt.Println("Error connecting to server:", err)
		}
	},
}
