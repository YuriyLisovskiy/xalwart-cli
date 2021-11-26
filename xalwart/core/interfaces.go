package core

type Template interface {
	Render(Component) error
	Path() string
	String() string
	SetTargetPath(string)
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

type Header interface {
	Year() int
	UserName() string
	FrameworkName() string
	FrameworkNamespace() string
	CLikeCopyrightNotice() string
	NumberSignCopyrightNotice() string
}
