package utils

import (
	core2 "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"
)

func GetCommandTemplateBox() core2.TemplateBox {
	return core2.NewFileTemplateBox("command")
}

func GetControllerTemplateBox() core2.TemplateBox {
	return core2.NewFileTemplateBox("controller")
}

func GetCopyrightNoticesTemplateBox() core2.TemplateBox {
	return core2.NewFileTemplateBox("copyright_notices")
}

func GetMiddlewareTemplateBox() core2.TemplateBox {
	return core2.NewFileTemplateBox("middleware")
}

func GetMigrationTemplateBox() core2.TemplateBox {
	return core2.NewFileTemplateBox("migration")
}

func GetModelTemplateBox() core2.TemplateBox {
	return core2.NewFileTemplateBox("model")
}

func GetModuleTemplateBox() core2.TemplateBox {
	return core2.NewFileTemplateBox("module")
}

func GetProjectTemplateBox() core2.TemplateBox {
	return core2.NewFileTemplateBox("project")
}
