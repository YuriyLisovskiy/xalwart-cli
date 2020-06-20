package config

import "github.com/gobuffalo/packr/v2"

type App struct {
	Year int
	Username string

	FrameworkName string
	FrameworkNamespace string

	AppName string

	// Example: v2/, v3/, ...
	FrameworkVersionSubDir string

	ProjectRoot string
	AppRoot string

	Templates* packr.Box
}
