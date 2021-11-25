package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core/components"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/templates"
	"github.com/iancoleman/strcase"
)

const controllerCommandDescription = `Create new controller component.
Controller files will have snake case '{name_flag}_controller' names by default.`

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
		templates.ControllerTemplateBox,
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
