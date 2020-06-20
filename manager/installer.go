package manager

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
)

func downloadGithubRelease(archiveFile string, version string) error {
	output, err := os.Create(archiveFile)
	if output != nil {
		defer output.Close()
	}

	client := &http.Client{Timeout: 10 * time.Second}
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

//	fmt.Println(response.StatusCode)

//	bodyAsBytes := make([]byte, response.ContentLength)
//	_, err = response.Body.Read(bodyAsBytes)
//	if err != nil {
//		return "", err
//	}

//	var data map[string]interface{}
//	err = json.Unmarshal(bodyAsBytes, &data)

	target := Releases{}
	err = json.NewDecoder(response.Body).Decode(&target)
	if err != nil {
		return "", err
	}

	return strings.TrimLeft(target.VersionTag, "v"), nil
}

func InstallFramework(targetDir string, version string) error {
	archiveFile := "/tmp/" + config.FrameworkName + "-framework.tar.gz"

	fmt.Print("Downloading '" + config.FrameworkName + "' framework...")
	err := downloadGithubRelease(archiveFile, version)
	if err != nil {
		return err
	}

	fmt.Println("Done.")
	fmt.Print("Installing...")
	reader, err := os.Open(archiveFile)
	if err != nil {
		return err
	}

	err = extractTarGz(targetDir, reader)
	if err != nil {
		return err
	}

	fmt.Println("Done.")

	err = os.Remove(archiveFile)
	if err != nil {
		return err
	}

	return nil
}
