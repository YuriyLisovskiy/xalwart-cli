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
	"regexp"
	"strings"
	"text/template"
	"xalwart-cli/config"
	"xalwart-cli/managers"
	"xalwart-cli/utils"
)

const (
	installCmdName = "install"
	InstallCmdDescription = installCmdName + ":\tinstalls '" + config.FrameworkName + "' framework"
)

var (
	InstallCmd = flag.NewFlagSet(installCmdName, flag.ExitOnError)
	iVerbose bool
	iLocal bool
	iGlobal bool
	iCustom bool
	iRoot string
	iProjectRoot string
	iVersion string
	iUpdateProject bool
	iUpgrade bool
)

func InitInstallCmd() {
	description := "Print steps during the installation"
	InstallCmd.BoolVar(&iVerbose, "verbose", true, description)

	description = "Local installation, using '-p' value"
//	InstallCmd.BoolVar(&iLocal, "local", true, description)
	InstallCmd.BoolVar(&iLocal, "l", true, description)

	description = "Global installation at '" + config.GlobalInstallationRoot + "'"
//	InstallCmd.BoolVar(&iGlobal, "global", false, description)
	InstallCmd.BoolVar(&iGlobal, "g", false, description)

	description = "Custom installation, '-i' argument is required"
//	InstallCmd.BoolVar(&iCustom, "custom", false, description)
	InstallCmd.BoolVar(&iCustom, "c", false, description)

	description = "Path to '" + config.FrameworkName + "' framework installation directory"
//	InstallCmd.StringVar(&iRoot, "root", "", description)
	InstallCmd.StringVar(&iRoot, "i", "", description)

	description = "Root directory of a project (default is current working directory)"
//	InstallCmd.StringVar(&iProjectRoot, "project-root", "", description)
	InstallCmd.StringVar(&iProjectRoot, "p", "", description)

	description = "Version of '" + config.FrameworkName + "' framework to install"
//	InstallCmd.StringVar(&iVersion, "version", "latest", description)
	InstallCmd.StringVar(&iVersion, "v", "latest", description)

	description = "Append framework setup to 'CMakeLists.txt' and update '.project.xw' if this files exist"
//	InstallCmd.BoolVar(&iUpdateProject, "update-project", false,	description)
	InstallCmd.BoolVar(&iUpdateProject, "u", true,	description)

	description = "Upgrades an existing version of '" + config.FrameworkName + "' framework"
	InstallCmd.BoolVar(&iUpgrade, "upgrade", false,	description)
}

type templateModel struct {
	FrameworkName string
	FrameworkVersion string
	InstallationRoot string
}

func versionIsValid(version string) bool {
	matched, err := regexp.MatchString("(\\d+)\\.(\\d+)\\.(\\d+)(-(alpha|beta))?", version)
	return matched && err == nil
}

func (c *Cmd) InstallFramework() error {
	b2i := map[bool]int8{true: 1, false: 0}
	if (b2i[iLocal] + b2i[iGlobal] + b2i[iCustom]) != 1 {
		return errors.New(
			"unknown installation type use exactly one of -l, -g or -c",
		)
	}

	if iGlobal {
		iRoot = config.GlobalInstallationRoot
	}

	if iCustom && len(iRoot) == 0 {
		return errors.New("-i argument is required when custom installation is chosen")
	}

	if iUpdateProject && len(iProjectRoot) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		iProjectRoot = cwd
	}

	if iLocal {
		iRoot = iProjectRoot
	}

	if len(iVersion) == 0 {
		iVersion = "latest"
	}

	version, err := c.getVersionOfFramework(iVersion, iVerbose)
	if err != nil {
		return err
	}

	var meta projectMeta
	metaIsInitialized := false
	exists := managers.FrameworkExists(iRoot)
	if iUpgrade {
		if exists {
			meta, err = c.loadMeta(iProjectRoot)
			if err == nil {
				metaIsInitialized = true
				if versionIsValid(meta.FrameworkVersion) {
					if version == meta.FrameworkVersion {
						if iVerbose {
							fmt.Println(config.FrameworkName + " is already the newest version (" + version + ")")
						}

						return nil
					} else if version < meta.FrameworkVersion {
						q := utils.NewIO()
						confirmed, err := q.ReadBool(
							"Current version of framework is newer than the one you want to install. " +
							"Are you sure you want to downgrade to earlier version? [Y/n] ",
						)
						if err != nil {
							return err
						}

						if !confirmed {
							return nil
						}
					}
				}
			}

			err = os.RemoveAll(path.Join(iRoot, "include", config.FrameworkName))
			if err != nil {
				return err
			}

			err = os.RemoveAll(path.Join(iRoot, "lib", config.FrameworkName))
			if err != nil {
				return err
			}
		} else {
			return errors.New(
				"unable to upgrade framework in '" + iRoot + "' because it is not installed",
			)
		}
	} else {
		if exists {
			return errors.New(
				"unable to install framework in '" + iRoot + "' because it already exists, use '--upgrade'",
			)
		}
	}

	err = managers.InstallFramework(iRoot, version, iVerbose)
	if err != nil {
		return err
	}

	if iUpdateProject {
		var msgEnding string
		if !metaIsInitialized {
			meta, err = c.loadMeta(iProjectRoot)
			if err != nil {
				if iVerbose {
					fmt.Println("Warning: " + err.Error())
				}
			} else {
				metaIsInitialized = true
			}
		}

		if metaIsInitialized {
			if len(meta.ProjectName) != 0 {
				msgEnding = " of '" + meta.ProjectName + "' project"
			}

			meta.FrameworkVersion = version
			err = c.saveMeta(iProjectRoot, meta)
			if err != nil {
				if iVerbose {
					fmt.Println("Warning: " + err.Error())
				}
			}
		}

		if !iUpgrade {
			cMakeListsPath := path.Join(iProjectRoot, "CMakeLists.txt")
			cMakeListsTxtBytes, err := ioutil.ReadFile(cMakeListsPath)
			if err != nil {
				if iVerbose {
					fmt.Println("Warning: unable to update 'CMakeLists.txt'" + msgEnding)
				}
			} else {
				model := templateModel{
					FrameworkName:    config.FrameworkName,
					FrameworkVersion: version,
				}
				if iLocal {
					model.InstallationRoot = "${PROJECT_SOURCE_DIR}"
				} else if iGlobal {
					model.InstallationRoot = config.GlobalInstallationRoot
				} else if iCustom {
					model.InstallationRoot = iRoot
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
					if iVerbose {
						fmt.Println("Warning: unable to update 'CMakeLists.txt'" + msgEnding)
					}
				}
			}
		}
	}

	return nil
}
