package utils

import "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"

func GetCommandTemplateBox() core.TemplateBox {
	return core.NewFileTemplateBox("command")
}

func GetControllerTemplateBox() core.TemplateBox {
	return core.NewFileTemplateBox("controller")
}

func GetCopyrightNoticesTemplateBox() core.TemplateBox {
	return core.NewFileTemplateBox("copyright_notices")
}

func GetMiddlewareTemplateBox() core.TemplateBox {
	return core.NewFileTemplateBox("middleware")
}

func GetMigrationTemplateBox() core.TemplateBox {
	return core.NewFileTemplateBox("migration")
}

func GetModelTemplateBox() core.TemplateBox {
	return core.NewFileTemplateBox("model")
}

func GetModuleTemplateBox() core.TemplateBox {
	return core.NewFileTemplateBox("module")
}

func GetProjectTemplateBox() core.TemplateBox {
	return core.NewFileTemplateBox("project")
}
