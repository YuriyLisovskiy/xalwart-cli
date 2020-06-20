package config

const (
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