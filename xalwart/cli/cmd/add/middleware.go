package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/cli/utils"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core/components"
)

var middlewareIsClassBased = false

const middlewareCommandDescription = `Create new middleware component.
Middleware files will have snake case value of 'name' flag as names by default.`

var middlewareCommand = getComponentCommandBuilder("middleware", middlewareCommandDescription).
	SetComponentBuilder(buildMiddlewareComponent).
	SetPostRunMessageBuilder(middlewareSuccess).
	Command(&overwriteVar)

func init() {
	flags := middlewareCommand.Flags()
	initDefaultFlags("middleware", flags)
	flags.BoolVarP(
		&middlewareIsClassBased,
		"class-based",
		"c",
		middlewareIsClassBased,
		"add class-based middleware instead of function-based",
	)
}

func buildMiddlewareComponent() (core.Component, error) {
	header, err := getDefaultHeader()
	if err != nil {
		return nil, err
	}

	return components.NewMiddlewareComponent(
		header,
		utils.GetMiddlewareTemplateBox(),
		nameVar,
		rootPathVar,
		customFileNameVar,
		middlewareIsClassBased,
	)
}

func middlewareSuccess(component core.Component) string {
	middleware := component.(*components.MiddlewareComponent)
	fullName := middleware.FullName()
	initializationName := fullName
	if middleware.IsClassBased() {
		initializationName += "()"
	}

	return fmt.Sprintf(
		`Success.

Register '%s' in 'register_middleware()' method in application settings:
  
  this->middleware("%s", %s);

If there is not 'register_middleware()' method in application settings, overwrite it:

  // Declare public method of 'Settings' class in header file:
  void register_middleware() override;

  // Define method in source file:
  void Settings::register_middleware()
  {
  }

Do not forget to enable '%s' in configuration (yaml):

  middleware:
    ...
    - %s
    ...
`, fullName, fullName, initializationName, fullName, fullName,
	)
}
