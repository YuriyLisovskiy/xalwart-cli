package main

import (
	"os"
	generators "wasp-cli/internal"
	"wasp-cli/internal/config"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cfg := config.ProjectConfig{
		WorkingDirectory:    cwd,
		ProjectName:         "HelloWorld",
		FrameworkName:       "wasp",
		FrameworkNamespace:  "wasp",
		SecretKey:           "+s6cv712&nw4gsk)1dmgpje+f#%^4lhp@!up+=p3ts+hxz(fr2",
		CMakeCPPStandard:    17,
		CMakeMinimumVersion: "3.13",
	}

	generators.GenerateProject(&cfg)
}
