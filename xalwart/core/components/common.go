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
	filePath, fileName := path.Split(templatePath)
	filePath = path.Join(component.RootPath(), filePath)
	filePath = strings.ReplaceAll(filePath, "_name_", strcase.ToSnake(component.Name()))
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	fileName = strings.ReplaceAll(fileName, "_name_", strcase.ToSnake(component.FileName()))
	return path.Join(filePath, fileName)
}
