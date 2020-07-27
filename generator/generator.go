package generator

import (
	"errors"
	"github.com/gobuffalo/packd"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"xalwart-cli/config"
	"xalwart-cli/utils"
)

type Generator struct {
	UnitExists func (cfg *ProjectUnit) error
	EmptyDirsToCreateInUnit []string

	FilePathSetup func (fp string, fn string) (string, string)
	ErrorIfFileExists func () error

	CheckIfNameIsSet bool
}

func (g *Generator) executeTemplate(fp string, file packd.File, cfg *ProjectUnit) error {
	filePath, fileName := path.Split(fp)
	if g.FilePathSetup != nil {
		filePath, fileName = g.FilePathSetup(filePath, fileName)
	}

	filePath = path.Join(cfg.Root, filePath)
	fullPath := path.Join(filePath, strings.TrimSuffix(fileName, filepath.Ext(fileName)))
	if utils.FileExists(fullPath) && g.ErrorIfFileExists != nil {
		return g.ErrorIfFileExists()
	}

	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return err
	}

	stream, err := os.Create(fullPath)
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

func (g *Generator) NewUnit(cfg *ProjectUnit, unitName string) error {
	cfg.Init()
	if g.CheckIfNameIsSet && len(cfg.Name) == 0 {
		return errors.New("name of a new " + unitName + " is now set")
	}

	if g.UnitExists != nil {
		if err := g.UnitExists(cfg); err != nil {
			return err
		}
	}

	err := os.MkdirAll(cfg.Root, os.ModePerm)
	if err != nil {
		return err
	}

	err = cfg.Templates.Walk(func(fp string, file packd.File) error {
		return g.executeTemplate(fp, file, cfg)
	})
	if err != nil {
		return err
	}

	err = utils.MakeDirs(cfg.Root, g.EmptyDirsToCreateInUnit)
	if err != nil {
		return err
	}

	return nil
}
