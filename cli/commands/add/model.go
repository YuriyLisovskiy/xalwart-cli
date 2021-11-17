package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
)

var (
	modelCustomTableName    string
	modelIsJsonSerializable = false
)

const modelCommandLongDescription = `Create new model component.
Model files will have lowercase '{name}' names by default.`

var modelCommand = makeCommand(
	"model",
	modelCommandLongDescription,
	func() (core.Component, error) {
		return components.NewModelComponent(
			componentName,
			rootPath,
			componentCustomFileName,
			modelCustomTableName,
			modelIsJsonSerializable,
		)
	}, func(component core.Component) string {
		return fmt.Sprintf(`Success.

Do not forget to create a new migration for '%s' and apply changes to the database.
`, component.(*components.ModelComponent).ClassName())
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
		"custom name which will be used for table in SQL database (value of 'name' flag will be used by default)",
	)
	flags.BoolVarP(
		&modelIsJsonSerializable,
		"json-serializable",
		"j",
		modelIsJsonSerializable,
		fmt.Sprintf(
			"inherit from '%s::IJsonSerializable' interface and implement 'to_json()' method",
			core.FrameworkNamespace,
		),
	)
}
