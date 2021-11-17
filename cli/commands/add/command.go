package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
)

const commandCommandLongDescription = `Create new command component.
Command files will have {lower_case_name} names by default.`

var commandCommand = makeCommand(
	"command",
	commandCommandLongDescription,
	func() (core.Component, error) {
		return components.NewCommandComponent(componentName, rootPath, componentCustomFileName)
	},
	func(component core.Component) string {
		className := component.(*components.ClassComponent).ClassName()
		return fmt.Sprintf(`Success.

Register '%s' at the end of 'commands()' method of the preferred module:
  
  this->command<%s>(this->settings->LOGGER);

If there is not 'commands()' method in application settings, overwrite it:

  // Declare public method of preferred module class in header file:
  void commands() override;

  // Define method in source file:
  void _MODULE_NAME_::commands()
  {
  }
`, className, className)
	},
)

func init() {
	addCommonFlags("command", commandCommand.Flags())
}
