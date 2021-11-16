package components

import (
	"errors"
	"path"
	"path/filepath"
	"strings"

	"github.com/YuriyLisovskiy/xalwart-cli/core"
)

type ProjectComponent struct {
	common    CommonComponent
	secretKey string

	useStandardServer bool
	useStandardORM    bool
}

func (s ProjectComponent) Name() string {
	return s.common.name
}

func (s ProjectComponent) FileName() string {
	panic(errors.New("ProjectComponent is not single-file component"))
}

func (s ProjectComponent) RootPath() string {
	return path.Join(s.common.rootPath, s.ProjectName())
}

func (s ProjectComponent) GetTargetPath(templatePath string) string {
	filePath, fileName := path.Split(templatePath)
	filePath = path.Join(s.RootPath(), filePath)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return path.Join(filePath, fileName)
}

func (s ProjectComponent) TemplateBox() core.TemplateBox {
	return s.common.templateBox
}

func (s ProjectComponent) Header() Header {
	return s.common.header
}

func (s ProjectComponent) SecretKey() string {
	return s.secretKey
}

func (s ProjectComponent) ProjectName() string {
	return s.Name()
}

func (s ProjectComponent) UseStandardServer() bool {
	return s.useStandardServer
}

func (s ProjectComponent) UseStandardORM() bool {
	return s.useStandardORM
}

func NewProjectComponent(
	projectName string,
	rootPath string,
	secretKeyLength uint,
	useStandardORM bool,
	useStandardServer bool,
) (*ProjectComponent, error) {
	secretKey, err := core.RandomString(secretKeyLength)
	if err != nil {
		return nil, err
	}

	commonComponent, err := newCommonComponent("project", projectName, rootPath)
	if err != nil {
		return nil, err
	}

	return &ProjectComponent{
		common:            *commonComponent,
		secretKey:         secretKey,
		useStandardServer: useStandardServer,
		useStandardORM:    useStandardORM,
	}, nil
}
