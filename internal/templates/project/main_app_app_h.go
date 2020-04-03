package project

const Template_MainApp_AppH = `#pragma once

#include <{{ .FrameworkName | lower }}/apps/config.h>
#include <{{ .FrameworkName | lower }}/conf/settings.h>

#include "./views.h"


class MainAppConfig : public {{ .FrameworkNamespace | lower }}::apps::AppConfig
{
public:
	explicit MainAppConfig({{ .FrameworkNamespace | lower }}::conf::Settings* settings)
		: AppConfig(__FILE__, settings)
	{
		this->init(this->__type__());
	}

	void urlpatterns() override
	{
		this->url<MainView>(R"(index/?)", "index");
		this->url<RedirectView>(R"(/?)", "root");
	}
};
`
