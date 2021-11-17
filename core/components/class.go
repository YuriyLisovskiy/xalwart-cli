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

func (s ClassComponent) Header() Header {
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

func newClassComponent(componentName, rootPath, componentType, customFileName string) (
	*ClassComponent,
	error,
) {
	commonComponent, err := newCommonComponent(componentType, componentName, rootPath)
	if err != nil {
		return nil, err
	}

	return &ClassComponent{
		common:         *commonComponent,
		componentType:  componentType,
		customFileName: customFileName,
	}, nil
}

func NewCommandComponent(command, rootPath, customFileName string) (*ClassComponent, error) {
	return newClassComponent(command, rootPath, "command", customFileName)
}

func NewControllerComponent(controllerName, rootPath, customFileName string) (*ClassComponent, error) {
	return newClassComponent(controllerName, rootPath, "controller", customFileName)
}

func NewModuleComponent(moduleName, rootPath, customFileName string) (*ClassComponent, error) {
	return newClassComponent(moduleName, rootPath, "module", customFileName)
}
