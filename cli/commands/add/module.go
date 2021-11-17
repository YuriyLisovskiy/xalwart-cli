package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
	"github.com/iancoleman/strcase"
)

const moduleCommandLongDescription = `Create new module component.
Module files will have 'module' names by default and will be placed in {name} directory.`

var moduleCommand = makeCommand(
	"module",
	moduleCommandLongDescription,
	func() (core.Component, error) {
		return components.NewModuleComponent(componentName, rootPath, componentCustomFileName)
	},
	func(component core.Component) string {
		module := component.(*components.ClassComponent)
		className := module.ClassName()
		nameSnake := strcase.ToSnake(module.Name())
		return fmt.Sprintf(`Success.

Register '%s' at the end of 'register_modules()' method in application settings:

  this->module<%s>();

Include '%s' in 'urlpatterns()' method of main module:

  this->include("%s", R"(%s)", "%s");

Do not forget to enable '%s' in configuration (yaml):

  modules:
    ...
    - %s
`, className, className, className, className, nameSnake, nameSnake, className, className)
	},
)

func init() {
	addCommonFlags("module", moduleCommand.Flags())
}
