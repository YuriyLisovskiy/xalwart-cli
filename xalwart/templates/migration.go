package templates

import "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"

var MigrationTemplateBox = core.NewFileTemplateBoxWithTemplates(
	[]core.Template{
		core.NewFileTemplateWithContent(
			"_name_.h.txt", `<% .Header.CLikeCopyrightNotice %>

#pragma once

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>.orm/db/migration.h>


class <% .ClassName %> : public <% .Header.FrameworkNamespace %>::orm::db::Migration
{
public:
	explicit <% .ClassName %>(<% .Header.FrameworkNamespace %>::orm::IBackend* backend);
};
`,
		),
		core.NewFileTemplateWithContent(
			"_name_.cpp.txt", `<% .Header.CLikeCopyrightNotice %>

#include "./<% .FileName %>.h"


<% .ClassName %>::<% .ClassName %>(<% .Header.FrameworkNamespace %>::orm::IBackend* backend)
	: <% .Header.FrameworkNamespace %>::orm::db::Migration(backend, "<% .MigrationName %>", <% .IsInitial %>)
{
    // Register changes here.
}
`,
		),
	},
)
