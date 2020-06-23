package main

import (
	"fmt"
	"os"
	"strings"
	"xalwart-cli/cmd/commands"
	"xalwart-cli/config"
)

func usage(printWelcome bool) {
	if printWelcome {
		fmt.Println(
			"Welcome to CLI tool for managing projects written in '" +
			config.FrameworkName + "' framework.",
		)
		println()
	}

	fmt.Println("General usage:")
	fmt.Println(
		"  " + config.FrameworkName + " [--version] [--help]\n  " +
		strings.Repeat(" ", len(config.FrameworkName)) +" <command> [arguments]",
	)
	println()
	fmt.Println(
		"Use \"" + strings.ToLower(config.FrameworkName) +
		" <command> --help\" for more information about a command.",
	)
	println()
	fmt.Println("The commands are:")
	fmt.Println("  " + commands.InstallCmdDescription)
	fmt.Println("  " + commands.NewProjectCmd.Name() + ":\tcreates a new project based on cmake lists")
	fmt.Println(
		"  " + commands.NewAppCmd.Name() +
		":\tadds a new application to existing '" + config.FrameworkName + "' application",
	)
	fmt.Println("  " + commands.NewLibraryCmd.Name() + ":\tcreates a new library for template engine")
	println()
	commands.InstallCmd.Usage()
	println()
	commands.NewProjectCmd.Usage()
	println()
	commands.NewAppCmd.Usage()
	println()
	commands.NewLibraryCmd.Usage()
	println()
}

func main() {
	if len(os.Args) < 2 {
		usage(true)
	} else {
		var err error
		cmd := commands.Cmd{}
		switch os.Args[1] {
		case commands.InstallCmd.Name():
			if commands.InstallCmd.Parse(os.Args[2:]) != nil {
				commands.InstallCmd.Usage()
			} else {
				err = cmd.InstallFramework()
			}
		case commands.NewProjectCmd.Name():
			if commands.NewProjectCmd.Parse(os.Args[2:]) != nil {
				commands.NewProjectCmd.Usage()
			} else {
				err = cmd.CreateProject()
			}
		case commands.NewAppCmd.Name():
			if commands.NewAppCmd.Parse(os.Args[2:]) != nil {
				commands.NewAppCmd.Usage()
			} else {
				err = cmd.CreateApp()
			}
		case commands.NewLibraryCmd.Name():
			if commands.NewLibraryCmd.Parse(os.Args[2:]) != nil {
				commands.NewLibraryCmd.Usage()
			} else {
				err = cmd.CreateLibrary()
			}
		case "-h", "--help", "help":
			usage(false)
		case "-v", "--version", "version":
			fmt.Println(config.FrameworkName + " version " + config.AppVersion)
		default:
			fmt.Println(
				config.FrameworkName + ": '" + os.Args[1] + "' is not a " +
				config.FrameworkName + " command. See '" + config.FrameworkName + " --help'.",
			)
		}

		if err != nil {
			panic(err)
			// TODO: uncomment in release version
			// fmt.Println("Error: " + err.Error())
		}
	}
}
