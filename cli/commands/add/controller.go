package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
	"github.com/iancoleman/strcase"
)

const controllerCommandLongDescription = `Create new controller component.
Controller files will have lowercase '{name}_controller' names by default.`

var controllerCommand = makeCommand(
	"controller",
	controllerCommandLongDescription,
	func() (core.Component, error) {
		return components.NewControllerComponent(componentName, rootPath, componentCustomFileName)
	},
	func(component core.Component) string {
		command := component.(*components.ClassComponent)
		className := command.ClassName()
		nameSnake := strcase.ToSnake(command.Name())
		return fmt.Sprintf(`Success.

Register '%s' in 'urlpatterns()' method of the preferred module:
  
  this->url<%s>(R"(%s/?)", "%s");
`, className, className, nameSnake, nameSnake)
	},
)

func init() {
	addCommonFlags("controller", controllerCommand.Flags())
}
