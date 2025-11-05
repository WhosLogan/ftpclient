package commands

import (
	"fmt"
	"strings"
)

var helpCommand = &Cmd{
	name:        "help",
	description: "Provides a help menu",
	execute: func(arg string, commands *[]*Cmd) {
		for _, command := range *commands {
			builder := strings.Builder{}
			builder.WriteString(command.name)
			if command.argName != "" {
				builder.WriteString(" <")
				builder.WriteString(command.argName)
				builder.WriteString(">")
			}
			builder.WriteString(": ")
			builder.WriteString(command.description)
			builder.WriteRune('\n')
			fmt.Print(builder.String())
		}
	},
}
