package commands

import (
	"fmt"
	"ftpclient/ftpclient"
)

var dirCommand = &Cmd{
	name:               "dir",
	description:        "Lists all files and folders in the current directory",
	requiresConnection: true,
	execute: func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd) {
		dir, err := client.GetDirectoryList()
		if err != nil {
			fmt.Printf("An error has occured: %s\n", err.Error())
			return
		}

		fmt.Println(dir)
	},
}
