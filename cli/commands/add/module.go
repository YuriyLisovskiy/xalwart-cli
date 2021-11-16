package add

import (
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
)

const moduleCommandLongDescription = ``

var moduleCommand = makeCommand(
	"module", moduleCommandLongDescription, func() (core.Component, error) {
		return components.NewModuleComponent(componentName, rootPath, componentCustomFileName)
	},
)

func init() {
	addCommonFlags("module", moduleCommand.Flags())
}
