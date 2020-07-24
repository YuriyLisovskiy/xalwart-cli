package managers

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
	"xalwart-cli/src/config"
	"xalwart-cli/src/utils"
)

type Asset struct {
	Name string                 `json:"name"`
	BrowserDownloadUrl string   `json:"browser_download_url"`
}

type Release struct {
	VersionTag string   `json:"tag_name"`
	Assets []Asset      `json:"assets"`
}

func downloadGithubRelease(archiveFile string, url string) error {
	output, err := os.Create(archiveFile)
	if output != nil {
		defer output.Close()
	}

	client := &http.Client{Timeout: 20 * time.Minute}
	response, err := client.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func extractTarGz(targetDir string, gzipStream io.Reader) error {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		absPath := path.Join(targetDir, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(absPath, os.ModePerm); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(path.Dir(absPath), os.ModePerm); err != nil {
				return err
			}

			outFile, err := os.Create(absPath)
			if err != nil {
				return err
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}

			err = outFile.Close()
			if err != nil {
				return err
			}
		default:
			return errors.New("extractTarGz: unknown type in " + header.Name)
		}
	}

	return nil
}

func findSuitableAsset(assets []Asset) (Asset, error) {
	for _, asset := range assets {
		if strings.Contains(asset.Name, runtime.GOOS) {
			return asset, nil
		}
	}

	return Asset{}, errors.New(
		"'" + config.FrameworkName + "' framework is not supported under '" +
		runtime.GOOS + "' operating system",
	)
}

func CheckIfVersionIsAvailable(version string) (bool, error) {
	client := &http.Client{Timeout: 1 * time.Minute}
	response, err := client.Get(strings.Replace(config.ReleaseByTagUrl, "<version>", version, 1))
	if err != nil {
		return false, err
	}

	defer response.Body.Close()
	return response.StatusCode == 200, nil
}

func GetRelease(url string) (Release, error) {
	client := &http.Client{Timeout: 1 * time.Minute}
	response, err := client.Get(url)
	target := Release{}
	if err != nil {
		return target, err
	}

	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&target)
	if err != nil {
		return target, err
	}

	target.VersionTag = strings.TrimLeft(target.VersionTag, "v")
	return target, nil
}

func GetLatestRelease() (Release, error) {
	return GetRelease(config.LatestReleaseUrl)
}

func FrameworkExists(rootDir string) bool {
	includeExists := !utils.DirIsEmpty(path.Join(rootDir, "include", config.FrameworkName))
	libExists := !utils.DirIsEmpty(path.Join(rootDir, "lib", config.FrameworkName))
	return includeExists || libExists
}

func InstallFramework(rootDir string, version string, verbose bool) error {
	release, err := GetRelease(strings.Replace(config.ReleaseByTagUrl, "<version>", version, 1))
	if err != nil {
		return err
	}

	asset, err := findSuitableAsset(release.Assets)
	if err != nil {
		return err
	}

	archiveFile := path.Join(config.TempDirectory, asset.Name)
	if !utils.FileExists(archiveFile) {
		if verbose {
			fmt.Print("Downloading '" + config.FrameworkName + "' framework...")
		}

		err := downloadGithubRelease(archiveFile, asset.BrowserDownloadUrl)
		if err != nil {
			return err
		}

		if verbose {
			fmt.Println(" Done.")
		}
	} else {
		if verbose {
			fmt.Println("Using cached archive: '" + archiveFile + "'.")
		}
	}

	if verbose {
		fmt.Print("Installing...")
	}

	reader, err := os.Open(archiveFile)
	if err != nil {
		return err
	}

	err = extractTarGz(rootDir, reader)
	if err != nil {
		return err
	}

	if verbose {
		fmt.Println(" Done.")
	}

	return nil
}
