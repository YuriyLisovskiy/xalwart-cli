package templates

import "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"

var MiddlewareTemplateBox = core.NewFileTemplateBoxWithTemplates(
	[]core.Template{
		core.NewFileTemplateWithContent(
			"_name_.h.txt", `<% .Header.CLikeCopyrightNotice %>

#pragma once

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>/middleware/base.h>

<% if .IsClassBased %>
class <% .FullName %>
{
public:
	<% .FullName %>();

	<% .Header.FrameworkNamespace %>::middleware::Function operator() (const <% .Header.FrameworkNamespace %>::middleware::Function& next);
};
<% else %>
extern <% .Header.FrameworkNamespace %>::middleware::Function <% .FullName %>(const <% .Header.FrameworkNamespace %>::middleware::Function& next);
<% end %>`,
		),
		core.NewFileTemplateWithContent(
			"_name_.cpp.txt", `<% .Header.CLikeCopyrightNotice %>

#include "./<% .FileName %>.h"

<% if .IsClassBased %>
<% .FullName %>::<% .FullName %>()
{
	// One-time configuration.
}

<% .Header.FrameworkNamespace %>::middleware::Function <% .FullName %>::operator() (const <% .Header.FrameworkNamespace %>::middleware::Function& next)
{
	// Pass only copies of any variables of current scope.
	return [*this, next](<% .Header.FrameworkNamespace %>::http::IRequest* request) -> std::unique_ptr<<% .Header.FrameworkNamespace %>::http::IResponse>
	{
		// Code to be executed for each request before
		// the controller (and later middleware) are called.

		auto response = next(request);

		// Code to be executed for each request/response after
		// the controller is called.

		return response;
	};
}
<% else %>
<% .Header.FrameworkNamespace %>::middleware::Function <% .FullName %>(const <% .Header.FrameworkNamespace %>::middleware::Function& next)
{
	return [next](<% .Header.FrameworkNamespace %>::http::IRequest* request) -> std::unique_ptr<<% .Header.FrameworkNamespace %>::http::IResponse>
	{
		// Code to be executed for each request before
        // the controller (and later middleware) are called.

        auto response = next(request);

        // Code to be executed for each request/response after
        // the controller is called.

        return response;
	};
}
<% end %>`,
		),
	},
)
