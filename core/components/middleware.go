package components

import (
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/iancoleman/strcase"
)

type MiddlewareComponent struct {
	class        ClassComponent
	isClassBased bool
}

func (m MiddlewareComponent) Name() string {
	return m.class.Name()
}

func (m MiddlewareComponent) FileName() string {
	if len(m.class.customFileName) != 0 {
		return m.class.customFileName
	}

	return strcase.ToSnake(m.Name())
}

func (m MiddlewareComponent) RootPath() string {
	return m.class.RootPath()
}

func (m MiddlewareComponent) GetTargetPath(templatePath string) string {
	return getTargetPath(templatePath, m)
}

func (m MiddlewareComponent) TemplateBox() core.TemplateBox {
	return m.class.TemplateBox()
}

func (m MiddlewareComponent) Header() Header {
	return m.class.Header()
}

func (m MiddlewareComponent) FullName() string {
	className := m.class.ClassName()
	if m.IsClassBased() {
		return strcase.ToCamel(className)
	}

	return strcase.ToSnake(className)
}

func (m MiddlewareComponent) IsClassBased() bool {
	return m.isClassBased
}

func NewMiddlewareComponent(modelName, rootPath, customFileName string, isClassBased bool) (
	*MiddlewareComponent,
	error,
) {
	classComponent, err := newClassComponent(modelName, rootPath, "middleware", customFileName)
	if err != nil {
		return nil, err
	}

	return &MiddlewareComponent{
		class:        *classComponent,
		isClassBased: isClassBased,
	}, nil
}
