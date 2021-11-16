package add

import (
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
)

const controllerCommandLongDescription = `Create new controller component.

Controller files will have '{lower_case_name}_controller' names by default.`

var controllerCommand = makeCommand(
	"controller", controllerCommandLongDescription, func() (core.Component, error) {
		return components.NewControllerComponent(componentName, rootPath, componentCustomFileName)
	},
)

func init() {
	addCommonFlags("controller", controllerCommand.Flags())
}
