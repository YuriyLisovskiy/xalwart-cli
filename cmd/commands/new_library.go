package commands

import (
	"errors"
	"flag"
	"github.com/gobuffalo/packr/v2"
	"path"
	"xalwart-cli/generator"
	"xalwart-cli/utils"
)

const newLibraryCmdName = "new-lib"

var (
	NewLibraryCmd = flag.NewFlagSet(newLibraryCmdName, flag.ExitOnError)

	nlLibPathFlag = NewLibraryCmd.String("path", "", "Location of a new library")
	nlNameFlag = NewLibraryCmd.String("name", "", "Name of a new library")
)

func (c *Cmd) CreateLibrary() error {
	c.customizeUnit = func(cwd string, unit *generator.ProjectUnit) error {
		unit.Name = *nlNameFlag
		unit.ProjectRoot = cwd
		unit.Templates = packr.New("Library Templates Box", "../../templates/library")
		unit.Customize = func(pu *generator.ProjectUnit) {
			libPath := *nlLibPathFlag
			if len(libPath) != 0 {
				pu.Root = libPath
			} else {
				pu.Root = path.Join(pu.ProjectRoot, pu.ProjectName, "libs", pu.Name)
			}
		}

		return nil
	}

	c.makeGenerator = func(pu *generator.ProjectUnit) generator.Generator {
		return generator.Generator{
			CheckIfNameIsSet: true,
			UnitExists: func(unit *generator.ProjectUnit) error {
				if utils.DirExists(unit.Root) {
					return errors.New("'" + unit.Name + "' library already exists")
				}

				return nil
			},
			EmptyDirsToCreateInUnit: []string{"tags", "filters"},
		}
	}

	err := c.execute("library", false)
	if err != nil {
		return err
	}

	return nil
}
