package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/cli/utils"
	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
)

var migrationIsInitial = false

const migrationCommandDescription = `Create new migration component.
Migration files will have snake case value of 'name' flag as names by default.

Recommended migration name structure is '{number}_{ShortDescription}', for example: '001_Initial'.
In this case, migration class will have 'Migration{number}_{ShortDescription}' name.`

var migrationCommand = getComponentCommandBuilder("migration", migrationCommandDescription).
	SetComponentBuilder(buildMigrationComponent).
	SetPostRunMessageBuilder(migrationSuccess).
	Command(&overwriteVar)

func init() {
	flags := migrationCommand.Flags()
	initDefaultFlags("migration", flags)
	flags.BoolVarP(&migrationIsInitial, "initial", "i", migrationIsInitial, "mark migration as initial one")
}

func buildMigrationComponent() (core.Component, error) {
	header, err := getDefaultHeader()
	if err != nil {
		return nil, err
	}

	return components.NewMigrationComponent(
		header,
		utils.GetMigrationTemplateBox(),
		nameVar,
		rootPathVar,
		customFileNameVar,
		migrationIsInitial,
	)
}

func migrationSuccess(component core.Component) string {
	className := component.(*components.MigrationComponent).ClassName()
	return fmt.Sprintf(
		`Success.

Register '%s' at the end of 'register_migrations()' method in application settings:
  
  this->migration<%s>();

If there is not 'register_migrations()' method in application settings, overwrite it:

  // Declare public method of 'Settings' class in header file:
  void register_migrations() override;

  // Define method in source file:
  void Settings::register_migrations()
  {
  }
`, className, className,
	)
}
