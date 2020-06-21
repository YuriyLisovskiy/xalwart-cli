package commands

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"xalwart-cli/config"
)

const metaFilePath = "." + config.FrameworkName + ".project"

type projectMeta struct {
	FrameworkVersion string  `json:"framework_version"`
	ProjectName string       `json:"project_name"`
}

func loadMeta(projectRoot string) (projectMeta, error) {
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
