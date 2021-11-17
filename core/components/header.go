package components

import (
	"errors"
	"os/user"
	"strings"
	"text/template"
	"time"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
)

type Header struct {
	year            int
	userName        string
	cLikeCopyrightNotice string
	numberSignCopyrightNotice string
}

func (h Header) Year() int {
	return h.year
}

func (h Header) UserName() string {
	return h.userName
}

func (h Header) FrameworkName() string {
	return core.FrameworkName
}

func (h Header) FrameworkNamespace() string {
	return core.FrameworkNamespace
}

func (h Header) CLikeCopyrightNotice() string {
	return h.cLikeCopyrightNotice
}

func (h Header) NumberSignCopyrightNotice() string {
	return h.numberSignCopyrightNotice
}

func (h Header) renderTemplate(text string) (string, error) {
	tmpl, err := template.New("text template").
		Funcs(core.DefaultFunctions).
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

func newHeader(templateBox core.TemplateBox) (*Header, error) {
	if templateBox == nil {
		return nil, errors.New("template box is nil")
	}

	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	header := &Header{
		year:     time.Now().Year(),
		userName: currentUser.Username,
	}

	cLikeNotice, err := templateBox.FindString("c-like.txt")
	if err != nil {
		return nil, err
	}

	header.cLikeCopyrightNotice, err = header.renderTemplate(cLikeNotice)
	if err != nil {
		return nil, err
	}

	numberSignNotice, err := templateBox.FindString("number-sign.txt")
	if err != nil {
		return nil, err
	}

	header.numberSignCopyrightNotice, err = header.renderTemplate(numberSignNotice)
	if err != nil {
		return nil, err
	}

	return header, nil
}
