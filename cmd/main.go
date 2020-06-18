package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'new-project' sub-command")
		os.Exit(1)
	} else {
		switch os.Args[1] {
		case newProjectCmd.Name():
			if newProjectCmd.Parse(os.Args[2:]) != nil {
				newProjectCmd.Usage()
			} else {
				err := createProject()
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
