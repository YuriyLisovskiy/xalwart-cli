package components

import (
	"errors"
	"os/user"
	"strings"
	"text/template"
	"time"

	core2 "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
)

type HeaderComponent struct {
	year            int
	userName        string
	cLikeCopyrightNotice string
	numberSignCopyrightNotice string
}

func (h HeaderComponent) Year() int {
	return h.year
}

func (h HeaderComponent) UserName() string {
	return h.userName
}

func (h HeaderComponent) FrameworkName() string {
	return core2.FrameworkName
}

func (h HeaderComponent) FrameworkNamespace() string {
	return core2.FrameworkNamespace
}

func (h HeaderComponent) CLikeCopyrightNotice() string {
	return h.cLikeCopyrightNotice
}

func (h HeaderComponent) NumberSignCopyrightNotice() string {
	return h.numberSignCopyrightNotice
}

func (h HeaderComponent) renderTemplate(text string) (string, error) {
	tmpl, err := template.New("text template").
		Funcs(core2.DefaultFunctions).
		Delims("<%", "%>").
		Parse(text)
	if err != nil {
		return "", err
	}

	noticeStream := new(strings.Builder)
	err = tmpl.Execute(noticeStream, h)
	if err != nil {
		return "", err
	}

	return noticeStream.String(), nil
}

func NewHeaderComponent(copyrightTemplatesBox core2.TemplateBox) (*HeaderComponent, error) {
	if copyrightTemplatesBox == nil {
		return nil, errors.New("copyright templates box is nil")
	}

	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	header := &HeaderComponent{
		year:     time.Now().Year(),
		userName: currentUser.Username,
	}

	cLikeNotice, err := copyrightTemplatesBox.FindString("c-like.txt")
	if err != nil {
		return nil, err
	}

	header.cLikeCopyrightNotice, err = header.renderTemplate(cLikeNotice)
	if err != nil {
		return nil, err
	}

	numberSignNotice, err := copyrightTemplatesBox.FindString("number-sign.txt")
	if err != nil {
		return nil, err
	}

	header.numberSignCopyrightNotice, err = header.renderTemplate(numberSignNotice)
	if err != nil {
		return nil, err
	}

	return header, nil
}
