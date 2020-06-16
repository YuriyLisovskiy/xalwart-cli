package generator

import (
	"github.com/gobuffalo/packd"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"xalwart-cli/config"
)

type Generator struct {

}

func (g *Generator) NewProject(cfg *config.Project) {
	var projectRoot string
	if len(cfg.ProjectName) == 0 {
		projectRoot = cfg.WorkingDirectory
	} else {
		projectRoot = path.Join(cfg.WorkingDirectory, cfg.ProjectName)
	}

	err := cfg.Templates.Walk(func(fp string, file packd.File) error {
		filePath, fileName := path.Split(fp)
		filePath = strings.Replace(filePath, "_app_", cfg.ProjectName, -1)
		filePath = path.Join(projectRoot, filePath)
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			panic(err)
		}

		stream, err := os.Create(path.Join(filePath, strings.TrimSuffix(fileName, filepath.Ext(fileName))))
		if err != nil {
			panic(err)
		}

		tmpl, err := template.New(fp).Funcs(config.DefaultFunctions).Delims("<%", "%>").Parse(file.String())
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

		return nil
	})

	if err != nil {
		panic(err)
	}
}
