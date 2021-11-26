package core

import (
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

const (
	FrameworkName            = "xalwart"
	FrameworkNamespace       = "xw"
	FrameworkLatestVersion   = "0.0.0"
	FrameworkBaseDockerImage = FrameworkName + "/framework:" + FrameworkLatestVersion + "-alpine"

	AppName    = FrameworkName
	AppVersion = "0.1.0"

	CLIAppDocumentationLink = "https://github.com/YuriyLisovskiy/xalwart-cli/wiki"
)

var DefaultFunctions = template.FuncMap{
	"upper":         strings.ToUpper,
	"lower":         strings.ToLower,
	"to_camel_case": strcase.ToCamel,
	"to_snake_case": strcase.ToSnake,
}
