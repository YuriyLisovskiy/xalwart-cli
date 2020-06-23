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
	"strings"
	"time"
	"xalwart-cli/config"
	"xalwart-cli/utils"
)

func downloadGithubRelease(archiveFile string, version string) error {
	output, err := os.Create(archiveFile)
	if output != nil {
		defer output.Close()
	}

	client := &http.Client{Timeout: 20 * time.Minute}
	response, err := client.Get(
		strings.Replace(config.DownloadReleaseUrl, "<version>", version, 1),
	)
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

func CheckIfVersionIsAvailable(version string) (bool, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Get(strings.Replace(config.ReleaseByTagUrl, "<version>", version, 1))
	if err != nil {
		return false, err
	}

	defer response.Body.Close()
	return response.StatusCode == 200, nil
}

type Releases struct {
	VersionTag string `json:"tag_name"`
}

func GetLatestVersionOfFramework() (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Get(config.LatestReleaseUrl)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	target := Releases{}
	err = json.NewDecoder(response.Body).Decode(&target)
	if err != nil {
		return "", err
	}

	return strings.TrimLeft(target.VersionTag, "v"), nil
}

func InstallFramework(rootDir string, version string, verbose bool) error {
	archiveFile := path.Join(config.TempDirectory, config.FrameworkName + "-framework.tar.gz")
	if !utils.FileExists(archiveFile) {
		if verbose {
			fmt.Print("Downloading '" + config.FrameworkName + "' framework...")
		}

		err := downloadGithubRelease(archiveFile, version)
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
