package commands

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"xalwart-cli/config"
	"xalwart-cli/generator"
	"xalwart-cli/managers"
	"xalwart-cli/utils"
)

const newProjectCmdName = "new-project"

var (
	NewProjectCmd = flag.NewFlagSet(newProjectCmdName, flag.ExitOnError)

	npAFlag = NewProjectCmd.Bool(
		"a",
		false,
		"Execute '"+newProjectCmdName+"' command using arguments",
	)
	npNameFlag            = NewProjectCmd.String("name", "", "Project name")
	npRootFlag            = NewProjectCmd.String("root", "", "Project root")
	npCMakeMinVersionFlag = NewProjectCmd.String(
		"cmake-version", config.MinimumCmakeVersion, "Cmake minimum version",
	)
	npCppStandardFlag      = NewProjectCmd.Int("cpp", config.MinimumCppVersion, "C++ standard")
	npFrameworkVersionFlag = NewProjectCmd.String(
		"fw-version",
		"latest",
		"A version of '"+config.FrameworkName+"' framework to install",
	)
)

func normalizeAndCheckProjectConfig(cfg *generator.ProjectUnit) error {
	// Check C++ standard.
	if cfg.CMakeCPPStandard < config.MinimumCppVersion {
		fmt.Println("Warning: minimum required C++ standard is " + strconv.Itoa(config.MinimumCppVersion))
		cfg.CMakeCPPStandard = config.MinimumCppVersion
	}

	return nil
}

func generateSecretKey(n int) string {
	charset := "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*(-_=+)"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func (c *Cmd) CreateProject() error {
	var (
		projectPath        string
		projectName        string
		frameworkVer       string
		cppStandard        int
		cmakeMinVer        string
	)

	if !*npAFlag {
		var err error
		reader := utils.NewIO()
		if projectPath, err = reader.ReadString(
			"Type path to folder where a new project will be created (default is current working directory): ",
		); err != nil {
			return err
		}

		if projectName, err = reader.ReadString(
			"Enter name or leave empty to use current folder name: ",
		); err != nil {
			return err
		}

		if frameworkVer, err = reader.ReadString(
			"Input a version of '" + config.FrameworkName + "' framework which you want to use (default is latest): ",
		); err != nil {
			return err
		}

		if len(frameworkVer) == 0 {
			frameworkVer = "latest"
		}

		if cppStandard, err = reader.ReadInt(
			"Setup C++ standard (minimum required is " + strconv.Itoa(config.MinimumCppVersion) + "): ",
		); err != nil {
			return err
		}

		if cppStandard == 0 {
			cppStandard = config.MinimumCppVersion
		}

		if cmakeMinVer, err = reader.ReadString(
			"Type minimum version of cmake (default is " + config.MinimumCmakeVersion + "): ",
		); err != nil {
			return err
		}

		if len(cmakeMinVer) == 0 {
			cmakeMinVer = config.MinimumCmakeVersion
		}
	} else {
		projectPath = *npRootFlag
		projectName = *npNameFlag
		frameworkVer = *npFrameworkVersionFlag
		cppStandard = *npCppStandardFlag
		cmakeMinVer = *npCMakeMinVersionFlag
	}

	if len(projectPath) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		projectPath = cwd
	}

	c.customizeUnit = func(cwd string, unit *generator.ProjectUnit) error {
		unit.WorkingDirectory = projectPath
		unit.ProjectName = projectName
		unit.FrameworkVersion = frameworkVer
		unit.CMakeCPPStandard = cppStandard
		unit.CMakeMinimumVersion = cmakeMinVer
		unit.SecretKey = generateSecretKey(config.SecretKeyLength)
		unit.Templates = packr.New("Project Templates Box", "../../templates/project")
		unit.Customize = func(pu *generator.ProjectUnit) {
			if len(pu.ProjectName) == 0 {
				pu.Root = pu.WorkingDirectory
				pu.ProjectName = path.Base(pu.Root)
			} else {
				pu.Root = path.Join(pu.WorkingDirectory, pu.ProjectName)
			}
		}

		err := normalizeAndCheckProjectConfig(unit)
		if err != nil {
			return err
		}

		return nil
	}

	c.makeGenerator = func(unit *generator.ProjectUnit) generator.Generator {
		return generator.Generator{
			UnitExists: func(cfg *generator.ProjectUnit) error {
				if !utils.DirIsEmpty(cfg.Root) {
					return errors.New("root directory of a new project is not empty")
				}

				return nil
			},
			FilePathSetup: func(fp string, fn string) (string, string) {
				return strings.Replace(fp, "_proj_name_", unit.ProjectName, -1), fn
			},
			EmptyDirsToCreateInUnit: []string{"media"},
		}
	}

	c.postProcess = func(unit *generator.ProjectUnit) error {
		err := managers.InstallFramework(unit.Root, unit.FrameworkVersion)
		if err != nil {
			return err
		}

		return nil
	}

	err := c.execute("project", true)
	if err != nil {
		return err
	}

	return nil
}
