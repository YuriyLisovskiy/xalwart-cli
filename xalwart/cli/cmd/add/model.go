package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/cli/utils"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core/components"
)

var (
	modelCustomTableName    string
	modelIsJsonSerializable = false
)

const modelCommandDescription = `Create new model component.
Model files will have snake case value of 'name' flag as names by default.`

var modelCommand = getComponentCommandBuilder("model", modelCommandDescription).
	SetComponentBuilder(buildModelComponent).
	SetPostRunMessageBuilder(modelSuccess).
	Command(&overwriteVar)

func init() {
	flags := modelCommand.Flags()
	initDefaultFlags("model", flags)
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

func buildModelComponent() (core.Component, error) {
	header, err := getDefaultHeader()
	if err != nil {
		return nil, err
	}

	return components.NewModelComponent(
		header,
		utils.GetModelTemplateBox(),
		nameVar,
		rootPathVar,
		customFileNameVar,
		modelCustomTableName,
		modelIsJsonSerializable,
	)
}

func modelSuccess(component core.Component) string {
	return fmt.Sprintf(`Success.

Do not forget to create a new migration for '%s' and apply changes to the database.
`, component.(*components.ModelComponent).ClassName())
}
