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

// TODO: search for `Settings::register_apps()` method in `settings.cpp` file;
//      append a new application here and to `installed_apps` in `config.yml`.

func (g *Generator) appCreateTemplate(fp string, file packd.File, cfg *config.App) error {
	filePath, fileName := path.Split(fp)
	filePath = path.Join(cfg.AppRoot, filePath)
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

func (g *Generator) NewApp(cfg *config.App) error {
	if len(cfg.AppName) == 0 {
		return errors.New("name of a new application is empty")
	}

	if !strings.HasSuffix(strings.ToLower(cfg.AppName), "_app") {
		cfg.AppName += "_app"
	}

	cfg.AppRoot = path.Join(cfg.ProjectRoot, cfg.AppName)
	if _, err := os.Stat(cfg.AppRoot); !os.IsNotExist(err) {
		return errors.New("'" + cfg.AppName + "' application already exists")
	}

	err := os.MkdirAll(cfg.AppRoot, os.ModePerm)
	if err != nil {
		return err
	}

	err = cfg.Templates.Walk(func(fp string, file packd.File) error {
		return g.appCreateTemplate(fp, file, cfg)
	})
	if err != nil {
		return err
	}

	err = utils.MakeDirs(cfg.AppRoot, []string{"views"})
	if err != nil {
		return err
	}

	return nil
}
