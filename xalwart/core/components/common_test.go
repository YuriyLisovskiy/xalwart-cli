package components

import (
	"testing"

	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
)

type ComponentMock struct {
	NameValue        string
	FileNameValue    string
	RootPathValue    string
	TargetPathValue  string
	TemplateBoxValue core.TemplateBox
}

func (c ComponentMock) Name() string {
	return c.NameValue
}

func (c ComponentMock) FileName() string {
	return c.FileNameValue
}

func (c ComponentMock) RootPath() string {
	return c.RootPathValue
}

func (c ComponentMock) GetTargetPath(string) string {
	return c.TargetPathValue
}

func (c ComponentMock) TemplateBox() core.TemplateBox {
	return c.TemplateBoxValue
}

func TestGetTargetPath(t *testing.T) {
	component := ComponentMock{
		NameValue:        "super",
		FileNameValue:    "super_file",
		RootPathValue:    "/tmp",
	}
	actual := getTargetPath("./some/_name_/to/_name_.yaml.txt", &component)
	expected := "/tmp/some/super/to/super_file.yaml"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}
