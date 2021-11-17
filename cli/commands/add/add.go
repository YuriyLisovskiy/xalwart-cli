package add

import (
	"fmt"
	"log"
	"os"

	"github.com/YuriyLisovskiy/xalwart-cli/cli/util"
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var RootCommand = &cobra.Command{
	Use:   "add",
	Short: "Create new component of project",
}

func init() {
	RootCommand.AddCommand(commandCommand)
	RootCommand.AddCommand(controllerCommand)
	RootCommand.AddCommand(middlewareCommand)
	RootCommand.AddCommand(migrationCommand)
	RootCommand.AddCommand(modelCommand)
	RootCommand.AddCommand(moduleCommand)
}

func makeCommand(
	command, longDescription string,
	componentBuilder func() (core.Component, error),
	postCreateMessageBuilder func(core.Component) string,
) *cobra.Command {
	builder := util.CommandBuilder{}
	builder.SetName(command)
	builder.SetShortDescription("Create new " + command)
	builder.SetLongDescription(longDescription)
	builder.SetNameValidator(
		func() {
			if len(componentName) == 0 {
				log.Fatalf("%s name should be set", command)
			}
		},
	)
	builder.SetComponentBuilder(componentBuilder)
	builder.SetPostCreateMessageBuilder(postCreateMessageBuilder)
	return builder.Command(&overwrite)
}

func addCommonFlags(component string, flags *pflag.FlagSet) {
	flags.StringVarP(
		&componentName,
		"name",
		"n",
		"",
		fmt.Sprintf("name of new %s", component),
	)
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flags.StringVarP(
		&rootPath,
		"root",
		"r",
		currentDirectory,
		fmt.Sprintf("root path of new %s", component),
	)
	flags.StringVarP(
		&componentCustomFileName,
		"file",
		"f",
		"",
		fmt.Sprintf("custom file name of new %s", component),
	)
	flags.BoolVarP(
		&overwrite,
		"overwrite",
		"o",
		overwrite,
		"overwrite files if exist",
	)
}
