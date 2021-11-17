package core

type Template interface {
	Execute(Component) error
}

type TemplateBox interface {
	Walk(func(Template) error, Component, bool) error
	FindString(string) (string, error)
}

type Component interface {
	Name() string
	FileName() string
	RootPath() string
	GetTargetPath(string) string
	TemplateBox() TemplateBox
}
