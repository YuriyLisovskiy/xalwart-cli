package core

import (
	"errors"
	"fmt"
	"os"
	"path"
	"text/template"
)

type FileTemplate struct {
	templatePath string
	targetPath   string
	content      string
}

func NewFileTemplateWithContent(templatePath string, content string) *FileTemplate {
	return &FileTemplate{
		templatePath: templatePath,
		content:      content,
	}
}

func (t *FileTemplate) Render(component Component) error {
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
		Parse(t.String())
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

func (t FileTemplate) Path() string {
	return t.templatePath
}

func (t FileTemplate) String() string {
	return t.content
}

func (t *FileTemplate) SetTargetPath(targetPath string) {
	t.targetPath = targetPath
}

type FileTemplateBox struct {
	templates map[string]Template
}

func (b *FileTemplateBox) Walk(function func(Template) error, component Component, overwrite bool) error {
	for _, item := range b.templates {
		targetPath := component.GetTargetPath(item.Path())
		if b.fileExists(targetPath) && !overwrite {
			fmt.Println(fmt.Sprintf("Warning: ignoring '%s', file already exists", targetPath))
			return nil
		}

		item.SetTargetPath(targetPath)
		if err := function(item); err != nil {
			return err
		}
	}

	return nil
}

func (b *FileTemplateBox) FindString(name string) (string, error) {
	item, exists := b.templates[name]
	if !exists {
		return "", errors.New("template does not exist")
	}

	return item.String(), nil
}

func (b FileTemplateBox) fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func NewFileTemplateBoxWithTemplates(templates []Template) *FileTemplateBox {
	templatesMap := map[string]Template{}
	for _, item := range templates {
		templatesMap[item.Path()] = item
	}

	return &FileTemplateBox{templates: templatesMap}
}
