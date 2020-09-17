package generator

import "github.com/gobuffalo/packr/v2"

type ProjectUnit struct {
	Year int
	Username string

	FrameworkName string
	FrameworkNamespace string
	UseConan bool
	CMakeListsTxtToDoLine string

	ProjectName string
	ProjectRoot string

	Name string
	Root string

	Templates* packr.Box
	TemplatesToExclude []string

	ConanRequiredPackages []string

	Customize func (pu *ProjectUnit)

	WorkingDirectory string

	SecretKey string

	CMakeCPPStandard int
	CMakeMinimumVersion string
}

func (pu *ProjectUnit) Init() {
	if pu.Customize != nil {
		pu.Customize(pu)
	}
}
