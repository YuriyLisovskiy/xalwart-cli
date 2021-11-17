package add

import (
	"fmt"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
	"github.com/YuriyLisovskiy/xalwart-cli/core/components"
)

var middlewareIsClassBased = false

const middlewareCommandLongDescription = `Create new middleware component.
Middleware files will have lowercase '{name}' names by default.`

var middlewareCommand = makeCommand(
	"middleware",
	middlewareCommandLongDescription,
	func() (core.Component, error) {
		return components.NewMiddlewareComponent(componentName, rootPath, componentCustomFileName, middlewareIsClassBased)
	},
	func(component core.Component) string {
		middleware := component.(*components.MiddlewareComponent)
		fullName := middleware.FullName()
		initializationName := fullName
		if middleware.IsClassBased() {
			initializationName += "()"
		}

		return fmt.Sprintf(`Success.

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
`, fullName, fullName, initializationName, fullName, fullName)
	},
)

func init() {
	flags := middlewareCommand.Flags()
	addCommonFlags("middleware", flags)
	flags.BoolVarP(
		&middlewareIsClassBased,
		"class-based",
		"c",
		middlewareIsClassBased,
		"add class-based middleware instead of function-based",
	)
}
