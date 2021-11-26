package templates

import "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"

var ModelTemplateBox = core.NewFileTemplateBoxWithTemplates(
	[]core.Template{
		core.NewFileTemplateWithContent(
			"_name_.h.txt", `<% .Header.CLikeCopyrightNotice %>

#pragma once

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>.orm/db/model.h>


class <% .ClassName %> : public <% .Header.FrameworkNamespace %>::orm::db::Model<% if .IsJsonSerializable %>, public <% .Header.FrameworkNamespace %>::IJsonSerializable<% end %>
{
public:
	long long int id{};

	static constexpr const char* meta_table_name = "<% .TableName %>";

	inline static const std::tuple meta_columns = {
		<% .Header.FrameworkNamespace %>::orm::db::make_pk_column_meta("id", &<% .ClassName %>::id)
	};

	<% .ClassName %>() = default;

	inline void __orm_set_column__(const std::string& column_name, const char* data) override
	{
		this->__orm_set_column_data__(<% .ClassName %>::meta_columns, column_name, data);
	}

	[[nodiscard]]
	std::string to_string() const override;<% if .IsJsonSerializable %>

	[[nodiscard]]
	nlohmann::json to_json() const override;<% end %>
};
`,
		),
		core.NewFileTemplateWithContent(
			"_name_.cpp.txt", `<% .Header.CLikeCopyrightNotice %>

#include "./<% .FileName %>.h"


std::string <% .ClassName %>::to_string() const
{
    if (this->is_null())
    {
     	return "null";
    }

    return "<% .Name | to_camel_case %> [" + std::to_string(this->id) + "]";
}
<% if .IsJsonSerializable %>

nlohmann::json <% .ClassName %>::to_json() const
{
	if (this->is_null())
	{
		return nullptr;
	}

	return {{"id", this->id}};
}
<% end %>`,
		),
	},
)
