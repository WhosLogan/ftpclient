package commands

import (
	"fmt"
	"ftpclient/ftpclient"
	"strings"
)

type Cmd struct {
	name        string
	description string
	argName     string
	execute     func(client *ftpclient.FtpClient, arg string, commands *[]*Cmd)
}

var commands = []*Cmd{
	helpCommand,
	userCommand,
	passCommand,
	quitCommand,
	closeCommand,
	dirCommand,
	openCommand,
	cdCommand,
	getCommand,
	putCommand,
}

func ExecuteCmd(client *ftpclient.FtpClient, text string) {
	text = strings.TrimSpace(text)

	split := strings.Split(text, " ")
	if len(split) < 1 || split[0] == "" {
		return
	}

	split[0] = strings.ToLower(split[0])

	for _, cmd := range commands {
		if split[0] == cmd.name {
			if cmd.argName == "" && len(split) > 1 {
				fmt.Println("Invalid command syntax")
				return
			}

			cmd.execute(client, strings.Join(split[1:], " "), &commands)
			return
		}
	}

	fmt.Println("Unknown command")
}
