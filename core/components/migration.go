package components

import (
	"fmt"
	"regexp"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/iancoleman/strcase"
)

type MigrationComponent struct {
	class                     ClassComponent
	isInitial                 bool
	migrationShortDescription string
	className                 string
}

func (m MigrationComponent) Name() string {
	return m.class.Name()
}

func (m MigrationComponent) FileName() string {
	if len(m.class.customFileName) != 0 {
		return m.class.customFileName
	}

	return strcase.ToSnake(m.MigrationShortDescription())
}

func (m MigrationComponent) RootPath() string {
	return m.class.RootPath()
}

func (m MigrationComponent) GetTargetPath(templatePath string) string {
	return getTargetPath(templatePath, m)
}

func (m MigrationComponent) TemplateBox() core.TemplateBox {
	return m.class.TemplateBox()
}

func (m MigrationComponent) Header() Header {
	return m.class.Header()
}

func (m MigrationComponent) ClassName() string {
	return m.className
}

func (m MigrationComponent) IsInitial() bool {
	return m.isInitial
}

func (m MigrationComponent) MigrationShortDescription() string {
	return m.migrationShortDescription
}

func NewMigrationComponent(migrationName, rootPath, customFileName string, isInitial bool) (
	*MigrationComponent,
	error,
) {
	className := migrationName
	matched, err := regexp.MatchString(`\d+_[a-zA-Z0-9_]+`, migrationName)
	if err != nil {
		return nil, err
	}

	if matched {
		className = fmt.Sprintf("Migration%s", className)
	}

	classComponent, err := newClassComponent("migration", className, rootPath, "migration", customFileName)
	if err != nil {
		return nil, err
	}

	return &MigrationComponent{
		class:                     *classComponent,
		isInitial:                 isInitial,
		migrationShortDescription: migrationName,
		className:                 className,
	}, nil
}
