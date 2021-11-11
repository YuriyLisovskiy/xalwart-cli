package generator

import (
	"errors"
	"fmt"
	"github.com/YuriyLisovskiy/xalwart-cli/config"
	"github.com/gobuffalo/packr/v2"
	"os/user"
	"time"
)

type Unit interface {
	GetRootPath() string
	GetTemplates() *packr.Box
	GetFileExistsError(string) error
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
			fmt.Sprintf("../../templates/%s", boxName),
		),
	}, nil
}

func (c *CommonUnit) GetRootPath() string {
	return c.RootPath
}

func (c *CommonUnit) GetTemplates() *packr.Box {
	return c.Templates
}

type CommandUnit struct {
	Common *CommonUnit
}

func NewCommandUnit(unitName, rootPath string) (*CommandUnit, error) {
	unit := &CommandUnit{}
	var err error = nil
	unit.Common, err = NewCommonUnit("command", unitName, rootPath)
	if err != nil {
		return nil, err
	}

	return unit, err
}

func (c *CommandUnit) GetRootPath() string {
	return c.Common.RootPath
}

func (c *CommandUnit) GetTemplates() *packr.Box {
	return c.Common.Templates
}

func (c *CommandUnit) GetFileExistsError(fileName string) error {
	return errors.New(
		fmt.Sprintf("command with target filename '%s' already exists", fileName),
	)
}

type ControllerUnit struct {
	Common *CommonUnit
}

func NewControllerUnit(unitName, rootPath string) (*ControllerUnit, error) {
	unit := &ControllerUnit{}
	var err error = nil
	unit.Common, err = NewCommonUnit("controller", unitName, rootPath)
	if err != nil {
		return nil, err
	}

	return unit, err
}

func (c *ControllerUnit) GetRootPath() string {
	return c.Common.RootPath
}

func (c *ControllerUnit) GetTemplates() *packr.Box {
	return c.Common.Templates
}

func (c *ControllerUnit) GetFileExistsError(fileName string) error {
	return errors.New(
		fmt.Sprintf("controller with target filename '%s' already exists", fileName),
	)
}

type ModuleUnit struct {
	Common *CommonUnit
}

func NewModuleUnit(unitName, rootPath string) (*ModuleUnit, error) {
	unit := &ModuleUnit{}
	var err error = nil
	unit.Common, err = NewCommonUnit("module", unitName, rootPath)
	if err != nil {
		return nil, err
	}

	return unit, err
}

func (m *ModuleUnit) GetRootPath() string {
	return m.Common.RootPath
}

func (m *ModuleUnit) GetTemplates() *packr.Box {
	return m.Common.Templates
}

func (m *ModuleUnit) GetFileExistsError(string) error {
	return errors.New("module in target directory already exists")
}
