package templates

import "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"

var CommandTemplateBox = core.NewFileTemplateBoxWithTemplates([]core.Template{
	core.NewFileTemplateWithContent(
		"_name_.h.txt", `<% .Header.CLikeCopyrightNotice %>

#pragma once

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>.base/interfaces/base.h>
#include <<% .Header.FrameworkName %>/commands/command.h>


class <% .ClassName %> : public <% .Header.FrameworkNamespace %>::cmd::Command
{
public:
	inline explicit <% .ClassName %>(const std::shared_ptr<<% .Header.FrameworkNamespace %>::ILogger>& logger) :
        <% .Header.FrameworkNamespace %>::cmd::Command("<% .Name | to_snake_case %>", "Calls '<% .Name | to_snake_case %>' command", logger)
	{
	}

protected:
	void add_flags() override;

	void validate() const override;

	bool handle() override;
};
`,
	),
	core.NewFileTemplateWithContent(
		"_name_.cpp.txt", `<% .Header.CLikeCopyrightNotice %>

#include "./<% .FileName %>.h"


void <% .ClassName %>::add_flags()
{
	<% .Header.FrameworkNamespace %>::cmd::Command::add_flags();

	// Register flags here.
}

void <% .ClassName %>::validate() const
{
	<% .Header.FrameworkNamespace %>::cmd::Command::validate();

	// Perform validation of registered flags here.
}

bool <% .ClassName %>::handle()
{
	if (<% .Header.FrameworkNamespace %>::cmd::Command::handle())
	{
		return true;
	}

	// Implement command logic here.
	this->logger->info("Greetings from <% .ClassName %>!");

	return true;
}
`,
	),
})
