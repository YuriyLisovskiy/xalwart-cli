package templates

import "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"

var CopyrightNoticesTemplateBox = core.NewFileTemplateBoxWithTemplates(
	[]core.Template{
		core.NewFileTemplateWithContent(
			"c-like.txt", `/*
 * Copyright (c) <% .Year %> <% .UserName %>
 */`,
		),
		core.NewFileTemplateWithContent("number-sign.txt", `# Copyright (c) <% .Year %> <% .UserName %>`),
	},
)
