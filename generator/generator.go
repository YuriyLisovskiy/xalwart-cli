package generator

import (
	"github.com/YuriyLisovskiy/xalwart-cli/config"
	"github.com/gobuffalo/packd"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

func fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}


func executeTemplate(
	templateFilePath string, file packd.File, rootPath string,
	unit interface{}, fileExistsError func(string) error,
) error {
	filePath, fileName := path.Split(templateFilePath)
	filePath = path.Join(rootPath, filePath)
	fullPath := path.Join(filePath, strings.TrimSuffix(fileName, filepath.Ext(fileName)))
	if fileExists(fullPath) && fileExistsError != nil {
		return fileExistsError(fullPath)
	}

	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return err
	}

	stream, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	tmpl, err := template.New(templateFilePath).
		Funcs(config.DefaultFunctions).
		Delims("<%", "%>").
		Parse(file.String())
	if err != nil {
		return err
	}

	err = tmpl.Execute(stream, unit)
	if err != nil {
		panic(err)
	}

	err = stream.Close()
	if err != nil {
		return err
	}

	return nil
}

func GenerateUnit(unit Unit) error {
	err := os.MkdirAll(unit.GetRootPath(), os.ModePerm)
	if err != nil {
		return err
	}

	err = unit.GetTemplates().Walk(func(filePath string, file packd.File) error {
		return executeTemplate(filePath, file, unit.GetRootPath(), unit, unit.GetFileExistsError)
	})
	if err != nil {
		return err
	}

	return nil
}
