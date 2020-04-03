package project

const Template_MainApp_ViewsH = `#pragma once

#include <string>

#include <{{ .FrameworkName | lower }}/views/template_view.h>
#include <{{ .FrameworkName | lower }}/views/redirect_view.h>


class MainView : public {{ .FrameworkNamespace | lower }}::views::TemplateView
{
public:
	explicit MainView({{ .FrameworkNamespace | lower }}::conf::Settings* settings)
		: TemplateView({"get"}, settings)
	{
		this->_template_name = "index.html";
	}
};


class RedirectView : public {{ .FrameworkNamespace | lower }}::views::RedirectView
{
public:
	explicit RedirectView(wasp::conf::Settings* settings)
		: {{ .FrameworkNamespace | lower }}::views::RedirectView(settings, "/index", false, false)
	{
	};
};
`
