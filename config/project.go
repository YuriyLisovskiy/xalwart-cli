package config

import "github.com/gobuffalo/packr"

type Project struct {
	Year int
	Username string
	WorkingDirectory string

	FrameworkName string
	FrameworkNamespace string

	ProjectName string

	SecretKey string

	CMakeCPPStandard int
	CMakeMinimumVersion string

	Templates packr.Box
}
