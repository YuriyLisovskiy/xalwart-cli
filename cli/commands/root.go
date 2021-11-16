package commands

import (
	"fmt"
	"log"

	"github.com/YuriyLisovskiy/xalwart-cli/cli/commands/add"
	"github.com/YuriyLisovskiy/xalwart-cli/config"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   config.AppName,
	Short: fmt.Sprintf(
		"%s is project and component generator tool", strcase.ToCamel(config.AppName),
	),
	Long: fmt.Sprintf(
		`%s CLI tool for generating project (and it's components) written in Xalwart framework.
Complete documentation is available at https://github.com/YuriyLisovskiy/xalwart-cli`,
		strcase.ToCamel(config.AppName),
	),
}

func init() {
	rootCommand.AddCommand(add.RootCommand)
	rootCommand.AddCommand(projectCommand)
}

func Execute() {
	err := rootCommand.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
