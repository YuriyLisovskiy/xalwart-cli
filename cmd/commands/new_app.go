package commands

import (
	"errors"
	"flag"
	"github.com/gobuffalo/packr/v2"
	"os"
	"path"
	"strings"
	"xalwart-cli/generator"
)

const newAppCmdName = "new-app"

var (
	NewAppCmd = flag.NewFlagSet(newAppCmdName, flag.ExitOnError)

	naNameFlag = NewAppCmd.String("name", "", "Name of a new application")
)

func (c *Cmd) CreateApp() error {
	c.customizeUnit = func(cwd string, unit *generator.ProjectUnit) error {
		unit.Name = *naNameFlag
		unit.ProjectRoot = cwd
		unit.Templates = packr.New("App Templates Box", "../../templates/app")
		unit.Customize = func(pu *generator.ProjectUnit) {
			if !strings.HasSuffix(strings.ToLower(pu.Name), "_app") {
				pu.Name += "_app"
			}

			pu.Root = path.Join(pu.ProjectRoot, pu.Name)
		}

		return nil
	}
	
	c.makeGenerator = func(pu *generator.ProjectUnit) generator.Generator {
		return generator.Generator{
			CheckIfNameIsSet: true,
			UnitExists: func(unit *generator.ProjectUnit) error {
				if _, err := os.Stat(unit.Root); !os.IsNotExist(err) {
					return errors.New("'" + unit.Name + "' application already exists")
				}

				return nil
			},
			EmptyDirsToCreateInUnit: []string{"views"},
		}
	}

	err := c.execute("application", false)
	if err != nil {
		return err
	}

	return nil
}
