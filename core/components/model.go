package components

import (
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

type ModelComponent struct {
	class              ClassComponent
	isJsonSerializable bool
	customTableName    string
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

func (m ModelComponent) Header() core.Header {
	return m.class.Header()
}

func (m ModelComponent) ClassName() string {
	return m.class.ClassName()
}

func (m ModelComponent) IsJsonSerializable() bool {
	return m.isJsonSerializable
}

func (m ModelComponent) TableName() string {
	if len(m.customTableName) != 0 {
		return m.customTableName
	}

	return pluralize.NewClient().Plural(m.Name())
}

func NewModelComponent(
	header core.Header,
	templateBox core.TemplateBox,
	modelName string,
	rootPath string,
	customFileName string,
	customTableName string,
	isJsonSerializable bool,
) (
	*ModelComponent,
	error,
) {
	classComponent, err := newClassComponent(header, templateBox, modelName, rootPath, "model", customFileName)
	if err != nil {
		return nil, err
	}

	return &ModelComponent{
		class:              *classComponent,
		isJsonSerializable: isJsonSerializable,
		customTableName:    customTableName,
	}, nil
}
