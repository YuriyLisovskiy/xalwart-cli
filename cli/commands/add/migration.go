package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
)

var migrationIsInitial = false

const migrationCommandLongDescription = `Create new migration component.
Migration files will have lower-case migration name by default.

Recommended migration name structure is '{number}_{ShortDescription}', for example: '001_Initial'.
In this case, migration class will have 'Migration{number}_{ShortDescription}' name.`

var migrationCommand = makeCommand(
	"migration",
	migrationCommandLongDescription,
	func() (core.Component, error) {
		return components.NewMigrationComponent(componentName, rootPath, componentCustomFileName, migrationIsInitial)
	},
	func(component core.Component) string {
		className := component.(*components.MigrationComponent).ClassName()
		return fmt.Sprintf(`Success.

Register '%s' at the end of 'register_migrations()' method in application settings:
  
  this->migration<%s>();

If there is not 'register_migrations()' method in application settings, overwrite it:

  // Declare public method of 'Settings' class in header file:
  void register_migrations() override;

  // Define method in source file:
  void Settings::register_migrations()
  {
  }
`, className, className)
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
