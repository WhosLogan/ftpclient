package main

type Cmd struct {
	name    string
	help    string
	args    []string
	execute func(args []string)
}

func executeCmd() {

}
