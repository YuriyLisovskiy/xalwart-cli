package commands

import (
	"fmt"
	"github.com/YuriyLisovskiy/xalwart-cli/config"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   config.AppName,
	Short: fmt.Sprintf("%s is project and component generator tool", strcase.ToCamel(config.AppName)),
	Long: fmt.Sprintf(`%s CLI tool for generating project (and it's components) written in Xalwart framework.
Complete documentation is available at https://github.com/YuriyLisovskiy/xalwart-cli`, strcase.ToCamel(config.AppName)),
}

func init() {
	rootCmd.AddCommand(addCommand)
	rootCmd.AddCommand(projectCommand)
}

func Execute() {
	must(rootCmd.Execute())
}
