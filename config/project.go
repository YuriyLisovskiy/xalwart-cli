package config

import (
	"github.com/gobuffalo/packr/v2"
	"path"
)

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

func (p *Project) MakeRoot() {
	if len(p.ProjectName) == 0 {
		p.ProjectRoot = p.WorkingDirectory
		p.ProjectName = path.Base(p.ProjectRoot)
	} else {
		p.ProjectRoot = path.Join(p.WorkingDirectory, p.ProjectName)
	}
}
