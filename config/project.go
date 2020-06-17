package config

import (
	"github.com/gobuffalo/packr"
	"strings"
	"text/template"
)

var DefaultFunctions = template.FuncMap {
	"upper": strings.ToUpper,
}

type Project struct {
	Year int
	Username string
	WorkingDirectory string

	FrameworkName string
	FrameworkNamespace string
	FrameworkVersion string

	ProjectName string
	ProjectRoot string

	SecretKey string

	CMakeCPPStandard int
	CMakeMinimumVersion string

	Templates packr.Box
}
