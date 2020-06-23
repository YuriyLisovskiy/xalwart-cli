// +build windows

package config

const (
	GlobalInstallationRoot = "C:\\" + FrameworkName
)

var (
	TempDirectory = os.Getenv("LOCALAPPDATA")
)
