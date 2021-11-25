package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core/components"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/templates"
)

const commandCommandDescription = `Create new command component.
Command files will have snake case value of 'name' flag as names by default.`

var commandCommand = getComponentCommandBuilder("command", commandCommandDescription).
	SetComponentBuilder(buildCommandComponent).
	SetPostRunMessageBuilder(commandSuccess).
	Command(&overwriteVar)

func init() {
	initDefaultFlags("command", commandCommand.Flags())
}

func buildCommandComponent() (core.Component, error) {
	header, err := getDefaultHeader()
	if err != nil {
		return nil, err
	}

	return components.NewCommandComponent(
		header,
		templates.CommandTemplateBox,
		nameVar,
		rootPathVar,
		customFileNameVar,
	)
}

func commandSuccess(component core.Component) string {
	className := component.(*components.ClassComponent).ClassName()
	return fmt.Sprintf(
		`Success.

Register '%s' at the end of 'commands()' method of the preferred module:
  
  this->command<%s>(this->settings->LOGGER);

If there is not 'commands()' method in application settings, overwrite it:

  // Declare public method of preferred module class in header file:
  void commands() override;

  // Define method in source file:
  void _MODULE_NAME_::commands()
  {
  }
`, className, className,
	)
}
