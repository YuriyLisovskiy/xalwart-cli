package main

import (
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"os"
	"os/user"
	"path"
	"regexp"
	"time"
	"xalwart-cli/config"
	"xalwart-cli/generator"
	"xalwart-cli/manager"
)

const (
	frameworkName = "xalwart"
	frameworkNamespace = "xw"
)

var (
	newProjectCmd = flag.NewFlagSet("new-project", flag.ExitOnError)

	newProjectAskQuestions = newProjectCmd.Bool(
		"q",
		true,
		"Execute 'new-project' command with questions or use arguments",
	)
	newProjectName = newProjectCmd.String("name", "", "Project name")
	newProjectRoot = newProjectCmd.String("root", "", "Project root")
	newProjectCMakeMinVersion = newProjectCmd.String(
		"cmake-min-v", "3.13", "Cmake minimum version",
	)
	newProjectCppStandard = newProjectCmd.Int("cpp", 17, "C++ standard")
	newProjectFrameworkVersion = newProjectCmd.String(
		"xalwart-version",
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
		cfg.FrameworkVersion = "0.0.1"
	} else {
		expr, _ := regexp.Compile("(?P<major>\\d+)\\.(?P<minor>\\d+)\\.(?P<patch>\\d+)")
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

	// TODO normalize the rest of arguments

	return nil
}

func generateSecretKey(n int) string {
	// TODO add random key generation.
	return "+s6cv712&nw4gsk)1dmgpje+f#%^4lhp@!up+=p3ts+hxz(fr2"
}

func createProject() error {
	var (
		projectRoot string
		frameworkVerSubDir = ""
		frameworkVer string
		cppStandard int
		cmakeMinVer string
	)
	if *newProjectAskQuestions {
		// TODO: ask questions
	} else {
		projectRoot = *newProjectRoot
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
		ProjectName:               *newProjectName,
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
		path.Join(cfg.ProjectRoot, "lib"),
		cfg.FrameworkVersion,
	)
	if err != nil {
		return err
	}

	return nil
}
