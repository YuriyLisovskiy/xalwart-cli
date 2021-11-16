package commands

import (
	"log"
	"os"

	"github.com/YuriyLisovskiy/xalwart-cli/cli/commands/util"
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
	"github.com/spf13/cobra"
)

var (
	projectName               string
	projectRootPath           string
	projectOverwrite               = false
	projectSecretKeyLength    uint = 50
	projectUsedStandardORM         = true
	projectUsedStandardServer      = true
)

var projectCommand *cobra.Command

func makeProjectCommand() *cobra.Command {
	builder := util.CommandBuilder{}
	builder.SetName("project")
	builder.SetShortDescription("Create new project")
	builder.SetNameValidator(
		func() {
			if len(projectName) == 0 {
				log.Fatalf("project name should be set")
			}
		},
	)
	builder.SetComponentBuilder(
		func() (core.Component, error) {
			return components.NewProjectComponent(
				projectName,
				projectRootPath,
				projectSecretKeyLength,
				projectUsedStandardORM,
				projectUsedStandardServer,
			)
		},
	)
	return builder.Command(&projectOverwrite)
}

func init() {
	projectCommand = makeProjectCommand()
	flags := projectCommand.Flags()
	flags.StringVarP(
		&projectName, "name", "n", "", "name of a new project",
	)
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flags.StringVarP(
		&projectRootPath,
		"root",
		"r",
		currentDirectory,
		"root path for a new project",
	)
	flags.BoolVarP(
		&projectOverwrite,
		"overwrite",
		"o",
		projectOverwrite,
		"overwrite files if exist",
	)
	flags.UintVarP(
		&projectSecretKeyLength,
		"key-length",
		"k",
		projectSecretKeyLength,
		"length of secret key to be generated in settings",
	)
	flags.BoolVarP(
		&projectUsedStandardORM,
		"use-orm",
		"d",
		projectUsedStandardORM,
		"use standard ORM, provided by framework",
	)
	flags.BoolVarP(
		&projectUsedStandardServer,
		"use-server",
		"s",
		projectUsedStandardServer,
		"use standard web server, provided by framework",
	)
}
