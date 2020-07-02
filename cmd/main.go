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
	fmt.Println("  " + commands.NewAppCmdDescription)
	fmt.Println("  " + commands.NewCommandCmdDescription)
	fmt.Println("  " + commands.NewLibraryCmdDescription)
	fmt.Println("  " + commands.NewMiddlewareCmdDescription)
	fmt.Println("  " + commands.NewProjectCmdDescription)
	println()

	commands.InitInstallCmd()
	commands.InitNewAppCmd()
	commands.InitNewCommandCmd()
	commands.InitNewLibraryCmd()
	commands.InitNewMiddlewareCmd()
	commands.InitNewProjectCmd()

	commands.InstallCmd.Usage()
	println()
	commands.NewAppCmd.Usage()
	println()
	commands.NewCommandCmd.Usage()
	println()
	commands.NewLibraryCmd.Usage()
	println()
	commands.NewMiddlewareCmd.Usage()
	println()
	commands.NewProjectCmd.Usage()
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
			commands.InitInstallCmd()
			if commands.InstallCmd.Parse(os.Args[2:]) != nil {
				commands.InstallCmd.Usage()
			} else {
				err = cmd.InstallFramework()
			}
		case commands.NewAppCmd.Name():
			commands.InitNewAppCmd()
			if commands.NewAppCmd.Parse(os.Args[2:]) != nil {
				commands.NewAppCmd.Usage()
			} else {
				err = cmd.CreateApp()
			}
		case commands.NewCommandCmd.Name():
			commands.InitNewCommandCmd()
			if commands.NewCommandCmd.Parse(os.Args[2:]) != nil {
				commands.NewCommandCmd.Usage()
			} else {
				err = cmd.CreateCommand()
			}
		case commands.NewLibraryCmd.Name():
			commands.InitNewLibraryCmd()
			if commands.NewLibraryCmd.Parse(os.Args[2:]) != nil {
				commands.NewLibraryCmd.Usage()
			} else {
				err = cmd.CreateLibrary()
			}
		case commands.NewMiddlewareCmd.Name():
			commands.InitNewMiddlewareCmd()
			if commands.NewMiddlewareCmd.Parse(os.Args[2:]) != nil {
				commands.NewMiddlewareCmd.Usage()
			} else {
				err = cmd.CreateMiddleware()
			}
		case commands.NewProjectCmd.Name():
			commands.InitNewProjectCmd()
			if commands.NewProjectCmd.Parse(os.Args[2:]) != nil {
				commands.NewProjectCmd.Usage()
			} else {
				err = cmd.CreateProject()
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
			println()
			panic(err)
			// TODO: uncomment in release version
			// fmt.Println("\nError: " + err.Error())
		}
	}
}
