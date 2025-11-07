package commands

import (
	"fmt"
	"ftpclient/ftpclient"
)

var cdCommand = &Cmd{
	name:        "cd",
	description: "Changes the current working directory",
	argName:     "directory",
	execute: func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd) {
		err := client.ChangeDirectory(arg)
		if err != nil {
			fmt.Println("Unable to change directory")
		}
	},
}
