package components

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/gobuffalo/packr/v2"
	"github.com/iancoleman/strcase"
)

type CommonComponent struct {
	header      Header
	name        string
	rootPath    string
	templateBox core.TemplateBox
}

func newCommonComponent(templateBoxName, componentName, rootPath string) (*CommonComponent, error) {
	noticeName := "copyright_notice.txt"
	headerBox := packr.New(noticeName, "../../templates")
	noticeContent, err := headerBox.FindString("copyright_notice.txt")
	if err != nil {
		return nil, err
	}

	header, err := newHeader(noticeContent)
	if err != nil {
		return nil, err
	}

	return &CommonComponent{
		header:      *header,
		name:        componentName,
		rootPath:    rootPath,
		templateBox: core.NewFileTemplateBox(templateBoxName),
	}, nil
}

func getTargetPath(templatePath string, component core.Component) string {
	filePath, fileName := path.Split(templatePath)
	filePath = path.Join(component.RootPath(), filePath)
	filePath = strings.ReplaceAll(filePath, "_name_", strcase.ToSnake(component.Name()))
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	fileName = strings.ReplaceAll(fileName, "_name_", strcase.ToSnake(component.FileName()))
	return path.Join(filePath, fileName)
}
