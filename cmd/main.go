package main

import (
	"github.com/gobuffalo/packr"
	"os"
	"os/user"
	"time"
	"xalwart-cli/config"
	"xalwart-cli/generator"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	cfg := config.Project{
		Year:                time.Now().Year(),
		Username:            usr.Username,
		WorkingDirectory:    cwd,
		FrameworkName:       "xalwart",
		FrameworkNamespace:  "xw",
		ProjectName:         "HelloWorld",
		SecretKey:           "+s6cv712&nw4gsk)1dmgpje+f#%^4lhp@!up+=p3ts+hxz(fr2",
		CMakeCPPStandard:    17,
		CMakeMinimumVersion: "3.13",
		Templates:           packr.NewBox("../templates/project"),
	}

	g := generator.Generator{}
	g.NewProject(&cfg)
}
