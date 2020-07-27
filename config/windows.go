// +build windows

package config

import "os"

const (
	GlobalInstallationRoot = "C:\\" + FrameworkName
)

var (
	TempDirectory = os.Getenv("LOCALAPPDATA")
)
