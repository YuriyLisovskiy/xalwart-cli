package utils

import (
	"io/ioutil"
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

func FileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func DirExists(dirName string) bool {
	info, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func DirIsEmpty(dirName string) bool {
	if !DirExists(dirName) {
		return true
	}

	entries, err := ioutil.ReadDir(dirName)
	if err != nil {
		return false
	}

	return len(entries) == 0
}
