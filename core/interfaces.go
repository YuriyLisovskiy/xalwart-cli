package core

type Component interface {
	Name() string
	FileName() string
	RootPath() string
	GetTargetPath(string) string
	TemplateBox() TemplateBox
}
