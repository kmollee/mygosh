package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kmollee/mygosh/shell"
	"github.com/kmollee/mygosh/shell/commands"
)

const (
	shellLoginText = "Welcome to Gosh"
	version        = "v0.1.0"
)

func main() {
	s, err := shell.NewShell(os.Stdin, shellLoginText+version)
	if err != nil {
		log.Fatal(err)
	}
	s.Init()
	for {
		s.Prompt()
		cmd, err := s.Read()
		if err != nil {
			log.Fatalf("could not read: %v\n", err)
		}
		err = commands.Run(cmd...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not execute command %s: %v\n", cmd, err)
		}
	}
}
