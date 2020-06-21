package commands

import (
	"errors"
	"flag"
	"github.com/gobuffalo/packr/v2"
	"os"
	"os/user"
	"path"
	"strings"
	"time"
	"xalwart-cli/config"
	"xalwart-cli/generator"
)

const newAppCmdName = "new-app"

var (
	NewAppCmd = flag.NewFlagSet(newAppCmdName, flag.ExitOnError)

	naNameFlag = NewAppCmd.String("name", "", "Name of a new application")
)

func CreateApp() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	usr, err := user.Current()
	if err != nil {
		return err
	}

	unitCfg := generator.ProjectUnit{
		Year:                      time.Now().Year(),
		Username:                  usr.Name,
		FrameworkName:             config.FrameworkName,
		FrameworkNamespace:        config.FrameworkNamespace,
		Name:                      *naNameFlag,
		ProjectRoot:               cwd,
		Templates:                 packr.New("App Templates Box", "../templates/app"),
		Customize: func(pu *generator.ProjectUnit) {
			if !strings.HasSuffix(strings.ToLower(pu.Name), "_app") {
				pu.Name += "_app"
			}

			pu.Root = path.Join(pu.ProjectRoot, pu.Name)
		},
	}

	meta, err := loadMeta(unitCfg.ProjectRoot)
	if err != nil {
		return err
	}

	unitCfg.ProjectName = meta.ProjectName

	_, unitCfg.FrameworkVersionSubDir, _ = getFWVersionAndSubDir(
		meta.FrameworkVersion, false,
	)
	g := generator.Generator{
		CheckIfNameIsSet: true,
		UnitExists: func(unit *generator.ProjectUnit) error {
			if _, err := os.Stat(unit.Root); !os.IsNotExist(err) {
				return errors.New("'" + unit.Name + "' application already exists")
			}

			return nil
		},
		EmptyDirsToCreateInUnit: []string{"views"},
	}
	err = g.NewUnit(&unitCfg, "application")
	if err != nil {
		return err
	}

	return nil
}
