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
)

const metaFilePath = ".project." + config.FrameworkNamespace

type projectMeta struct {
	ProjectName string `json:"project_name"`
}

type Cmd struct {
	process func (cwd string, unit *generator.ProjectUnit) error
	postProcess func (unit *generator.ProjectUnit) error
}

func (c *Cmd) execute(unitName string, isProject bool) error {
	fmt.Printf("Generating %s...", unitName)

	usr, err := user.Current()
	if err != nil {
		return err
	}

	cfg := generator.ProjectUnit{
		Year:               time.Now().Year(),
		Username:           usr.Name,
		FrameworkName:      config.FrameworkName,
		FrameworkNamespace: config.FrameworkNamespace,
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	if !isProject {
		meta, err := c.loadMeta(cfg.ProjectRoot)
		if err != nil {
			return err
		}

		cfg.ProjectName = meta.ProjectName
	}

	if c.process != nil {
		err = c.process(cwd, &cfg)
		if err != nil {
			return err
		}
	}

	fmt.Println(" Done.")

	if c.postProcess != nil {
		err := c.postProcess(&cfg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Cmd) buildPath(initialPath string) string {
	return initialPath
}

func (c *Cmd) loadMeta(projectRoot string) (projectMeta, error) {
	projectErr := errors.New("unable to read project meta: '" + metaFilePath + "' is missing or damaged")
	obj := projectMeta{}
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
