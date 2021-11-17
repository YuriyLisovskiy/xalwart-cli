package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/cli/utils"
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
	"github.com/iancoleman/strcase"
)

const moduleCommandDescription = `Create new module component.
Module files will have 'module' names by default and will be placed in lowercase '{name}' directory.`

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
		utils.GetModuleTemplateBox(),
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
