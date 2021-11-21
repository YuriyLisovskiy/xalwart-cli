package components

import (
	"errors"
	"path"
	"path/filepath"
	"strings"

	"github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
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
	templatePath = strings.TrimSuffix(templatePath, filepath.Ext(templatePath))
	return path.Join(s.RootPath(), templatePath)
}

func (s ProjectComponent) TemplateBox() core.TemplateBox {
	return s.common.templateBox
}

func (s ProjectComponent) Header() core.Header {
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
	header core.Header,
	templateBox core.TemplateBox,
	secretKey string,
	projectName string,
	rootPath string,
	useStandardORM bool,
	useStandardServer bool,
) *ProjectComponent {
	return &ProjectComponent{
		common: CommonComponent{
			header:      header,
			name:        projectName,
			rootPath:    rootPath,
			templateBox: templateBox,
		},
		secretKey:         secretKey,
		useStandardServer: useStandardServer,
		useStandardORM:    useStandardORM,
	}
}
