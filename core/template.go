package core

import (
	"errors"
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
)

type FileTemplate struct {
	file         packd.File
	templatePath string
	targetPath   string
}

func NewFileTemplate(file packd.File, templatePath, targetPath string) (*FileTemplate, error) {
	return &FileTemplate{
		file:         file,
		templatePath: templatePath,
		targetPath:   targetPath,
	}, nil
}

func (t *FileTemplate) Execute(component Component) error {
	fileDir, _ := path.Split(t.targetPath)
	err := os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		return err
	}

	stream, err := os.Create(t.targetPath)
	if err != nil {
		return err
	}

	tmpl, err := template.New(t.templatePath).
		Funcs(DefaultFunctions).
		Delims("<%", "%>").
		Parse(t.file.String())
	if err != nil {
		return err
	}

	err = tmpl.Execute(stream, component)
	if err != nil {
		return err
	}

	err = stream.Close()
	if err != nil {
		return err
	}

	return nil
}

type FileTemplateBox struct {
	box *packr.Box
}

func (b *FileTemplateBox) Walk(function func(Template) error, component Component, overwrite bool) error {
	if b.box == nil {
		return errors.New("template box is nil")
	}

	return b.box.Walk(
		func(templatePath string, file packd.File) error {
			targetPath := component.GetTargetPath(templatePath)
			if b.fileExists(targetPath) && !overwrite {
				fmt.Println(fmt.Sprintf("Warning: ignoring '%s', file already exists", targetPath))
				return nil
			}

			fileTemplate, err := NewFileTemplate(file, templatePath, targetPath)
			if err != nil {
				return err
			}

			return function(fileTemplate)
		},
	)
}

func (b *FileTemplateBox) FindString(name string) (string, error) {
	if b.box == nil {
		return "", errors.New("template box is nil")
	}

	return b.box.FindString(name)
}

func (b FileTemplateBox) fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func NewFileTemplateBox(boxName string) *FileTemplateBox {
	return &FileTemplateBox{box: packr.Folder(fmt.Sprintf("../templates/%s", boxName))}
}
