package core

import (
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

const (
	FrameworkName      = "xalwart"
	FrameworkNamespace = "xw"

	AppName    = FrameworkName
	AppVersion = "0.1.0"

	CLIAppDocumentationLink = "https://github.com/YuriyLisovskiy/xalwart-cli/wiki/Version-" + AppVersion
)

var DefaultFunctions = template.FuncMap{
	"upper":         strings.ToUpper,
	"to_camel_case": strcase.ToCamel,
	"to_snake_case": strcase.ToSnake,
}
