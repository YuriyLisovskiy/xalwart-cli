package commands

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"os"
	"os/user"
	"path"
	"time"
	"xalwart-cli/config"
	"xalwart-cli/generator"
	"xalwart-cli/utils"
)

const newLibraryCmdName = "new-lib"

var (
	NewLibraryCmd = flag.NewFlagSet(newLibraryCmdName, flag.ExitOnError)

	nlLibPathFlag = NewLibraryCmd.String("path", "", "Location of a new library")
	nlNameFlag = NewLibraryCmd.String("name", "", "Name of a new library")
)

func CreateLibrary() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	usr, err := user.Current()
	if err != nil {
		return err
	}

	unitCfg := config.ProjectUnit{
		Year:                      time.Now().Year(),
		Username:                  usr.Name,
		FrameworkName:             config.FrameworkName,
		FrameworkNamespace:        config.FrameworkNamespace,
		Name:                      *nlNameFlag,
		ProjectRoot:               cwd,
		Templates:                 packr.New("Library Templates Box", "../templates/library"),
		Customize: func(pu *config.ProjectUnit) {
			libPath := *nlLibPathFlag
			if len(libPath) != 0 {
				pu.Root = libPath
			} else {
				pu.Root = path.Join(pu.ProjectRoot, pu.ProjectName, "libs", pu.Name)
			}
		},
	}

	meta, err := loadMeta(unitCfg.ProjectRoot)
	if err != nil {
		fmt.Println("Warning: " + err.Error())
	}

	unitCfg.ProjectName = meta.ProjectName

	_, unitCfg.FrameworkVersionSubDir, _ = getFWVersionAndSubDir(
		meta.Version, false,
	)
	g := generator.Generator{
		CheckIfUnitExists: true,
		UnitExists: func(unit *config.ProjectUnit) error {
			if utils.DirExists(unit.Root) {
				return errors.New("'" + unit.Name + "' library already exists")
			}

			return nil
		},
		EmptyDirsToCreateInUnit: []string{"tags", "filters"},
	}
	err = g.NewUnit(&unitCfg, "library")
	if err != nil {
		return err
	}

	return nil
}
