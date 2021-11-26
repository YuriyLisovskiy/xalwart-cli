package new_

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core/components"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/templates"
	"github.com/iancoleman/strcase"
)

const moduleCommandDescription = `Create new module component.
Module files will have 'module' names by default and will be placed in the directory with snake case
name from value of 'name' flag.`

var moduleCommand = getComponentCommandBuilder("module", moduleCommandDescription).
	SetComponentBuilder(buildModuleComponent).
	SetPostRunMessageBuilder(moduleSuccess).
	Command(&overwriteVar)

func init() {
	initDefaultFlags("module", moduleCommand.Flags())
}

func buildModuleComponent() (core.Component, error) {
	header, err := getDefaultHeader()
	if err != nil {
		return nil, err
	}

	return components.NewModuleComponent(
		header,
		templates.ModuleTemplateBox,
		nameVar,
		rootPathVar,
		customFileNameVar,
	)
}

func moduleSuccess(component core.Component) string {
	module := component.(*components.ClassComponent)
	className := module.ClassName()
	nameSnake := strcase.ToSnake(module.Name())
	return fmt.Sprintf(
		`Success.

Register '%s' at the end of 'register_modules()' method in application settings:

  this->module<%s>();

Include '%s' in 'urlpatterns()' method of main module:

  this->include("%s", R"(%s)", "%s");

Do not forget to enable '%s' in configuration (yaml):

  modules:
    ...
    - %s
`, className, className, className, className, nameSnake, nameSnake, className, className,
	)
}
