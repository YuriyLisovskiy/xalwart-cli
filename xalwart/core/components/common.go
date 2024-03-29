package components

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
	"github.com/iancoleman/strcase"
)

type CommonComponent struct {
	header      core.Header
	name        string
	rootPath    string
	templateBox core.TemplateBox
}

func getTargetPath(templatePath string, component core.Component) string {
	if !path.IsAbs(templatePath) {
		templatePath = path.Join(component.RootPath(), templatePath)
	}

	templatePath = strings.TrimSuffix(templatePath, filepath.Ext(templatePath))
	filePath, fileName := path.Split(templatePath)
	filePath = strings.ReplaceAll(filePath, "_name_", strcase.ToSnake(component.Name()))
	fileName = strings.ReplaceAll(fileName, "_name_", strcase.ToSnake(component.FileName()))
	return path.Join(filePath, fileName)
}
