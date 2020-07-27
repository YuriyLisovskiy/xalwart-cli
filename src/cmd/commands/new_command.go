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
	newCommandCmdName = "new-command"
	NewCommandCmdDescription = newCommandCmdName + ":\tcreates a new command"
)

var (
	NewCommandCmd = flag.NewFlagSet(newCommandCmdName, flag.ExitOnError)
	ncAppNameFlag string
	ncNameFlag string
)

func InitNewCommandCmd() {
	NewCommandCmd.StringVar(
		&ncAppNameFlag,
		"a", "", "Name of application where you want to create command",
	)
	NewCommandCmd.StringVar(
		&ncNameFlag,
		"n", "", "Name of a new command",
	)
}

func trimCommandSuffix(appName string) string {
	appName = strings.TrimSuffix(appName, "command")
	appName = strings.TrimSuffix(appName, "Command")
	return strings.TrimSuffix(appName, "_")
}

func (c *Cmd) CreateCommand() error {
	if len(ncAppNameFlag) == 0 {
		return errors.New("application is not specified")
	}

	ncAppNameFlag = trimAppSuffix(ncAppNameFlag)
	c.customizeUnit = func(cwd string, unit *generator.ProjectUnit) error {
		unit.Name = ncNameFlag
		unit.ProjectRoot = cwd
		unit.Templates = packr.New("Command Templates Box", "../../templates/command")
		appPath := path.Join(unit.ProjectRoot, ncAppNameFlag + "_app")
		if !utils.DirExists(appPath) {
			return errors.New("'" + ncAppNameFlag + "' application does not exist")
		}

		unit.Customize = func(pu *generator.ProjectUnit) {
			pu.Root = path.Join(appPath, "commands")
			pu.Name = trimCommandSuffix(pu.Name)
		}

		return nil
	}

	c.makeGenerator = func(pu *generator.ProjectUnit) generator.Generator {
		return generator.Generator{
			CheckIfNameIsSet: true,
			ErrorIfFileExists: func() error {
				return errors.New("'" + pu.Name + "' command already exists")
			},
			FilePathSetup: func(fp string, fn string) (string, string) {
				return fp, strings.Replace(fn, "_name_", strcase.ToSnake(pu.Name), 1)
			},
		}
	}

	c.postCreateHelp = func(unit *generator.ProjectUnit) {
		fmt.Printf("\nTo use '%s' command it must be registered.\n", unit.Name)

		unitNameSnake := strcase.ToSnake(unit.Name)

		fmt.Printf(
			"\nInclude in '" + ncAppNameFlag + "_app' configuration:\n  #include \"%s\"\n",
			"./commands/" + unitNameSnake + ".h",
		)
		fmt.Println(
			"\nRegister command in '" + strcase.ToCamel(ncAppNameFlag) +
			"AppConfig::commands()' method:",
		)

		unitNameCamel := strcase.ToCamel(unit.Name) + "Command"

		fmt.Printf("  this->command<%s>();\n", unitNameCamel)
	}

	err := c.execute("command", false)
	if err != nil {
		return err
	}

	return nil
}
