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
	newCommandCmdName = "new-command"
	NewCommandCmdDescription = newCommandCmdName + ":\tcreates a new command"
)

var (
	NewCommandCmd = flag.NewFlagSet(newCommandCmdName, flag.ExitOnError)
	ncAppNameFlag string
	ncNameFlag string
	ncCommandPathFlag string
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
	NewCommandCmd.StringVar(
		&ncCommandPathFlag,
		"l", "", "Location of a new command",
	)
}

func trimCommandSuffix(appName string) string {
	appName = strings.TrimSuffix(appName, "command")
	appName = strings.TrimSuffix(appName, "Command")
	return strings.TrimSuffix(appName, "_")
}

func (c *Cmd) CreateCommand() error {
	if len(ncNameFlag) == 0 {
		return errors.New("command name is not specified")
	}

	if len(ncAppNameFlag) == 0 && len(ncCommandPathFlag) == 0 {
		return errors.New("application or path to a new command is required")
	}

	if len(ncAppNameFlag) != 0 {
		ncAppNameFlag = trimAppSuffix(ncAppNameFlag)
	}

	c.process = func(cwd string, cfg *generator.ProjectUnit) error {
		cfg.Name = ncNameFlag
		cfg.ProjectRoot = cwd
		cfg.Templates = packr.New("Command Templates Box", "../../templates/command")
		var appPath string
		if len(ncAppNameFlag) != 0 {
			appPath = path.Join(cfg.ProjectRoot, ncAppNameFlag + "_app")
			if !utils.DirExists(appPath) {
				return errors.New("'" + ncAppNameFlag + "' application does not exist")
			}
		}

		cfg.Customize = func(pu *generator.ProjectUnit) {
			if len(ncAppNameFlag) == 0 {
				if path.IsAbs(ncCommandPathFlag) {
					pu.Root = ncCommandPathFlag
				} else {
					pu.Root = path.Join(pu.ProjectRoot, ncCommandPathFlag)
				}
			} else {
				pu.Root = path.Join(appPath, "commands")
			}

			pu.Name = trimCommandSuffix(pu.Name)
		}

		gen := generator.Generator{
			CheckIfNameIsSet: true,
			ErrorIfFileExists: func() error {
				return errors.New("'" + cfg.Name + "' command already exists")
			},
			FilePathSetup: func(fp string, fn string) (string, string) {
				return fp, strings.Replace(fn, "_name_", strcase.ToSnake(cfg.Name), 1)
			},
		}

		err := gen.NewUnit(cfg, "command")
		if err != nil {
			return err
		}

		return nil
	}

	c.postProcess = func(unit *generator.ProjectUnit) error {
		fmt.Printf("\nTo use '%s' command it must be registered.\n", unit.Name)
		fmt.Printf(
			"\nInclude command id application config and register it in '" +
			"commands()' method:\n",
		)

		unitNameCamel := strcase.ToCamel(unit.Name) + "Command"

		fmt.Printf("  this->command<%s>();\n", unitNameCamel)

		return nil
	}

	err := c.execute("command", false)
	if err != nil {
		return err
	}

	return nil
}
