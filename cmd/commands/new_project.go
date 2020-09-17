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
	"xalwart-cli/utils"
)

const (
	newProjectCmdName = "new-project"
	NewProjectCmdDescription = newProjectCmdName + ":\tcreates a new project based on cmake lists"
)

var (
	NewProjectCmd = flag.NewFlagSet(newProjectCmdName, flag.ExitOnError)
	npQFlag bool
	npNameFlag string
	npRootFlag string
	npCMakeMinVersionFlag string
	npCppStandardFlag int
	npConanFlag bool
)

func InitNewProjectCmd() {
	NewProjectCmd.BoolVar(&npQFlag,
		"q",
		false,
		"Create project using questions instead of explicit arguments",
	)
	NewProjectCmd.StringVar(&npNameFlag, "n", "", "Project name")
	NewProjectCmd.StringVar(&npRootFlag, "r", "", "Project root")
	NewProjectCmd.StringVar(&npCMakeMinVersionFlag,
		"cmake", config.MinimumCmakeVersion, "Cmake minimum version",
	)
	NewProjectCmd.IntVar(&npCppStandardFlag, "cpp", config.MinimumCppStandard, "Standard of C++ language")
	NewProjectCmd.BoolVar(&npConanFlag,
		"conan", true, "Use Conan package manager",
	)
}

func normalizeAndCheckProjectConfig(cfg *generator.ProjectUnit) error {
	// Check C++ standard.
	if cfg.CMakeCPPStandard < config.MinimumCppStandard {
		fmt.Println("Warning: minimum required C++ standard is " + strconv.Itoa(config.MinimumCppStandard))
		cfg.CMakeCPPStandard = config.MinimumCppStandard
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
		cppStandard        int
		cmakeMinVer        string
		useConan           bool
	)

	if npQFlag {
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

		if cppStandard, err = reader.ReadInt(
			"Setup C++ standard (minimum required is " + strconv.Itoa(config.MinimumCppStandard) + "): ",
		); err != nil {
			return err
		}

		if cppStandard == 0 {
			cppStandard = config.MinimumCppStandard
		}

		if cmakeMinVer, err = reader.ReadString(
			"Input minimum version of cmake (default is " + config.MinimumCmakeVersion + "): ",
		); err != nil {
			return err
		}

		if useConan, err = reader.ReadBool(
			"Do you want to use Conan package manager? [Y/n] ",
		); err != nil {
			return err
		}

		if len(cmakeMinVer) == 0 {
			cmakeMinVer = config.MinimumCmakeVersion
		}
	} else {
		projectPath = npRootFlag
		projectName = npNameFlag
		cppStandard = npCppStandardFlag
		cmakeMinVer = npCMakeMinVersionFlag
		useConan = npConanFlag
	}

	if len(projectPath) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		projectPath = cwd
	}

	c.process = func(cwd string, cfg *generator.ProjectUnit) error {
		cfg.WorkingDirectory = projectPath
		cfg.ProjectName = projectName
		cfg.CMakeListsTxtToDoLine = config.CMakeListsTxtToDoLine
		cfg.CMakeCPPStandard = cppStandard
		cfg.UseConan = useConan
		cfg.CMakeMinimumVersion = cmakeMinVer
		cfg.SecretKey = generateSecretKey(config.SecretKeyLength)
		cfg.Templates = packr.New("Project Templates Box", "../../templates/project")
		cfg.Customize = func(pu *generator.ProjectUnit) {
			if len(pu.ProjectName) == 0 {
				pu.Root = pu.WorkingDirectory
				pu.ProjectName = path.Base(pu.Root)
			} else {
				pu.Root = path.Join(pu.WorkingDirectory, pu.ProjectName)
			}
		}

		if !cfg.UseConan {
			cfg.TemplatesToExclude = []string{"conanfile.txt.txt"}
		} else {
			cfg.ConanRequiredPackages = config.ConanRequiredPackages
		}

		err := normalizeAndCheckProjectConfig(cfg)
		if err != nil {
			return err
		}

		gen := generator.Generator{
			UnitExists: func(cfg *generator.ProjectUnit) error {
				if !utils.DirIsEmpty(cfg.Root) {
					return errors.New("root directory of a new project is not empty")
				}

				return nil
			},
			FilePathSetup: func(fp string, fn string) (string, string) {
				return strings.Replace(fp, "_proj_name_", cfg.ProjectName, -1), fn
			},
			EmptyDirsToCreateInUnit: []string{"media"},
		}

		err = gen.NewUnit(cfg, "project")
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
