package main

import (
	"fmt"
	"os"
	"xalwart-cli/cmd/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'new-project' sub-command")
		os.Exit(1)
	} else {
		switch os.Args[1] {
		case commands.NewProjectCmd.Name():
			if commands.NewProjectCmd.Parse(os.Args[2:]) != nil {
				commands.NewProjectCmd.Usage()
			} else {
				err := commands.CreateProject()
				if err != nil {
					panic(err)
				}
			}
		default:
			fmt.Println("expected 'new-project' sub-command")
			os.Exit(1)
		}
	}
}
