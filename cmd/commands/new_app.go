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
)

func InitNewAppCmd() {
	NewAppCmd.StringVar(&naNameFlag, "n", "", "Name of a new application")
}

func trimAppSuffix(appName string) string {
	appName = strings.TrimSuffix(appName, "app")
	appName = strings.TrimSuffix(appName, "App")
	return strings.TrimSuffix(appName, "_")
}

func (c *Cmd) CreateApp() error {
	c.customizeUnit = func(cwd string, unit *generator.ProjectUnit) error {
		unit.Name = naNameFlag
		unit.ProjectRoot = cwd
		unit.Templates = packr.New("App Templates Box", "../../templates/app")
		unit.Customize = func(pu *generator.ProjectUnit) {
			pu.Name = trimAppSuffix(pu.Name) + "App"
			pu.Root = path.Join(pu.ProjectRoot, strcase.ToSnake(pu.Name))
		}

		return nil
	}

	c.makeGenerator = func(pu *generator.ProjectUnit) generator.Generator {
		return generator.Generator{
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
	}

	c.postCreateHelp = func(unit *generator.ProjectUnit) {
		appName := trimAppSuffix(unit.Name)

		fmt.Printf("\nTo make '%s' application work it must be registered.\n", appName)

		folderName := strcase.ToSnake(appName)
		if !strings.HasSuffix(strings.ToLower(folderName), "_app") {
			folderName += "_app"
		}

		fmt.Printf("\nInclude in 'Settings':\n  #include \"../%s/app.h\"\n", folderName)
		fmt.Println("\nRegister application in 'Settings::register_apps()' method:")

		appNameSnake := strcase.ToSnake(appName)
		appNameSnake = strings.TrimSuffix(appNameSnake, "_app")
		appNameCamel := strcase.ToCamel(appNameSnake)

		fmt.Printf("  this->app<%sAppConfig>(\"%sApp\");\n", appNameCamel, appNameCamel)
		fmt.Println("\nActivate app in 'config.yml' in 'installed_apps':")
		fmt.Printf("  - %sApp\n", appNameCamel)
		fmt.Printf("\nInclude in main application:\n  #include \"../%s/app.h\"\n", folderName)
		fmt.Println(
			"\nInsert the next code in 'urlpatterns()':",
		)
		fmt.Printf(
			"  this->include<%sAppConfig>(R\"(%s/)\", \"%s\");\n\n",
			appNameCamel,
			appNameSnake,
			appNameSnake,
		)
	}

	err := c.execute("application", false)
	if err != nil {
		return err
	}

	return nil
}
