package generator

import (
	"crypto/rand"
	"fmt"
	"github.com/YuriyLisovskiy/xalwart-cli/config"
	"github.com/gobuffalo/packr/v2"
	"github.com/iancoleman/strcase"
	"math/big"
	"os/user"
	"path"
	"time"
)

type Unit interface {
	GetRootPath() string
	GetTemplates() *packr.Box
}

type SingleUnit interface {
	GetFileName() string
	GetUnitName() string
}

type CommonUnit struct {
	Year               int
	Username           string
	FrameworkName      string
	FrameworkNamespace string
	UnitName           string
	RootPath           string
	Templates          *packr.Box
}

func NewCommonUnit(boxName, unitName, rootPath string) (*CommonUnit, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	return &CommonUnit{
		Year:               time.Now().Year(),
		Username:           usr.Username,
		FrameworkName:      config.FrameworkName,
		FrameworkNamespace: config.FrameworkNamespace,
		UnitName:           unitName,
		RootPath:           rootPath,
		Templates: packr.New(
			fmt.Sprintf("%s box", boxName),
			fmt.Sprintf("../templates/%s", boxName),
		),
	}, nil
}

func (c *CommonUnit) GetRootPath() string {
	return c.RootPath
}

func (c *CommonUnit) GetTemplates() *packr.Box {
	return c.Templates
}

func (c *CommonUnit) GetFileName() string {
	return c.UnitName
}

func (c *CommonUnit) GetUnitName() string {
	return c.UnitName
}

type SingleFileUnit struct {
	Common         *CommonUnit
	UnitType       string
	CustomFileName string
	FullUnitName   string
}

func NewSingleFileUnit(
	boxName, unitName, unitType, rootPath, customFileName string,
) (*SingleFileUnit, error) {
	unit := &SingleFileUnit{
		UnitType:       unitType,
		CustomFileName: customFileName,
	}
	var err error = nil
	unit.Common, err = NewCommonUnit(boxName, unitName, rootPath)
	if err != nil {
		return nil, err
	}

	unit.FullUnitName = strcase.ToCamel(unit.Common.UnitName) + strcase.ToCamel(unit.UnitType)
	return unit, nil
}

func (c *SingleFileUnit) GetRootPath() string {
	return c.Common.RootPath
}

func (c *SingleFileUnit) GetTemplates() *packr.Box {
	return c.Common.Templates
}

func (c *SingleFileUnit) GetFileName() string {
	result := c.Common.UnitName
	if len(c.CustomFileName) != 0 {
		result = c.CustomFileName
	}

	return strcase.ToSnake(result)
}

func (c *SingleFileUnit) GetUnitName() string {
	return strcase.ToSnake(c.Common.UnitName)
}

func NewCommandUnit(unitName, rootPath, customFileName string) (*SingleFileUnit, error) {
	return NewSingleFileUnit(
		"command", unitName, "command", rootPath, customFileName,
	)
}

func NewControllerUnit(unitName, rootPath, customFileName string) (*SingleFileUnit, error) {
	return NewSingleFileUnit(
		"controller", unitName, "controller", rootPath, customFileName,
	)
}

func NewModuleUnit(unitName, rootPath, customFileName string) (*SingleFileUnit, error) {
	return NewSingleFileUnit(
		"module", unitName, "module", rootPath, customFileName,
	)
}

type ProjectUnit struct {
	Common      *CommonUnit
	SecretKey   string
	ProjectName string

	UseStandardServer bool
	UseStandardORM    bool
}

func NewProjectUnit(
	projectName, rootPath string, secretKeyLength uint,
	useStandardORM bool,
	useStandardServer bool,
) (*ProjectUnit, error) {
	secretKey, err := generateRandomString(secretKeyLength)
	if err != nil {
		return nil, err
	}

	unit := &ProjectUnit{
		SecretKey:         secretKey,
		ProjectName:       projectName,
		UseStandardServer: useStandardServer,
		UseStandardORM:    useStandardORM,
	}
	unit.Common, err = NewCommonUnit("project", projectName, rootPath)
	if err != nil {
		return nil, err
	}

	return unit, nil
}

func (p *ProjectUnit) GetRootPath() string {
	return path.Join(p.Common.RootPath, p.ProjectName)
}

func (p *ProjectUnit) GetTemplates() *packr.Box {
	return p.Common.Templates
}

func generateRandomString(length uint) (string, error) {
	n := int(length)
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*(-_=+)"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}

		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
