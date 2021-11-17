package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/cli/utils"
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
	"github.com/iancoleman/strcase"
)

const controllerCommandDescription = `Create new controller component.
Controller files will have lowercase '{name}_controller' names by default.`

var controllerCommand = getComponentCommandBuilder("controller", controllerCommandDescription).
	SetComponentBuilder(buildControllerComponent).
	SetPostRunMessageBuilder(controllerSuccess).
	Command(&overwriteVar)

func init() {
	initDefaultFlags("controller", controllerCommand.Flags())
}

func buildControllerComponent() (core.Component, error) {
	header, err := getDefaultHeader()
	if err != nil {
		return nil, err
	}

	return components.NewControllerComponent(
		header,
		utils.GetControllerTemplateBox(),
		nameVar,
		rootPathVar,
		customFileNameVar,
	)
}

func controllerSuccess(component core.Component) string {
	command := component.(*components.ClassComponent)
	className := command.ClassName()
	nameSnake := strcase.ToSnake(command.Name())
	return fmt.Sprintf(
		`Success.

Register '%s' in 'urlpatterns()' method of the preferred module:
  
  this->url<%s>(R"(%s/?)", "%s");
`, className, className, nameSnake, nameSnake,
	)
}
