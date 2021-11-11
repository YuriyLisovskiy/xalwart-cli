package commands

import (
	"github.com/spf13/cobra"
	"log"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Create new project component",
}

func init() {
	addCommand.AddCommand(controllerCommand)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
