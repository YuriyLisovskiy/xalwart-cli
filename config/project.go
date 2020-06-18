package config

import (
	"github.com/gobuffalo/packr/v2"
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

	// Example: v2/, v3/, ...
	FrameworkVersionSubDir string

	ProjectName string
	ProjectRoot string

	SecretKey string

	CMakeCPPStandard int
	CMakeMinimumVersion string

	Templates* packr.Box
}
