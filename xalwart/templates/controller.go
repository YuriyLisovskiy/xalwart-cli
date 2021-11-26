package templates

import "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"

var ControllerTemplateBox = core.NewFileTemplateBoxWithTemplates(
	[]core.Template{
		core.NewFileTemplateWithContent(
			"_name_.h.txt", `<% .Header.CLikeCopyrightNotice %>

#pragma once

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>/controllers/controller.h>


class <% .ClassName %> : public <% .Header.FrameworkNamespace %>::ctrl::Controller<>
{
public:
	explicit inline <% .ClassName %>(const <% .Header.FrameworkNamespace %>::ILogger* logger) :
		<% .Header.FrameworkNamespace %>::ctrl::Controller<>({"get"}, logger)
	{
	}

	std::unique_ptr<<% .Header.FrameworkNamespace %>::http::IResponse> get(<% .Header.FrameworkNamespace %>::http::IRequest* request) const override;
};
`,
		),
		core.NewFileTemplateWithContent(
			"_name_.cpp.txt", `<% .Header.CLikeCopyrightNotice %>

#include "./<% .FileName %>.h"

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>/http/response.h>


std::unique_ptr<<% .Header.FrameworkNamespace %>::http::IResponse> <% .ClassName %>::get(<% .Header.FrameworkNamespace %>::http::IRequest* request) const
{
    // Implement controller logic here.

	return std::make_unique<<% .Header.FrameworkNamespace %>::http::Response>("Greetings from <% .ClassName %>!");
}
`,
		),
	},
)
