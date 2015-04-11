package main

import (
	"os"
	"log"
	"github.com/mitchellh/cli"
	"github.com/kellydunn/sfdc/commands"
)

func main() {
	c := cli.NewCLI("sfdc", "0.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"15to18": commands.NewExpandIdCommand,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)	
}

