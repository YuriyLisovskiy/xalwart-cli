package commands

import (
	"flag"
	"github.com/gobuffalo/packr/v2"
	"os"
	"os/user"
	"time"
	"xalwart-cli/config"
	"xalwart-cli/generator"
)

const newAppCmdName = "new-app"

var (
	NewAppCmd = flag.NewFlagSet(newAppCmdName, flag.ExitOnError)

	naRootFlag = NewAppCmd.String("root", "", "Root of a project")
	naNameFlag = NewAppCmd.String("name", "", "Name of a new application")
	naFrameworkVersionFlag = NewAppCmd.String(
		"fw-version",
		"latest",
		"A version of '" + config.FrameworkName + "' framework used",
	)
)

func CreateApp() error {
	if len(*naRootFlag) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		*naRootFlag = cwd
	}

	usr, err := user.Current()
	if err != nil {
		return err
	}

	cfg := config.App{
		Year:                      time.Now().Year(),
		Username:                  usr.Name,
		FrameworkName:             config.FrameworkName,
		FrameworkNamespace:        config.FrameworkNamespace,
		AppName:                   *naNameFlag,
		ProjectRoot:               *naRootFlag,
		Templates:                 packr.New("App Templates Box", "../templates/app"),
	}

	_, cfg.FrameworkVersionSubDir, _ = getFWVersionAndSubDir(
		*naFrameworkVersionFlag, false,
	)

	g := generator.Generator{}
	err = g.NewApp(&cfg)
	if err != nil {
		return err
	}

	return nil
}
