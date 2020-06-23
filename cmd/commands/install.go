package commands

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
	"xalwart-cli/config"
	"xalwart-cli/managers"
)

const (
	installCmdName = "install"
	InstallCmdDescription = "Installs '" + config.FrameworkName + "' framework"
)

var (
	InstallCmd = flag.NewFlagSet(installCmdName, flag.ExitOnError)

	iVerbose = InstallCmd.Bool("verbose", true, "Print steps during the installation")
	iLocal = InstallCmd.Bool("local", true, "Local installation, using '--project-root' value")
	iGlobal = InstallCmd.Bool(
		"global", false, "Global installation at '" + config.GlobalInstallationRoot + "'",
	)
	iCustom = InstallCmd.Bool(
		"custom", false, "Custom installation, '--root' argument is required",
	)
	iRoot = InstallCmd.String(
		"root", "", "Path to '" + config.FrameworkName + "' framework installation directory",
	)
	iProjectRoot = InstallCmd.String(
		"project-root", "", "Root directory of a project (default is current working directory)",
	)
	iVersion = InstallCmd.String(
		"version", "latest", "Version of '" + config.FrameworkName + "' framework to install",
	)
	iUpdateProject = InstallCmd.Bool(
		"update-project",
		false,
		"Append framework setup to 'CMakeLists.txt' and update '.project.xw' if this files exist",
	)
)

type templateModel struct {
	FrameworkName string
	FrameworkVersion string
	InstallationRoot string
}

func (c *Cmd) InstallFramework() error {
	b2i := map[bool]int8{true: 1, false: 0}
	if (b2i[*iLocal] + b2i[*iGlobal] + b2i[*iCustom]) != 1 {
		return errors.New(
			"unknown installation type use exactly one of '--local', '--global' or '--custom'",
		)
	}

	root := *iRoot
	projectRoot := *iProjectRoot
	if *iGlobal {
		root = config.GlobalInstallationRoot
	}

	if *iCustom && len(root) == 0 {
		return errors.New("'--root' argument is required when custom installation is chosen")
	}

	if *iUpdateProject && len(projectRoot) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		projectRoot = cwd
	}

	if *iLocal {
		root = projectRoot
	}

	version := *iVersion
	if len(version) == 0 {
		version = "latest"
	}

	verbose := *iVerbose
	version, err := c.getVersionOfFramework(version, verbose)
	if err != nil {
		return err
	}

	err = managers.InstallFramework(root, version, verbose)
	if err != nil {
		return err
	}

	if *iUpdateProject {
		var msgEnding string
		meta, err := c.loadMeta(projectRoot)
		if err != nil {
			if verbose {
				fmt.Println("Warning: " + err.Error())
			}
		} else {
			if len(meta.ProjectName) != 0 {
				msgEnding = " of '" + meta.ProjectName + "' project"
			}

			meta.FrameworkVersion = version
			err = c.saveMeta(projectRoot, meta)
			if err != nil {
				if verbose {
					fmt.Println("Warning: " + err.Error())
				}
			}
		}

		cMakeListsPath := path.Join(projectRoot, "CMakeLists.txt")
		cMakeListsTxtBytes, err := ioutil.ReadFile(cMakeListsPath)
		if err != nil {
			if verbose {
				fmt.Println("Warning: unable to update 'CMakeLists.txt'" + msgEnding)
			}
		} else {
			model := templateModel{
				FrameworkName:    config.FrameworkName,
				FrameworkVersion: version,
			}
			if *iLocal {
				model.InstallationRoot = "${PROJECT_SOURCE_DIR}"
			} else if *iGlobal {
				model.InstallationRoot = config.GlobalInstallationRoot
			} else if *iCustom {
				model.InstallationRoot = root
			}

			tmplBox := packr.New("Installation Templates Box", "../../templates/install")
			partialStr, err := tmplBox.FindString("cmake-lists-partial.txt")
			if err != nil {
				return err
			}

			tmpl, err := template.New(
				"cmake-lists").Funcs(
				config.DefaultFunctions).Delims(
				"<%", "%>").Parse(partialStr)
			if err != nil {
				return err
			}

			var bytesResult bytes.Buffer
			if err := tmpl.Execute(&bytesResult, model); err != nil {
				return err
			}

			result := bytesResult.String()
			cMakeListsTxt := string(cMakeListsTxtBytes)
			if strings.Contains(cMakeListsTxt, config.CMakeListsTxtToDoLine) {
				cMakeListsTxt = strings.Replace(cMakeListsTxt, config.CMakeListsTxtToDoLine, result, 1)
			} else {
				cMakeListsTxt += result
			}

			err = ioutil.WriteFile(cMakeListsPath, []byte(cMakeListsTxt), 0644)
			if err != nil {
				if verbose {
					fmt.Println("Warning: unable to update 'CMakeLists.txt'" + msgEnding)
				}
			}
		}
	}

	return nil
}
