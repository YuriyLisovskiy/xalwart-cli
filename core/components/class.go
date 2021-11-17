package components

import (
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/iancoleman/strcase"
)

type ClassComponent struct {
	common         CommonComponent
	componentType  string
	customFileName string
}

func (s ClassComponent) Header() core.Header {
	return s.common.header
}

func (s ClassComponent) ClassName() string {
	return strcase.ToCamel(s.common.name) + strcase.ToCamel(s.componentType)
}

func (s ClassComponent) Name() string {
	return s.common.name
}

func (s ClassComponent) FileName() string {
	if len(s.customFileName) != 0 {
		return s.customFileName
	}

	return strcase.ToSnake(s.ClassName())
}

func (s ClassComponent) RootPath() string {
	return s.common.rootPath
}

func (s ClassComponent) GetTargetPath(templatePath string) string {
	return getTargetPath(templatePath, s)
}

func (s ClassComponent) TemplateBox() core.TemplateBox {
	return s.common.templateBox
}

func newClassComponent(
	header core.Header,
	templateBox core.TemplateBox,
	componentName string,
	rootPath string,
	componentType string,
	customFileName string,
) (
	*ClassComponent,
	error,
) {
	return &ClassComponent{
		common: CommonComponent{
			header:      header,
			name:        componentName,
			rootPath:    rootPath,
			templateBox: templateBox,
		},
		componentType:  componentType,
		customFileName: customFileName,
	}, nil
}

func NewCommandComponent(
	header core.Header,
	templateBox core.TemplateBox,
	command string,
	rootPath string,
	customFileName string,
) (*ClassComponent, error) {
	return newClassComponent(header, templateBox, command, rootPath, "command", customFileName)
}

func NewControllerComponent(
	header core.Header,
	templateBox core.TemplateBox,
	controllerName string,
	rootPath string,
	customFileName string,
) (*ClassComponent, error) {
	return newClassComponent(header, templateBox, controllerName, rootPath, "controller", customFileName)
}

func NewModuleComponent(
	header core.Header,
	templateBox core.TemplateBox,
	moduleName string,
	rootPath string,
	customFileName string,
) (*ClassComponent, error) {
	return newClassComponent(header, templateBox, moduleName, rootPath, "module", customFileName)
}
