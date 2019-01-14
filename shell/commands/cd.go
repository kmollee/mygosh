package commands

import (
	"log"
	"os"
	"strings"
)

// Cd changes the shell's directory
func Cd(argv ...string) error {
	log.Printf("cd argv: %v", argv)
	// BUG: cd - is not work
	if len(argv) == 1 {
		os.Chdir(os.Getenv("HOME"))
	} else if argv[1][0:1] == "/" {
		os.Chdir(argv[1])
	} else if argv[1][0:1] == "~" {
		os.Chdir(os.Getenv("HOME") + strings.Join(argv[1:], ""))
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		os.Chdir(wd + "/" + argv[1])
	}
	return nil
}
