package components

import (
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/iancoleman/strcase"
)

type ModelComponent struct {
	class ClassComponent
}

func (m ModelComponent) Name() string {
	return m.class.Name()
}

func (m ModelComponent) FileName() string {
	if len(m.class.customFileName) != 0 {
		return m.class.customFileName
	}

	return strcase.ToSnake(m.Name())
}

func (m ModelComponent) RootPath() string {
	return m.class.RootPath()
}

func (m ModelComponent) GetTargetPath(templatePath string) string {
	return getTargetPath(templatePath, m)
}

func (m ModelComponent) TemplateBox() core.TemplateBox {
	return m.class.TemplateBox()
}

func (m ModelComponent) Header() Header {
	return m.class.Header()
}

func (m ModelComponent) ClassName() string {
	return m.class.ClassName()
}

func (m ModelComponent) WithId() bool {
	// TODO:
	return false
}

func (m ModelComponent) IsJsonSerializable() bool {
	// TODO:
	return false
}

func (m ModelComponent) TableName() string {
	// TODO:
	return ""
}

func NewModelComponent(modelName, rootPath, customFileName string) (
	*ModelComponent,
	error,
) {
	classComponent, err := newClassComponent("model", modelName, rootPath, "model", customFileName)
	if err != nil {
		return nil, err
	}

	return &ModelComponent{
		class: *classComponent,
	}, nil
}
