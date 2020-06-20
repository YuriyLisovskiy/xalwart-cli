package generator

import (
	"github.com/gobuffalo/packd"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"xalwart-cli/config"
	"xalwart-cli/utils"
)

func (g *Generator) projectCreateTemplate(fp string, file packd.File, cfg *config.Project) error {
	filePath, fileName := path.Split(fp)
	filePath = strings.Replace(filePath, "_proj_name_", cfg.ProjectName, -1)
	filePath = path.Join(cfg.ProjectRoot, filePath)
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return err
	}

	stream, err := os.Create(path.Join(filePath, strings.TrimSuffix(fileName, filepath.Ext(fileName))))
	if err != nil {
		return err
	}

	tmpl, err := template.New(fp).Funcs(config.DefaultFunctions).Delims("<%", "%>").Parse(file.String())
	if err != nil {
		return err
	}

	err = tmpl.Execute(stream, cfg)
	if err != nil {
		panic(err)
	}

	err = stream.Close()
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) NewProject(cfg *config.Project) error {
	err := cfg.Templates.Walk(func(fp string, file packd.File) error {
		return g.projectCreateTemplate(fp, file, cfg)
	})
	if err != nil {
		return err
	}

	err = utils.MakeDirs(cfg.ProjectRoot, []string{"media"})
	if err != nil {
		return err
	}

	return nil
}
