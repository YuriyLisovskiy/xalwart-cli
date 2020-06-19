package utils

import (
	"os"
	"path"
)

func MakeDirs(rootDir string, dirsToCreate []string) error {
	for _, dir := range dirsToCreate {
		var fullDir string
		if path.IsAbs(dir) {
			fullDir = dir
		} else {
			fullDir = path.Join(rootDir, dir)
		}

		err := os.MkdirAll(fullDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
