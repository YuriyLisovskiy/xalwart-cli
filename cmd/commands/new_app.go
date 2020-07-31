package commands

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/iancoleman/strcase"
	"os"
	"path"
	"strings"
	"xalwart-cli/config"
	"xalwart-cli/generator"
)

const (
	newAppCmdName = "new-app"
	NewAppCmdDescription = newAppCmdName +
		":\tadds a new application to existing '" + config.FrameworkName + "' application"
)

var (
	NewAppCmd = flag.NewFlagSet(newAppCmdName, flag.ExitOnError)
	naNameFlag string
	naAppPathFlag string
)

func InitNewAppCmd() {
	NewAppCmd.StringVar(&naNameFlag, "n", "", "Name of a new application")
	NewAppCmd.StringVar(&naAppPathFlag, "l", "", "Location of a new application")
}

func trimAppSuffix(appName string) string {
	appName = strings.TrimSuffix(appName, "config")
	appName = strings.TrimSuffix(appName, "Config")
	appName = strings.TrimSuffix(appName, "_")
	appName = strings.TrimSuffix(appName, "app")
	appName = strings.TrimSuffix(appName, "App")
	return strings.TrimSuffix(appName, "_")
}

func (c *Cmd) CreateApp() error {
	if len(naNameFlag) == 0 {
		return errors.New("application name is not specified")
	}

	c.process = func(cwd string, cfg *generator.ProjectUnit) error {
		cfg.Name = naNameFlag
		cfg.ProjectRoot = cwd
		cfg.Templates = packr.New("App Templates Box", "../../templates/app")
		cfg.Customize = func(pu *generator.ProjectUnit) {
			pu.Name = trimAppSuffix(pu.Name) + "App"
			if len(naAppPathFlag) != 0 {
				if path.IsAbs(naAppPathFlag) {
					pu.Root = naAppPathFlag
				} else {
					pu.Root = path.Join(pu.ProjectRoot, naAppPathFlag)
				}
			} else {
				pu.Root = pu.ProjectRoot
			}

			pu.Root = path.Join(pu.Root, strcase.ToSnake(pu.Name))
		}

		gen := generator.Generator{
			CheckIfNameIsSet: true,
			UnitExists: func(unit *generator.ProjectUnit) error {
				if _, err := os.Stat(unit.Root); !os.IsNotExist(err) {
					return errors.New(
						"'" + strings.TrimSuffix(unit.Name, "App") +
							"' application already exists",
					)
				}

				return nil
			},
			EmptyDirsToCreateInUnit: []string{"views"},
		}

		err := gen.NewUnit(cfg, "application")
		if err != nil {
			return err
		}

		return nil
	}

	c.postProcess = func(unit *generator.ProjectUnit) error {
		appName := trimAppSuffix(unit.Name)

		fmt.Printf("\nTo make '%s' application work it must be registered.\n", appName)

		fmt.Println(
			"\nInclude app config in Settings and register an " +
			"application in 'Settings::register_apps()' method:",
		)

		appNameSnake := strcase.ToSnake(appName)
		appNameSnake = strings.TrimSuffix(appNameSnake, "_app")
		appNameCamel := strcase.ToCamel(appNameSnake)

		fmt.Printf("  this->app<%sAppConfig>(\"%sApp\");\n", appNameCamel, appNameCamel)
		fmt.Println("\nActivate app in 'config.yml' in 'installed_apps':")
		fmt.Printf("  - %sApp\n", appNameCamel)
		fmt.Printf(
			"\nInclude in main application and insert the next code in 'urlpatterns()':\n",
		)
		fmt.Printf(
			"  this->include<%sAppConfig>(R\"(%s/)\", \"%s\");\n\n",
			appNameCamel,
			appNameSnake,
			appNameSnake,
		)

		return nil
	}

	err := c.execute("application", false)
	if err != nil {
		return err
	}

	return nil
}
