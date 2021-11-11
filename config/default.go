package config

import (
	"github.com/iancoleman/strcase"
	"strings"
	"text/template"
)

const (
	AppName = "xalwart"
	AppVersion = "0.0.1"

	FrameworkName       = "xalwart"
	FrameworkNamespace  = "xw"

	SecretKeyLength = 50
)

var DefaultFunctions = template.FuncMap {
	"upper": strings.ToUpper,
	"to_camel_case": strcase.ToCamel,
	"to_snake_case": strcase.ToSnake,
}
