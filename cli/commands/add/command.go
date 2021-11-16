package add

import (
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
)

const commandCommandLongDescription = ``

var commandCommand = makeCommand(
	"command", commandCommandLongDescription, func() (core.Component, error) {
		return components.NewCommandComponent(componentName, rootPath, componentCustomFileName)
	},
)

func init() {
	addCommonFlags("command", commandCommand.Flags())
}
