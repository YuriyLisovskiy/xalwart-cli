package generator

import (
	"github.com/gobuffalo/packr/v2"
)

type ProjectUnit struct {
	Year int
	Username string

	FrameworkName string
	FrameworkNamespace string

	ProjectName string
	ProjectRoot string

	Name string
	Root string

	// Example: v2/, v3/, ...
//	FrameworkVersionSubDir string

	Templates* packr.Box

	Customize func (pu *ProjectUnit)



	WorkingDirectory string

	FrameworkVersion string

	SecretKey string

	CMakeCPPStandard int
	CMakeMinimumVersion string
}

func (pu *ProjectUnit) Init() {
	if pu.Customize != nil {
		pu.Customize(pu)
	}
}
