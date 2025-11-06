package commands

import (
	"fmt"
	"ftpclient/ftpclient"
)

var dirCommand = &Cmd{
	name:        "dir",
	description: "Lists all files and folders in the current directory",
	execute: func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd) {
		if !client.IsConnected() {
			fmt.Println("You are not connected to a server")
			return
		}

		dir, err := client.GetDirectoryList()
		if err != nil {
			fmt.Printf("An error has occured: %s\n", err.Error())
			return
		}

		fmt.Println(dir)
	},
}
