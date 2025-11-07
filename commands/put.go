package commands

import (
	"fmt"
	"ftpclient/ftpclient"
)

var putCommand = &Cmd{
	name:               "put",
	description:        "Uploads a file to the FTP server",
	argName:            "file",
	requiresConnection: true,
	execute: func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd) {
		if !client.IsConnected() {
			fmt.Println("Not connected to server")
			return
		}

		written, err := client.SendFile(arg)
		if err != nil {
			fmt.Println("Error sending file:", err)
			return
		}

		fmt.Printf("%d bytes uploaded to server\n", written)
	},
}
