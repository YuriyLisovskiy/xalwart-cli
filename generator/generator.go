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

func executeTemplate(templateFilePath string, file packd.File, unit Unit) error {
	filePath, fileName := path.Split(templateFilePath)
	filePath = path.Join(unit.GetRootPath(), filePath)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	if singleUnit, ok := unit.(SingleUnit); ok {
		fileName = strings.ReplaceAll(fileName, "_name_", singleUnit.GetFileName())
		filePath = strings.ReplaceAll(filePath, "_name_", singleUnit.GetUnitName())
	}

	fullPath := path.Join(filePath, fileName)
	//if fileExists(fullPath) {
	//	fmt.Println(fmt.Sprintf("Warning: ignoring '%s', file already exists", fullPath))
	//	return nil
	//}

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
		return executeTemplate(filePath, file, unit)
	})
	if err != nil {
		return err
	}

	return nil
}
