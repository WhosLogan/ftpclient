package main

type Cmd struct {
	name    string
	help    string
	argName string
	execute func(args string)
}

func executeCmd() {

}
