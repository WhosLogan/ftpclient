package commands

import (
	"fmt"
	"ftpclient/ftpclient"
	"os"
	"strings"
)

var getCommand = &Cmd{
	name:               "get",
	description:        "Download's a file from the FTP server",
	argName:            "file",
	requiresConnection: true,
	execute: func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd) {
		if arg == "" {
			fmt.Println("Please provide a file name")
			return
		}

		file, err := client.GetFile(arg)
		if err != nil {
			fmt.Println("Unable to download file:", err.Error())
			return
		}

		err = os.WriteFile("./"+strings.ReplaceAll(strings.ReplaceAll(arg, "/", "-"),
			"\\", "-"), file, 0666)
		if err != nil {
			fmt.Println("Unable to write downloaded file:", err.Error())
		}
	},
}
