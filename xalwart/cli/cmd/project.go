package cmd

import (
	"errors"
	"log"
	"os"

	utils2 "github.com/YuriyLisovskiy/xalwart-cli/xalwart/cli/utils"
	core2 "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
	components2 "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core/components"
)

var (
	projectName               string
	projectRootPath           string
	projectOverwrite               = false
	projectSecretKeyLength    uint = 50
	projectUsedStandardORM         = true
	projectUsedStandardServer      = true
)

var projectCommand = utils2.NewComponentCommandBuilder().
	SetName("project").
	SetShortDescription("Create new project").
	SetNameValidator(validateProjectName).
	SetComponentBuilder(buildProjectComponent).
	SetPostRunMessageBuilder(projectSuccess).
	Command(&projectOverwrite)

func init() {
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
		"O",
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

func validateProjectName() error {
	if len(projectName) == 0 {
		return errors.New("project name should be set")
	}

	return nil
}

func buildProjectComponent() (core2.Component, error) {
	header, err := components2.NewHeaderComponent(utils2.GetCopyrightNoticesTemplateBox())
	if err != nil {
		return nil, err
	}

	secretKey, err := core2.RandomString(projectSecretKeyLength)
	if err != nil {
		return nil, err
	}

	return components2.NewProjectComponent(
		*header,
		utils2.GetProjectTemplateBox(),
		secretKey,
		projectName,
		projectRootPath,
		projectUsedStandardORM,
		projectUsedStandardServer,
	), nil
}

func projectSuccess(core2.Component) string {
	return `Success.

Examine 'README.md' in the project root directory.
`
}
