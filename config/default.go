package config

import (
	"github.com/iancoleman/strcase"
	"strings"
	"text/template"
)

const (
	AppVersion = "0.0.1"

	FrameworkName = "xalwart"
	FrameworkNamespace = "xw"
	MinimumCmakeVersion = "3.13"
	MinimumCppVersion = 17

	SecretKeyLength = 50

	baseUrl = "https://api.github.com/repos/YuriyLisovskiy/TestRepo/releases"
	tarArchive = "/" + FrameworkName + "-framework.tar.gz"
	DownloadReleaseUrl = "https://github.com/YuriyLisovskiy/TestRepo/releases/download/v<version>/" + tarArchive
	ReleaseByTagUrl = baseUrl + "/tags/v<version>"
	LatestReleaseUrl = baseUrl + "/latest"
)

var DefaultFunctions = template.FuncMap {
	"upper": strings.ToUpper,
	"to_camel_case": strcase.ToCamel,
}
