package commands

import (
	"github.com/spf13/cobra"
	"log"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Create new component of project",
}

func init() {
	addCommand.AddCommand(commandCommand)
	addCommand.AddCommand(controllerCommand)
	addCommand.AddCommand(migrationCommand)
	addCommand.AddCommand(moduleCommand)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
