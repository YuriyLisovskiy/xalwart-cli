package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/config"
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
)

var (
	modelCustomTableName    string
	modelIsJsonSerializable = false
)

const modelCommandLongDescription = ``

var modelCommand = makeCommand(
	"model", modelCommandLongDescription, func() (core.Component, error) {
		return components.NewModelComponent(
			componentName,
			rootPath,
			componentCustomFileName,
			modelCustomTableName,
			modelIsJsonSerializable,
		)
	},
)

func init() {
	flags := modelCommand.Flags()
	addCommonFlags("model", flags)
	flags.StringVarP(
		&modelCustomTableName,
		"table",
		"t",
		"",
		"custom name which will be used for table in SQL database (flag 'name' will be used by default)",
	)
	flags.BoolVarP(
		&modelIsJsonSerializable,
		"json-serializable",
		"j",
		modelIsJsonSerializable,
		fmt.Sprintf(
			"inherit from '%s::IJsonSerializable' interface and add implement 'to_json()' method",
			config.FrameworkNamespace,
		),
	)
}
