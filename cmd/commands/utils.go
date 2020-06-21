package commands

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"xalwart-cli/config"
)

const metaFilePath = "." + config.FrameworkName + ".meta"

type projectMeta struct {
	Version string      `json:"version"`
	ProjectName string  `json:"project_name"`
}

func loadMeta(projectRoot string) (projectMeta, error) {
	obj := projectMeta{Version: "latest"}
	jsonFile, err := os.Open(path.Join(projectRoot, metaFilePath))
	if err != nil {
		return obj, err
	}

	defer jsonFile.Close()
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return obj, err
	}

	err = json.Unmarshal(bytes, &obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
