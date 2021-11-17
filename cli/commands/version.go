package commands

import (
	"fmt"
	"text/template"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Print version for %s", core.AppName),
	RunE: runVersion,
}

func runVersion(cmd *cobra.Command, args []string) error {
	t := template.New("top")
	t, err := t.Parse(rootCommand.VersionTemplate())
	if err != nil {
		return err
	}

	return t.Execute(cmd.OutOrStdout(), rootCommand)
}
