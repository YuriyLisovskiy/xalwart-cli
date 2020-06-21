package commands

import (
	"errors"
	"flag"
	"fmt"
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

	unitCfg := config.ProjectUnit{
		Year:                      time.Now().Year(),
		Username:                  usr.Name,
		FrameworkName:             config.FrameworkName,
		FrameworkNamespace:        config.FrameworkNamespace,
		Name:                      *naNameFlag,
		ProjectRoot:               cwd,
		Templates:                 packr.New("App Templates Box", "../templates/app"),
		Customize: func(pu *config.ProjectUnit) {
			if !strings.HasSuffix(strings.ToLower(pu.Name), "_app") {
				pu.Name += "_app"
			}

			pu.Root = path.Join(pu.ProjectRoot, pu.Name)
		},
	}

	meta, err := loadMeta(unitCfg.ProjectRoot)
	if err != nil {
		fmt.Println("Warning: " + err.Error())
	}

	unitCfg.ProjectName = meta.ProjectName

	_, unitCfg.FrameworkVersionSubDir, _ = getFWVersionAndSubDir(
		meta.Version, false,
	)
	g := generator.Generator{
		CheckIfUnitExists: true,
		UnitExists: func(unit *config.ProjectUnit) error {
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
