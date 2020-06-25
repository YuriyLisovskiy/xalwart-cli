package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"regexp"
	"time"
	"xalwart-cli/config"
	"xalwart-cli/generator"
	"xalwart-cli/managers"
)

const metaFilePath = ".project." + config.FrameworkNamespace

type projectMeta struct {
	FrameworkVersion string  `json:"framework_version"`
	ProjectName string       `json:"project_name"`
}

type Cmd struct {
	customizeUnit func (cwd string, unit *generator.ProjectUnit) error
	makeGenerator func (unit *generator.ProjectUnit) generator.Generator
	postProcess func (unit *generator.ProjectUnit) error
	postCreateHelp func (unit *generator.ProjectUnit)
}

func (c *Cmd) execute(unitName string, isProject bool) error {
	fmt.Printf("Generating %s...", unitName)

	usr, err := user.Current()
	if err != nil {
		return err
	}

	cfg := generator.ProjectUnit{
		Year:                      time.Now().Year(),
		Username:                  usr.Name,
		FrameworkName:             config.FrameworkName,
		FrameworkNamespace:        config.FrameworkNamespace,
	}
	if c.customizeUnit != nil {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		err = c.customizeUnit(cwd, &cfg)
		if err != nil {
			return err
		}
	}

	var initFWVer string
	if isProject {
		initFWVer = cfg.FrameworkVersion
	} else {
		meta, err := c.loadMeta(cfg.ProjectRoot)
		if err != nil {
			return err
		}

		cfg.ProjectName = meta.ProjectName
		initFWVer = meta.FrameworkVersion
	}

	if cfg.InstallFramework {
		cfg.FrameworkVersion, err = c.getVersionOfFramework(
			initFWVer, isProject,
		)
		if isProject && err != nil {
			return err
		}
	} else {
		cfg.FrameworkVersion = "null"
	}

	var g generator.Generator
	if c.makeGenerator != nil {
		g = c.makeGenerator(&cfg)
	}

	err = g.NewUnit(&cfg, unitName)
	if err != nil {
		return err
	}

	fmt.Println(" Done.")

	if c.postProcess != nil {
		err := c.postProcess(&cfg)
		if err != nil {
			return err
		}
	}

	if c.postCreateHelp != nil {
		c.postCreateHelp(&cfg)
	}

	return nil
}

func (c *Cmd) loadMeta(projectRoot string) (projectMeta, error) {
	projectErr := errors.New("unable to read project meta: '" + metaFilePath + "' is missing or damaged")
	obj := projectMeta{FrameworkVersion: "latest"}
	jsonFile, err := os.Open(path.Join(projectRoot, metaFilePath))
	if err != nil {
		return obj, projectErr
	}

	defer jsonFile.Close()
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return obj, projectErr
	}

	err = json.Unmarshal(bytes, &obj)
	if err != nil {
		return obj, projectErr
	}

	return obj, nil
}

func (c *Cmd) saveMeta(projectRoot string, meta projectMeta) error {
	savingErr := errors.New("unable to update project meta, file: '" + metaFilePath + "'")
	content, err := json.MarshalIndent(meta, "", " ")
	if err != nil {
		return savingErr
	}

	err = ioutil.WriteFile(path.Join(projectRoot, metaFilePath), content, 0644)
	if err != nil {
		return savingErr
	}

	return nil
}

func (c *Cmd) reSubMatchMap(r *regexp.Regexp, str string) (bool, map[string]string) {
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

func (c *Cmd) getVersionOfFramework(version string, verbose bool) (string, error) {
	retrieveErr := errors.New(
		"unable to retrieve latest release of '" + config.FrameworkName + "' framework",
	)
	if version == "latest" {
		var err error
		release, err := managers.GetLatestRelease()
		if err != nil {
			return "", retrieveErr
		}

		version = release.VersionTag
	} else {
		isAvailable, err := managers.CheckIfVersionIsAvailable(version)
		if err != nil {
			return "", retrieveErr
		}

		if !isAvailable {
			if verbose {
				fmt.Println(
					"\nWarning: release v" + version + " is not available, latest is used instead",
				)
			}

			release, err := managers.GetLatestRelease()
			if err != nil {
				return "", retrieveErr
			}

			version = release.VersionTag
		}
	}

	return version, nil
}
