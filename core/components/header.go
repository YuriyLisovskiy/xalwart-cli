package components

import (
	"os/user"
	"strings"
	"text/template"
	"time"

	"github.com/YuriyLisovskiy/xalwart-cli/config"
)

type Header struct {
	year            int
	userName        string
	copyrightNotice string
}

func (h Header) Year() int {
	return h.year
}

func (h Header) UserName() string {
	return h.userName
}

func (h Header) FrameworkName() string {
	return config.FrameworkName
}

func (h Header) FrameworkNamespace() string {
	return config.FrameworkNamespace
}

func (h Header) CopyrightNotice() string {
	return h.copyrightNotice
}

func newHeader(copyrightNotice string) (*Header, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	header := &Header{
		year:     time.Now().Year(),
		userName: currentUser.Username,
	}

	tmpl, err := template.New("copyright notice").
		Funcs(config.DefaultFunctions).
		Delims("<%", "%>").
		Parse(copyrightNotice)
	if err != nil {
		return nil, err
	}

	noticeStream := new(strings.Builder)
	err = tmpl.Execute(noticeStream, *header)
	if err != nil {
		return nil, err
	}

	header.copyrightNotice = noticeStream.String()
	return header, nil
}
