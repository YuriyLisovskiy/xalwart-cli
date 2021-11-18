package add

import (
	"errors"
	"fmt"
	"log"
	"os"

	utils2 "github.com/YuriyLisovskiy/xalwart-cli/xalwart/cli/utils"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core/components"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var RootCommand = &cobra.Command{
	Use:   "add",
	Short: "Create new component for existing project",
}

func init() {
	RootCommand.AddCommand(commandCommand)
	RootCommand.AddCommand(controllerCommand)
	RootCommand.AddCommand(middlewareCommand)
	RootCommand.AddCommand(migrationCommand)
	RootCommand.AddCommand(modelCommand)
	RootCommand.AddCommand(moduleCommand)
}

func initDefaultFlags(component string, flags *pflag.FlagSet) {
	flags.StringVarP(&nameVar, "name", "n", "", fmt.Sprintf("name of new %s", component))
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flags.StringVarP(&rootPathVar, "root", "r", currentDirectory, fmt.Sprintf("root path of new %s", component))
	flags.StringVarP(&customFileNameVar, "file", "f", "", fmt.Sprintf("custom file name of new %s", component))
	flags.BoolVarP(&overwriteVar, "overwrite", "o", overwriteVar, "overwrite files if exist")
}

func getComponentCommandBuilder(name, longDescription string) *utils2.ComponentCommandBuilder {
	builder := &utils2.ComponentCommandBuilder{}
	builder.SetName(name)
	builder.SetShortDescription(fmt.Sprintf("Create new %s component.", name))
	builder.SetLongDescription(longDescription)
	builder.SetNameValidator(
		func() error {
			if len(nameVar) == 0 {
				return errors.New(fmt.Sprintf("%s name should be set", name))
			}

			return nil
		},
	)
	return builder
}

func getDefaultHeader() (core.Header, error) {
	return components.NewHeaderComponent(utils2.GetCopyrightNoticesTemplateBox())
}
