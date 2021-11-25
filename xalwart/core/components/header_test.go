package components

import (
	"fmt"
	"testing"

	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
)

func TestHeaderComponent_CLikeCopyrightNoticeTemplates(t *testing.T) {
	header, _ := NewHeaderComponent(&templateBoxMock{Templates: map[string]string{
		"c-like.txt": "// Copyright (c) <% .Year %> <% .UserName %>",
		"number-sign.txt": "",
	}})
	cLikeExpected := fmt.Sprintf("// Copyright (c) %d %s", header.Year(), header.UserName())
	cLikeActual := header.CLikeCopyrightNotice()
	if cLikeActual != cLikeExpected {
		t.Errorf("Expected %s, received %s", cLikeExpected, cLikeActual)
	}
}

func TestHeaderComponent_NumberSignCopyrightNotice(t *testing.T) {
	header, _ := NewHeaderComponent(&templateBoxMock{Templates: map[string]string{
		"c-like.txt": "",
		"number-sign.txt": "# Copyright (c) <% .Year %> <% .UserName %>",
	}})
	numberSignExpected := fmt.Sprintf("# Copyright (c) %d %s", header.Year(), header.UserName())
	numberSignActual := header.NumberSignCopyrightNotice()
	if numberSignActual != numberSignExpected {
		t.Errorf("Expected %s, received %s", numberSignExpected, numberSignActual)
	}
}

func TestHeaderComponent_CLikeTemplateNotFound(t *testing.T) {
	_, err := NewHeaderComponent(&templateBoxMock{Templates: map[string]string{
		"number-sign.txt": "",
	}})
	if err == nil {
		t.Error("Expected not found error, received nil")
	}
}

func TestHeaderComponent_NumberSignTemplateNotFound(t *testing.T) {
	_, err := NewHeaderComponent(&templateBoxMock{Templates: map[string]string{
		"c-like.txt": "",
	}})
	if err == nil {
		t.Error("Expected not found error, received nil")
	}
}

func TestHeaderComponent_FrameworkName(t *testing.T) {
	header := HeaderComponent{}
	actual := header.FrameworkName()
	expected := core.FrameworkName
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestHeaderComponent_FrameworkNamespace(t *testing.T) {
	header := HeaderComponent{}
	actual := header.FrameworkNamespace()
	expected := core.FrameworkNamespace
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}

func TestHeaderComponent_FrameworkLatestVersion(t *testing.T) {
	header := HeaderComponent{}
	actual := header.FrameworkLatestVersion()
	expected := core.FrameworkLatestVersion
	if actual != expected {
		t.Errorf("Expected %s, received %s", expected, actual)
	}
}
