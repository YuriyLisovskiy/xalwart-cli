package commands

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/iancoleman/strcase"
	"path"
	"xalwart-cli/generator"
	"xalwart-cli/utils"
)

const (
	newLibraryCmdName = "new-lib"
	NewLibraryCmdDescription = newLibraryCmdName + ":\tcreates a new library for template engine"
)

var (
	NewLibraryCmd = flag.NewFlagSet(newLibraryCmdName, flag.ExitOnError)
	nlLibPathFlag string
	nlNameFlag string
)

func InitNewLibraryCmd() {
	NewLibraryCmd.StringVar(&nlLibPathFlag, "l", "", "Location of a new library")
	NewLibraryCmd.StringVar(&nlNameFlag, "n", "", "Name of a new library")
}

func (c *Cmd) CreateLibrary() error {
	if len(nlNameFlag) == 0 {
		return errors.New("library name is not specified")
	}

	c.process = func(cwd string, cfg *generator.ProjectUnit) error {
		cfg.Name = nlNameFlag
		cfg.ProjectRoot = cwd
		cfg.Templates = packr.New("Library Templates Box", "../../templates/library")
		cfg.Customize = func(pu *generator.ProjectUnit) {
			if len(nlLibPathFlag) != 0 {
				if path.IsAbs(nlLibPathFlag) {
					pu.Root = nlLibPathFlag
				} else {
					pu.Root = path.Join(pu.ProjectRoot, nlLibPathFlag)
				}

				pu.Root = path.Join(pu.Root, pu.Name)
			} else {
				pu.Root = path.Join(pu.ProjectRoot, pu.ProjectName, "libs", pu.Name)
			}
		}

		gen := generator.Generator{
			CheckIfNameIsSet: true,
			UnitExists: func(unit *generator.ProjectUnit) error {
				if utils.DirExists(unit.Root)  {
					return errors.New("'" + unit.Name + "' library already exists")
				}

				return nil
			},
		}

		err := gen.NewUnit(cfg, "library")
		if err != nil {
			return err
		}

		return nil
	}

	c.postProcess = func(unit *generator.ProjectUnit) error {
		fmt.Printf("\nTo use '%s' library to render engine it must be registered.\n", unit.Name)

		var libPath string
		if len(nlLibPathFlag) != 0 {
			libPath = path.Join(unit.Root, "library.h")
			if !path.IsAbs(nlLibPathFlag) {
				libPath = path.Join("..", libPath)
			}
		} else {
			libPath = "./libs/" + unit.Name + "/library.h"
		}

		fmt.Printf("\nInclude in 'Settings':\n  #include \"%s\"\n", libPath)
		fmt.Println("\nRegister library in 'Settings::register_libraries()' method:")

		libNameCamel := strcase.ToCamel(unit.Name)

		fmt.Printf("  this->library<%s>(\"%s\");\n", libNameCamel, libNameCamel)
		fmt.Println("\nActivate library in 'config.yml' in 'templates_env' -> 'libraries':")
		fmt.Printf("  - %s\n", libNameCamel)

		return nil
	}

	err := c.execute("library", false)
	if err != nil {
		return err
	}

	return nil
}
