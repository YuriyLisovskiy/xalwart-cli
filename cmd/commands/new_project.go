package commands

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"time"
	"xalwart-cli/config"
	"xalwart-cli/generator"
	"xalwart-cli/manager"
	"xalwart-cli/utils"
)

var (
	NewProjectCmd = flag.NewFlagSet("new-project", flag.ExitOnError)

	qFlag = NewProjectCmd.Bool(
		"q",
		true,
		"Execute 'new-project' command with questions or use arguments",
	)
	nameFlag = NewProjectCmd.String("name", "", "Project name")
	rootFlag = NewProjectCmd.String("root", "", "Project root")
	cMakeMinVersionFlag = NewProjectCmd.String(
		"cmake-version", config.MinimumCmakeVersion, "Cmake minimum version",
	)
	cppStandardFlag = NewProjectCmd.Int("cpp", config.MinimumCppVersion, "C++ standard")
	frameworkVersionFlag = NewProjectCmd.String(
		"fw-version",
		"latest",
		"A version of '" + config.FrameworkName + "' framework to install",
	)
)

func reSubMatchMap(r *regexp.Regexp, str string) (bool, map[string]string) {
	if r.MatchString(str) {
		match := r.FindStringSubmatch(str)
		subMatchMap := make(map[string]string)
		for i, name := range r.SubexpNames() {
			if i != 0 {
				subMatchMap[name] = match[i]
			}
		}

		return true, subMatchMap
	}

	return false, nil
}

func normalizeAndCheckProjectConfig(cfg *config.Project) error {
	// Check C++ standard.
	if cfg.CMakeCPPStandard < 17 {
		fmt.Println("Warning: minimum required C++ standard is 17")
		cfg.CMakeCPPStandard = 17
	}

	// Check version of framework.
	if cfg.FrameworkVersion == "latest" {
		var err error
		cfg.FrameworkVersion, err = manager.GetLatestVersionOfFramework()
		if err != nil {
			return errors.New(
				"unable to retrieve latest release of '" + config.FrameworkName + "' framework",
			)
		}
	} else {
		isAvailable, err := manager.CheckIfVersionIsAvailable(cfg.FrameworkVersion)
		if err != nil {
			return err
		}

		if !isAvailable {
			fmt.Println(
				"Warning: release v" + cfg.FrameworkVersion + " is not available, latest is used instead",
			)

			cfg.FrameworkVersion, err = manager.GetLatestVersionOfFramework()
			if err != nil {
				fmt.Println(err)
				return errors.New(
					"unable to retrieve latest release of '" + config.FrameworkName + "' framework",
				)
			}
		}
	}

	expr, _ := regexp.Compile(
		"(?P<major>\\d+)\\.(?P<minor>\\d+)\\.(?P<patch>\\d+)(-(?P<pre_release>(alpha|beta)))?",
	)
	match, ver := reSubMatchMap(expr, cfg.FrameworkVersion)
	if match {
		if ver["major"] > "1" {
			cfg.FrameworkVersionSubDir = "v" + ver["major"] + "/"
		}
	} else {
		fmt.Println(
			"Warning: invalid version of '" + config.FrameworkName + "' framework. Used latest by default.",
		)
	}

	return nil
}

func generateSecretKey(n int) string {
	// TODO add random key generation.
	return "+s6cv712&nw4gsk)1dmgpje+f#%^4lhp@!up+=p3ts+hxz(fr2"
}

func CreateProject() error {
	var (
		projectRoot string
		projectName string
		frameworkVerSubDir = ""
		frameworkVer string
		cppStandard int
		cmakeMinVer string
	)

	if *qFlag {
		var err error
		reader := utils.NewIO()
		if projectRoot, err = reader.ReadString(
			"Type root folder of a new project (default is current working directory): ",
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
		projectRoot = *rootFlag
		projectName = *nameFlag
		frameworkVer = *frameworkVersionFlag
		cppStandard = *cppStandardFlag
		cmakeMinVer = *cMakeMinVersionFlag
	}

	if len(projectRoot) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		projectRoot = cwd
	}

	usr, err := user.Current()
	if err != nil {
		return err
	}

	cfg := config.Project{
		Year:                      time.Now().Year(),
		Username:                  usr.Username,
		WorkingDirectory:          projectRoot,
		FrameworkName:             config.FrameworkName,
		FrameworkNamespace:        config.FrameworkNamespace,
		FrameworkVersion:          frameworkVer,
		FrameworkVersionSubDir:    frameworkVerSubDir,
		ProjectName:               projectName,
		SecretKey:                 generateSecretKey(config.SecretKeyLength),
		CMakeCPPStandard:          cppStandard,
		CMakeMinimumVersion:       cmakeMinVer,
		Templates:                 packr.New("Project Templates Box", "../templates/project"),
	}

	err = normalizeAndCheckProjectConfig(&cfg)
	if err != nil {
		return err
	}

	g := generator.Generator{}
	err = g.NewProject(&cfg)
	if err != nil {
		return err
	}

	err = manager.InstallFramework(cfg.ProjectRoot, cfg.FrameworkVersion)
	if err != nil {
		return err
	}

	return nil
}
