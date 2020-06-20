package main

import (
	"fmt"
	"os"
	"strings"
	"xalwart-cli/cmd/commands"
	"xalwart-cli/config"
)

func usage() {
	fmt.Println(
		"Welcome to CLI tool for managing projects written in '" +
		config.FrameworkName + "' framework.",
	)
	println()
	fmt.Println("General usage:")
	fmt.Println("  xalwart <command> [arguments]")
	println()
	fmt.Println(
		"Use \"" + strings.ToLower(config.FrameworkName) +
		" <command> --help\" for more information about a command.",
	)
	println()
	fmt.Println("The commands are:")
	fmt.Println("  " + commands.NewProjectCmd.Name() + ":\tcreates a new project based on cmake lists")
	fmt.Println(
		"  " + commands.NewAppCmd.Name() +
		":\tadds a new application to existing '" + config.FrameworkName + "' application",
	)
	println()
	commands.NewProjectCmd.Usage()
	println()
	commands.NewAppCmd.Usage()
	println()
}

func main() {
	if len(os.Args) < 2 {
		usage()
	} else {
		switch os.Args[1] {
		case commands.NewProjectCmd.Name():
			if commands.NewProjectCmd.Parse(os.Args[2:]) != nil {
				commands.NewProjectCmd.Usage()
			} else {
				err := commands.CreateProject()
				if err != nil {
					panic(err)
					// TODO: uncomment in release version
					// fmt.Println("Error: " + err.Error())
				}
			}
		case commands.NewAppCmd.Name():
			if commands.NewAppCmd.Parse(os.Args[2:]) != nil {
				commands.NewAppCmd.Usage()
			} else {
				fmt.Println("Not implemented.")
			}
		default:
			usage()
		}
	}
}
