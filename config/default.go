package config

import (
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

const (
	AppName = "xalwart"
	AppVersion = "0.0.1"

	FrameworkName       = "xalwart"
	FrameworkNamespace  = "xw"
)

var DefaultFunctions = template.FuncMap {
	"upper": strings.ToUpper,
	"to_camel_case": strcase.ToCamel,
	"to_snake_case": strcase.ToSnake,
}
