package config

import (
	"github.com/iancoleman/strcase"
	"strings"
	"text/template"
)

const (
	AppVersion = "0.0.1"

	FrameworkName       = "xalwart"
	FrameworkNamespace  = "xw"
	MinimumCmakeVersion = "3.13"
	MinimumCppStandard  = 17

	SecretKeyLength = 50

	CMakeListsTxtToDoLine = "# TODO: setup and link '" + FrameworkName + "' framework."
)

var ConanRequiredPackages = []string{
	"xalwart/0.0.1",
}

var DefaultFunctions = template.FuncMap {
	"upper": strings.ToUpper,
	"to_camel_case": strcase.ToCamel,
	"to_snake_case": strcase.ToSnake,
}
