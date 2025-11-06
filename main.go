package main

import (
	"bufio"
	"fmt"
	"ftpclient/commands"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Welcome to Logan & Harrison's FTP Client")
	fmt.Println("Type 'help' for a list of commands.")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	for {
		select {
		case <-sigs:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Print("> ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			commands.ExecuteCmd(text)
		}
	}
}
