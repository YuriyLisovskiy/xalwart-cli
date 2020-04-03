package internal

import (
	"os"
	"path"
	"text/template"
	"wasp-cli/internal/config"
)

func GenerateProject(cfg *config.ProjectConfig) {
	var projectRoot string
	if len(cfg.ProjectName) == 0 {
		projectRoot = cfg.WorkingDirectory
	} else {
		projectRoot = path.Join(cfg.WorkingDirectory, cfg.ProjectName)
	}

	for _, file := range config.ProjectTemplates {
		if file.Path == "PROJECT_NAME" {
			file.Path = cfg.ProjectName
		}

		filePath := path.Join(projectRoot, file.Path)
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			panic(err)
		}

		stream, err := os.Create(path.Join(filePath, file.Name))
		if err != nil {
			panic(err)
		}

		tmpl, err := template.New(file.Name).Funcs(config.DefaultFunctions).Parse(file.TemplateStr)
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(stream, cfg)
		if err != nil {
			panic(err)
		}

		err = stream.Close()
		if err != nil {
			panic(err)
		}
	}
}
