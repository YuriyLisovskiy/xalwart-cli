package components

import "testing"

var classTestHeader = HeaderComponent{
	year:                      2021,
	userName:                  "test_user",
	cLikeCopyrightNotice:      "// C-like copyright",
	numberSignCopyrightNotice: "# Number sign copyright",
}

var classTestTemplateBox = templateBoxMock{Templates: map[string]string{}}

var classTestClassComponent = ClassComponent{
	common: CommonComponent{
		header:      classTestHeader,
		name:        "test_name",
		rootPath:    "/tmp",
		templateBox: classTestTemplateBox,
	},
	componentType:  "super_component",
	customFileName: "",
}

func TestClassComponent_Header(t *testing.T) {
	actual := classTestClassComponent.Header()
	expected := classTestHeader
	if actual != expected {
		t.Errorf("Expected %v, received %v", expected, actual)
	}
}

func TestClassComponent_ClassName(t *testing.T) {
	actual := classTestClassComponent.ClassName()
	expected := "TestNameSuperComponent"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestClassComponent_Name(t *testing.T) {
	actual := classTestClassComponent.Name()
	expected := "test_name"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestClassComponent_FileName_Default(t *testing.T) {
	actual := classTestClassComponent.FileName()
	expected := "test_name_super_component"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestClassComponent_FileName_Custom(t *testing.T) {
	classTestClassComponent.customFileName = "custom_file_name"
	actual := classTestClassComponent.FileName()
	expected := classTestClassComponent.customFileName
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}

	classTestClassComponent.customFileName = ""
}

func TestClassComponent_RootPath(t *testing.T) {
	actual := classTestClassComponent.RootPath()
	expected := classTestClassComponent.common.rootPath
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestClassComponent_GetTargetPath(t *testing.T) {
	actual := classTestClassComponent.GetTargetPath("/tmp/_name_.txt.txt")
	expected := "/tmp/test_name_super_component.txt"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}
