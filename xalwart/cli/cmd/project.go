package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/cli/utils"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core/components"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/templates"
	"github.com/iancoleman/strcase"
)

var (
	projectName               string
	projectRootPath           string
	projectOverwrite               = false
	projectSecretKeyLength    uint = 50
	projectUsedStandardORM         = true
	projectUsedStandardServer      = true
)

var projectCommand = utils.NewComponentCommandBuilder().
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

func buildProjectComponent() (core.Component, error) {
	header, err := components.NewHeaderComponent(templates.CopyrightNoticesTemplateBox)
	if err != nil {
		return nil, err
	}

	secretKey, err := core.RandomString(projectSecretKeyLength)
	if err != nil {
		return nil, err
	}

	return components.NewProjectComponent(
		*header,
		templates.ProjectTemplateBox,
		secretKey,
		projectName,
		projectRootPath,
		projectUsedStandardORM,
		projectUsedStandardServer,
	), nil
}

func projectSuccess(component core.Component) string {
	project := component.(*components.ProjectComponent)
	nameLowercase := strcase.ToSnake(project.ProjectName())
	return fmt.Sprintf(
		`Success.

Examine 'README.md' in the project root directory for more info.

Check the application by running it in docker container:

  sudo docker build -t %s:latest .
  docker run -p 8000:8000 %s:latest ./application start-server --bind 0.0.0.0:8000 --workers=5

You can read CMake configuration and build process logs in /var/log/app after building the image:

  docker run %s:latest cat /var/log/app/configure.log
  docker run %s:latest cat /var/log/app/build.log
`, nameLowercase, nameLowercase, nameLowercase, nameLowercase,
	)
}
