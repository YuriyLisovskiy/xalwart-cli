package add

import (
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
)

var migrationIsInitial = false

const migrationCommandLongDescription = `Recommended migration name structure is 'Migration{number}_{ShortDescription}'.
Example: 'Migration001_Initial'.

In case of using recommended name structure, migration file name will have '{number}_{ShortDescription}' name by default.`

var migrationCommand = makeCommand(
	"migration", migrationCommandLongDescription, func() (core.Component, error) {
		return components.NewMigrationComponent(componentName, rootPath, componentCustomFileName, migrationIsInitial)
	},
)

func init() {
	addCommonFlags("migration", migrationCommand.Flags())
	migrationCommand.Flags().BoolVarP(
		&migrationIsInitial,
		"initial",
		"i",
		migrationIsInitial,
		"mark migration as initial one",
	)
}
