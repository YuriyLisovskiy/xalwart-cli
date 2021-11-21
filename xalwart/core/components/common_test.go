package components

import (
	"errors"
	"testing"

	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
)

type componentMock struct {
	NameValue        string
	FileNameValue    string
	RootPathValue    string
	TargetPathValue  string
	TemplateBoxValue core.TemplateBox
}

func (c componentMock) Name() string {
	return c.NameValue
}

func (c componentMock) FileName() string {
	return c.FileNameValue
}

func (c componentMock) RootPath() string {
	return c.RootPathValue
}

func (c componentMock) GetTargetPath(string) string {
	return c.TargetPathValue
}

func (c componentMock) TemplateBox() core.TemplateBox {
	return c.TemplateBoxValue
}

type templateBoxMock struct {
	Templates map[string]string
}

func (t templateBoxMock) Walk(func(core.Template) error, core.Component, bool) error {
	return errors.New("not implemented")
}

func (t templateBoxMock) FindString(name string) (string, error) {
	if s, ok := t.Templates[name]; ok {
		return s, nil
	}

	return "", errors.New("template does not exist")
}

var commonTestComponent = componentMock{
	NameValue:        "super",
	FileNameValue:    "super_file",
	RootPathValue:    "/tmp",
}

func TestGetTargetPath_WithGenericName(t *testing.T) {
	actual := getTargetPath("./some/name/to/_name_.yaml.txt", &commonTestComponent)
	expected := "/tmp/some/name/to/super_file.yaml"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestGetTargetPath_WithGenericPathPart(t *testing.T) {
	actual := getTargetPath("./some/_name_/to/file.yaml.txt", &commonTestComponent)
	expected := "/tmp/some/super/to/file.yaml"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestGetTargetPath_WithGenericNameAndPathPart(t *testing.T) {
	actual := getTargetPath("./some/_name_/to/_name_.yaml.txt", &commonTestComponent)
	expected := "/tmp/some/super/to/super_file.yaml"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestGetTargetPath_WithoutExt(t *testing.T) {
	actual := getTargetPath("./some/_name_/to/_name_", &commonTestComponent)
	expected := "/tmp/some/super/to/super_file"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}
