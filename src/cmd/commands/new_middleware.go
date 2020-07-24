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
)

const (
	newMiddlewareCmdName = "new-middleware"
	NewMiddlewareCmdDescription = newMiddlewareCmdName + ":\tcreates a new middleware"
)

var (
	NewMiddlewareCmd = flag.NewFlagSet(newMiddlewareCmdName, flag.ExitOnError)
	nmMiddlewarePathFlag string
	nmNameFlag string
)

func InitNewMiddlewareCmd() {
	NewMiddlewareCmd.StringVar(
		&nmMiddlewarePathFlag,
		"l", "", "Location of a new middleware",
	)
	NewMiddlewareCmd.StringVar(
		&nmNameFlag,
		"n", "", "Name of a new middleware",
	)
}

func trimMiddlewareSuffix(appName string) string {
	appName = strings.TrimSuffix(appName, "middleware")
	appName = strings.TrimSuffix(appName, "Middleware")
	return strings.TrimSuffix(appName, "_")
}

func (c *Cmd) CreateMiddleware() error {
	c.customizeUnit = func(cwd string, unit *generator.ProjectUnit) error {
		unit.Name = nmNameFlag
		unit.ProjectRoot = cwd
		unit.Templates = packr.New("Middleware Templates Box", "../../templates/middleware")
		unit.Customize = func(pu *generator.ProjectUnit) {
			if len(nmMiddlewarePathFlag) != 0 {
				if path.IsAbs(nmMiddlewarePathFlag) {
					pu.Root = nmMiddlewarePathFlag
				} else {
					pu.Root = path.Join(pu.ProjectRoot, nmMiddlewarePathFlag)
				}
			} else {
				pu.Root = path.Join(pu.ProjectRoot, pu.ProjectName, "middleware")
			}

			pu.Name = trimMiddlewareSuffix(pu.Name)
		}

		return nil
	}

	c.makeGenerator = func(pu *generator.ProjectUnit) generator.Generator {
		return generator.Generator{
			CheckIfNameIsSet: true,
			ErrorIfFileExists: func() error {
				return errors.New("'" + pu.Name + "' middleware already exists")
			},
			FilePathSetup: func(fp string, fn string) (string, string) {
				return fp, strings.Replace(fn, "_name_", strcase.ToSnake(pu.Name), 1)
			},
		}
	}

	c.postCreateHelp = func(unit *generator.ProjectUnit) {
		fmt.Printf("\nTo use '%s' middleware it must be registered.\n", unit.Name)

		unitNameSnake := strcase.ToSnake(unit.Name)
		var middlewarePath string
		if len(nmMiddlewarePathFlag) != 0 {
			middlewarePath = path.Join(unit.Root, unitNameSnake + ".h")
			if !path.IsAbs(nmMiddlewarePathFlag) {
				middlewarePath = path.Join("..", middlewarePath)
			}
		} else {
			middlewarePath = "./middleware/" + unitNameSnake + ".h"
		}

		fmt.Printf("\nInclude in 'Settings':\n  #include \"%s\"\n", middlewarePath)
		fmt.Println("\nRegister middleware in 'Settings::register_middleware()' method:")

		unitNameCamel := strcase.ToCamel(unit.Name) + "Middleware"

		fmt.Printf("  this->middleware<%s>(\"%s\");\n", unitNameCamel, unitNameCamel)
		fmt.Println("\nActivate middleware in 'config.yml' in 'middleware':")
		fmt.Printf("  - %s\n", unitNameCamel)
	}

	err := c.execute("middleware", false)
	if err != nil {
		return err
	}

	return nil
}
