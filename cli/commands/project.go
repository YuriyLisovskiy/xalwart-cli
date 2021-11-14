package commands

import (
	"github.com/YuriyLisovskiy/xalwart-cli/generator"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	projectName               string
	projectRootPath           string
	projectSecretKeyLength    uint = 50
	projectUsedStandardORM         = true
	projectUsedStandardServer      = true
)

var projectCommand = &cobra.Command{
	Use:   "project",
	Short: "Create new project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(projectName) == 0 {
			log.Fatal("Project name should be set")
		}

		unit, err := generator.NewProjectUnit(
			projectName,
			projectRootPath,
			projectSecretKeyLength,
			projectUsedStandardORM,
			projectUsedStandardServer,
		)
		if err != nil {
			log.Fatal(err)
		}

		must(generator.GenerateUnit(unit))
	},
}

func init() {
	projectCommand.Flags().StringVarP(
		&projectName, "name", "n", "", "name of a new project",
	)
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	projectCommand.Flags().StringVarP(
		&projectRootPath, "root", "r", currentDirectory, "root path for a new project",
	)
	projectCommand.Flags().UintVarP(
		&projectSecretKeyLength,
		"key-length",
		"k",
		projectSecretKeyLength,
		"length of secret key to be generated in settings",
	)
	projectCommand.Flags().BoolVarP(
		&projectUsedStandardORM,
		"use-orm",
		"o",
		projectUsedStandardORM,
		"use standard ORM, provided by framework",
	)
	projectCommand.Flags().BoolVarP(
		&projectUsedStandardServer,
		"use-server",
		"s",
		projectUsedStandardServer,
		"use standard web server, provided by framework",
	)
}
