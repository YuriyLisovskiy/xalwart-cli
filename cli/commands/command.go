package commands

import (
	"github.com/YuriyLisovskiy/xalwart-cli/generator"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	commandName string
	commandRootPath string
	commandCustomFileName string
)

var commandCommand = &cobra.Command{
	Use:   "command",
	Short: "Create new command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(commandName) == 0 {
			log.Fatal("Command name should be set")
		}

		unit, err := generator.NewCommandUnit(commandName, commandRootPath, commandCustomFileName)
		if err != nil {
			log.Fatal(err)
		}

		must(generator.GenerateUnit(unit))
	},
}

func init() {
	commandCommand.Flags().StringVarP(
		&commandName, "name", "n", "", "name of new command (example: Dump)",
	)
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	commandCommand.Flags().StringVarP(
		&commandRootPath, "root", "r", currentDirectory, "root path of new command",
	)
	commandCommand.Flags().StringVarP(
		&commandCustomFileName, "file", "f", "command", "custom file name for a new command",
	)
}
