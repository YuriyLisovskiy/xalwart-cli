package commands

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/iancoleman/strcase"
	"path"
	"strings"
	"xalwart-cli/generator"
	"xalwart-cli/utils"
)

const (
	newViewCmdName = "new-view"
	NewViewCmdDescription = newViewCmdName + ":\tcreates a new view"
)

var (
	NewViewCmd = flag.NewFlagSet(newViewCmdName, flag.ExitOnError)
	nvAppNameFlag string
	nvNameFlag string
	nvViewPathFlag string
)

func InitNewViewCmd() {
	NewViewCmd.StringVar(
		&nvAppNameFlag,
		"a", "", "Name of application where you want to add view",
	)
	NewViewCmd.StringVar(
		&nvNameFlag,
		"n", "", "Name of a new view",
	)
	NewViewCmd.StringVar(
		&nvViewPathFlag,
		"l", "", "Location of a new view",
	)
}

func trimViewSuffix(appName string) string {
	appName = strings.TrimSuffix(appName, "view")
	appName = strings.TrimSuffix(appName, "View")
	return strings.TrimSuffix(appName, "_")
}

func (c *Cmd) CreateView() error {
	if len(nvNameFlag) == 0 {
		return errors.New("view name is not specified")
	}

	if len(nvAppNameFlag) == 0 && len(nvViewPathFlag) == 0 {
		return errors.New("application or path to a new view is required")
	}

	if len(nvAppNameFlag) != 0 {
		nvAppNameFlag = trimAppSuffix(nvAppNameFlag)
	}

	c.process = func(cwd string, cfg *generator.ProjectUnit) error {
		cfg.Name = nvNameFlag
		cfg.ProjectRoot = cwd
		cfg.Templates = packr.New("View Templates Box", "../../templates/view")
		var appPath string
		if len(nvAppNameFlag) != 0 {
			appPath = path.Join(cfg.ProjectRoot, nvAppNameFlag + "_app")
			if !utils.DirExists(appPath) {
				return errors.New("'" + nvAppNameFlag + "' application does not exist")
			}
		}

		cfg.Customize = func(pu *generator.ProjectUnit) {
			if len(nvAppNameFlag) == 0 {
				if path.IsAbs(nvViewPathFlag) {
					pu.Root = nvViewPathFlag
				} else {
					pu.Root = path.Join(pu.ProjectRoot, nvViewPathFlag)
				}
			} else {
				pu.Root = path.Join(appPath, "views")
			}

			pu.Name = trimViewSuffix(pu.Name)
		}

		gen := generator.Generator{
			CheckIfNameIsSet: true,
			ErrorIfFileExists: func() error {
				return errors.New("'" + cfg.Name + "' view already exists")
			},
			FilePathSetup: func(fp string, fn string) (string, string) {
				return fp, strings.Replace(fn, "_name_", strcase.ToSnake(cfg.Name), 1)
			},
		}

		err := gen.NewUnit(cfg, "view")
		if err != nil {
			return err
		}

		return nil
	}

	c.postProcess = func(unit *generator.ProjectUnit) error {
		fmt.Printf("\nTo use '%s' view it must be registered.\n", unit.Name)

		unitNameSnake := strcase.ToSnake(unit.Name)

		fmt.Printf(
			"\nInclude view in application config and " +
			"setup URL in 'urlpatterns()' method:",
		)

		unitNameCamel := strcase.ToCamel(unit.Name) + "View"

		fmt.Printf(
			"  this->url<%s>(R\"(%s/?)\", \"%s\");\n",
			unitNameCamel, unitNameSnake, unitNameSnake,
		)

		return nil
	}

	return c.execute("view", false)
}
