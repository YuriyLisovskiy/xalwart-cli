package commands

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/iancoleman/strcase"
	"path"
	"strings"
	"xalwart-cli/src/generator"
	"xalwart-cli/src/utils"
)

const (
	newViewCmdName = "new-view"
	NewViewCmdDescription = newViewCmdName + ":\tcreates a new view"
)

var (
	NewViewCmd = flag.NewFlagSet(newViewCmdName, flag.ExitOnError)
	nvAppNameFlag string
	nvNameFlag string
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
}

func trimViewSuffix(appName string) string {
	appName = strings.TrimSuffix(appName, "view")
	appName = strings.TrimSuffix(appName, "View")
	return strings.TrimSuffix(appName, "_")
}

func (c *Cmd) CreateView() error {
	if len(nvAppNameFlag) == 0 {
		return errors.New("application is not specified")
	}

	nvAppNameFlag = trimAppSuffix(nvAppNameFlag)
	c.customizeUnit = func(cwd string, unit *generator.ProjectUnit) error {
		unit.Name = nvNameFlag
		unit.ProjectRoot = cwd
		unit.Templates = packr.New("View Templates Box", "../../templates/view")
		appPath := path.Join(unit.ProjectRoot, nvAppNameFlag + "_app")
		if !utils.DirExists(appPath) {
			return errors.New("'" + nvAppNameFlag + "' application does not exist")
		}

		unit.Customize = func(pu *generator.ProjectUnit) {
			pu.Root = path.Join(appPath, "views")
			pu.Name = trimViewSuffix(pu.Name)
		}

		return nil
	}

	c.makeGenerator = func(pu *generator.ProjectUnit) generator.Generator {
		return generator.Generator{
			CheckIfNameIsSet: true,
			ErrorIfFileExists: func() error {
				return errors.New("'" + pu.Name + "' view already exists")
			},
			FilePathSetup: func(fp string, fn string) (string, string) {
				return fp, strings.Replace(fn, "_name_", strcase.ToSnake(pu.Name), 1)
			},
		}
	}

	c.postCreateHelp = func(unit *generator.ProjectUnit) {
		fmt.Printf("\nTo use '%s' view it must be registered.\n", unit.Name)

		unitNameSnake := strcase.ToSnake(unit.Name)

		fmt.Printf(
			"\nInclude in '" + nvAppNameFlag + "_app' configuration:\n  #include \"%s\"\n",
			"./views/" + unitNameSnake + "_view.h",
		)
		fmt.Println(
			"\nSetup URL in '" + strcase.ToCamel(nvAppNameFlag) +
			"AppConfig::urlpatterns()' method:",
		)

		unitNameCamel := strcase.ToCamel(unit.Name) + "View"

		fmt.Printf(
			"  this->url<%s>(R\"(%s/?)\", \"%s\");\n",
			unitNameCamel, unitNameSnake, unitNameSnake,
		)
	}

	return c.execute("view", false)
}
