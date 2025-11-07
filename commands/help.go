package commands

import (
	"fmt"
	"ftpclient/ftpclient"
	"strings"
)

var helpCommand = &Cmd{
	name:        "help",
	description: "Provides a help menu",
	execute: func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd) {
		builder := strings.Builder{}
		for _, command := range *commands {
			builder.WriteString(command.name)
			if command.argName != "" {
				builder.WriteString(" <")
				builder.WriteString(command.argName)
				builder.WriteString(">")
			}
			builder.WriteString(": ")
			builder.WriteString(command.description)
			builder.WriteRune('\n')
		}
		fmt.Print(builder.String())
	},
}
