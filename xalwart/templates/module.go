package templates

import "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"

var ModuleTemplateBox = core.NewFileTemplateBoxWithTemplates(
	[]core.Template{
		core.NewFileTemplateWithContent(
			"_name_/module.h.txt", `<% .Header.CLikeCopyrightNotice %>

#pragma once

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>/conf/module.h>


class <% .ClassName %> : public <% .Header.FrameworkNamespace %>::conf::ModuleConfig
{
public:
	explicit inline <% .ClassName %>(const std::string& registration_name, <% .Header.FrameworkNamespace %>::conf::Settings* settings) :
        <% .Header.FrameworkNamespace %>::conf::ModuleConfig(registration_name, settings)
    {
    }

	void configure() override;

	void urlpatterns() override;
};
`,
		),
		core.NewFileTemplateWithContent(
			"_name_/module.cpp.txt", `<% .Header.CLikeCopyrightNotice %>

#include "./module.h"


void <% .ClassName %>::configure()
{
    // Set up module here instead of in constructor

	this->set_ready();
}

void <% .ClassName %>::urlpatterns()
{
    // Register URL patterns here.
}
`,
		),
	},
)
