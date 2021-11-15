package commands

import (
	"github.com/YuriyLisovskiy/xalwart-cli/generator"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	controllerName string
	controllerRootPath string
	controllerCustomFileName string
)

var controllerCommand = &cobra.Command{
	Use:   "controller",
	Short: "Create new controller",
	Run: func(cmd *cobra.Command, args []string) {
		if len(controllerName) == 0 {
			log.Fatal("Controller name should be set")
		}

		unit, err := generator.NewControllerUnit(
			controllerName,
			controllerRootPath,
			controllerCustomFileName,
		)
		if err != nil {
			log.Fatal(err)
		}

		must(generator.GenerateUnit(unit))
	},
}

func init() {
	controllerCommand.Flags().StringVarP(
		&controllerName,
		"name",
		"n",
		"",
		"name of new controller (example: Index)",
	)
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	controllerCommand.Flags().StringVarP(
		&controllerRootPath,
		"root",
		"r",
		currentDirectory,
		"root path of new controller",
	)
	controllerCommand.Flags().StringVarP(
		&controllerCustomFileName,
		"file",
		"f",
		"{lower_case_name}_controller",
		"custom file name of new controller",
	)
}
