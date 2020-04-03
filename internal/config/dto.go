package config

type ProjectFile struct {
	Name string
	Path string
	TemplateStr string
}

type ProjectConfig struct {
	WorkingDirectory string

	ProjectName string

	FrameworkName string
	FrameworkNamespace string

	SecretKey string

	CMakeCPPStandard int
	CMakeMinimumVersion string
}
