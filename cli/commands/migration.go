package commands

import (
	"github.com/YuriyLisovskiy/xalwart-cli/generator"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	migrationName           string
	migrationRootPath       string
	migrationCustomFileName string
	migrationIsInitial      = false
)

var migrationCommand = &cobra.Command{
	Use:   "migration",
	Short: "Create new migration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(migrationName) == 0 {
			log.Fatal("Migration name should be set")
		}

		unit, err := generator.NewMigrationUnit(
			migrationName,
			migrationRootPath,
			migrationCustomFileName,
			migrationIsInitial,
		)
		if err != nil {
			log.Fatal(err)
		}

		must(generator.GenerateUnit(unit))
	},
}

func init() {
	migrationCommand.Flags().StringVarP(
		&migrationName,
		"name",
		"n",
		"",
		"name of new migration (recommended: Migration{number}_{ShortDescription}) (example: Migration001_Initial)",
	)
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	migrationCommand.Flags().StringVarP(
		&migrationRootPath,
		"root",
		"r",
		currentDirectory,
		"root path of new migration",
	)
	migrationCommand.Flags().StringVarP(
		&migrationCustomFileName,
		"file",
		"f",
		"",
		"custom file name of new migration ({number}_{ShortDescription} if used recommended name pattern)",
	)
	migrationCommand.Flags().BoolVarP(
		&migrationIsInitial,
		"initial",
		"i",
		migrationIsInitial,
		"mark migration as initial one",
	)
}
