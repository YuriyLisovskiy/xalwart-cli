package commands

import (
	"errors"
	"flag"
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

	unitCfg := generator.ProjectUnit{
		Year:                      time.Now().Year(),
		Username:                  usr.Name,
		FrameworkName:             config.FrameworkName,
		FrameworkNamespace:        config.FrameworkNamespace,
		Name:                      *nlNameFlag,
		ProjectRoot:               cwd,
		Templates:                 packr.New("Library Templates Box", "../templates/library"),
		Customize: func(pu *generator.ProjectUnit) {
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
		return err
	}

	unitCfg.ProjectName = meta.ProjectName

	_, unitCfg.FrameworkVersionSubDir, _ = getFWVersionAndSubDir(
		meta.FrameworkVersion, false,
	)
	g := generator.Generator{
		CheckIfNameIsSet: true,
		UnitExists: func(unit *generator.ProjectUnit) error {
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
