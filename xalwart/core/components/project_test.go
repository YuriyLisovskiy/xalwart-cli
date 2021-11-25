package components

import (
	"testing"

	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
)

var projectTestProjectComponent = ProjectComponent{
	common: CommonComponent{
		name:     "MyTestProject",
		rootPath: "/tmp",
	},
	secretKey:         "super-secret-key",
	useStandardServer: true,
	useStandardORM:    true,
}

func TestProjectComponent_FileName(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic")
		}
	}()
	_ = projectTestProjectComponent.FileName()
}

func TestProjectComponent_RootPath(t *testing.T) {
	actual := projectTestProjectComponent.RootPath()
	expected := "/tmp/MyTestProject"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestProjectComponent_GetTargetPath(t *testing.T) {
	actual := projectTestProjectComponent.GetTargetPath("./some/dir/template.h.txt")
	expected := "/tmp/MyTestProject/some/dir/template.h"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestProjectComponent_SecretKey(t *testing.T) {
	actual := projectTestProjectComponent.SecretKey()
	expected := projectTestProjectComponent.secretKey
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestProjectComponent_ProjectName(t *testing.T) {
	actual := projectTestProjectComponent.ProjectName()
	expected := "MyTestProject"
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestProjectComponent_UseStandardServer(t *testing.T) {
	actual := projectTestProjectComponent.UseStandardServer()
	if !actual {
		t.Errorf("Expected true, received %v", actual)
	}
}

func TestProjectComponent_UseStandardORM(t *testing.T) {
	actual := projectTestProjectComponent.UseStandardORM()
	if !actual {
		t.Errorf("Expected true, received %v", actual)
	}
}

func TestProjectComponent_FrameworkBaseDockerImage(t *testing.T) {
	actual := projectTestProjectComponent.FrameworkBaseDockerImage()
	expected := core.FrameworkBaseDockerImage
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}
