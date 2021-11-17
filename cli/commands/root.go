package commands

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/YuriyLisovskiy/xalwart-cli/cli/commands/add"
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:     core.AppName,
	Version: core.AppVersion,
	Short: fmt.Sprintf(
		"%s is project and component generator tool", strcase.ToCamel(core.AppName),
	),
	Long: fmt.Sprintf(
		`%s CLI tool for generating project (and it's components) written in %s framework.
Complete documentation is available at %s`,
		strcase.ToCamel(core.AppName), strcase.ToCamel(core.FrameworkName), core.CLIAppDocumentationLink,
	),
}

func init() {
	rootCommand.InitDefaultVersionFlag()
	versionTemplate := strings.TrimRight(rootCommand.VersionTemplate(), "\n")
	versionTemplate = fmt.Sprintf("%s (%s, %s)\n", versionTemplate, runtime.GOOS, runtime.GOARCH)
	rootCommand.SetVersionTemplate(versionTemplate)

	rootCommand.AddCommand(add.RootCommand)
	rootCommand.AddCommand(projectCommand)
	rootCommand.AddCommand(versionCommand)
}

func Execute() {
	err := rootCommand.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
