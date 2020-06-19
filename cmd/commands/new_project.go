package commands

import (
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"os"
	"os/user"
	"path"
	"regexp"
	"strconv"
	"time"
	"xalwart-cli/config"
	"xalwart-cli/generator"
	"xalwart-cli/manager"
	"xalwart-cli/utils"
)

const (
	frameworkName = "xalwart"
	frameworkNamespace = "xw"
	minimumCmakeVersion = "3.13"
	minimumCppVersion = 17
)

var (
	NewProjectCmd = flag.NewFlagSet("new-project", flag.ExitOnError)

	newProjectAskQuestions = NewProjectCmd.Bool(
		"q",
		true,
		"Execute 'new-project' command with questions or use arguments",
	)
	newProjectName = NewProjectCmd.String("name", "", "Project name")
	newProjectRoot = NewProjectCmd.String("root", "", "Project root")
	newProjectCMakeMinVersion = NewProjectCmd.String(
		"cmake-min-v", minimumCmakeVersion, "Cmake minimum version",
	)
	newProjectCppStandard = NewProjectCmd.Int("cpp", minimumCppVersion, "C++ standard")
	newProjectFrameworkVersion = NewProjectCmd.String(
		frameworkName + "-version",
		"latest",
		"A version of 'xalwart' framework to install",
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
		// TODO: retrieve latest version of xalwart framework from github
	} else {
		expr, _ := regexp.Compile(
			"(?P<major>\\d+)\\.(?P<minor>\\d+)\\.(?P<patch>\\d+)(-(?P<pre_release>(alpha|beta)))?",
		)
		match, ver := reSubMatchMap(expr, cfg.FrameworkVersion)
		if match {
			if ver["major"] != "1" {
				cfg.FrameworkVersionSubDir = "v" + ver["major"] + "/"
			}
		} else {
			fmt.Println(
				"Warning: invalid version of '" + frameworkName + "' framework. Used latest by default.",
			)
		}
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

	if *newProjectAskQuestions {
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
			"Input a version of '" + frameworkName + "' framework which you want to use (default is latest): ",
		); err != nil {
			return err
		}

		if len(frameworkVer) == 0 {
			frameworkVer = "latest"
		}

		if cppStandard, err = reader.ReadInt(
			"Setup C++ standard (minimum required is " + strconv.Itoa(minimumCppVersion) + "): ",
		); err != nil {
			return err
		}

		if cppStandard == 0 {
			cppStandard = minimumCppVersion
		}

		if cmakeMinVer, err = reader.ReadString(
			"Type minimum version of cmake (default is " + minimumCmakeVersion + "): ",
		); err != nil {
			return err
		}

		if len(cmakeMinVer) == 0 {
			cmakeMinVer = minimumCmakeVersion
		}
	} else {
		projectRoot = *newProjectRoot
		projectName = *newProjectName
		frameworkVer = *newProjectFrameworkVersion
		cppStandard = *newProjectCppStandard
		cmakeMinVer = *newProjectCMakeMinVersion
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
		FrameworkName:             frameworkName,
		FrameworkNamespace:        frameworkNamespace,
		FrameworkVersion:          frameworkVer,
		FrameworkVersionSubDir:    frameworkVerSubDir,
		ProjectName:               projectName,
		SecretKey:                 generateSecretKey(50),
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

	err = manager.InstallFramework(
		path.Join(cfg.ProjectRoot, "include"),
		path.Join(cfg.ProjectRoot, "lib/" + frameworkName + "/" + frameworkVer),
		cfg.FrameworkVersion,
	)
	if err != nil {
		return err
	}

	return nil
}
