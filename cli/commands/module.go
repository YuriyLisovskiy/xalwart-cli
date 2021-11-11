package commands

import (
	"github.com/YuriyLisovskiy/xalwart-cli/generator"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	moduleName string
	moduleRootPath string
	moduleCustomFileName string
)

var moduleCommand = &cobra.Command{
	Use:   "module",
	Short: "Create new module",
	Run: func(cmd *cobra.Command, args []string) {
		if len(moduleName) == 0 {
			log.Fatal("Module name should be set")
		}

		unit, err := generator.NewModuleUnit(moduleName, moduleRootPath, moduleCustomFileName)
		if err != nil {
			log.Fatal(err)
		}

		must(generator.GenerateUnit(unit))
	},
}

func init() {
	moduleCommand.Flags().StringVarP(
		&moduleName, "name", "n", "", "name of new module (example: Main)",
	)
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	moduleCommand.Flags().StringVarP(
		&moduleRootPath, "root", "r", currentDirectory, "root path of new module",
	)
	moduleCommand.Flags().StringVarP(
		&moduleCustomFileName, "file", "f", "module", "custom file name of new module",
	)
}
