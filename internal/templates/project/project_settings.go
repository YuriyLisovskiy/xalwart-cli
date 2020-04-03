package project

const(
	Template_ProjectName_SettingsH = `#pragma once

#include <{{ .FrameworkName | lower }}/conf/settings.h>
#include <{{ .FrameworkName | lower }}/core/path.h>

#include <{{ .FrameworkName | lower }}/middleware/common.h>
#include <{{ .FrameworkName | lower }}/middleware/cookies.h>
#include <{{ .FrameworkName | lower }}/middleware/security.h>
#include <{{ .FrameworkName | lower }}/middleware/clickjacking.h>

#include <{{ .FrameworkName | lower }}/render/env/default.h>
#include <{{ .FrameworkName | lower }}/render/loaders.h>
#include <{{ .FrameworkName | lower }}/render/library/builtin.h>

#include "../main_app/app.h"


struct Settings final: public {{ .FrameworkNamespace }}::conf::Settings
{
	Settings() : {{ .FrameworkNamespace }}::conf::Settings()
	{
	}

	void init() final
	{
		using namespace {{ .FrameworkNamespace }};

		this->BASE_DIR = core::path::dirname(core::path::dirname(__FILE__));

		this->SECRET_KEY = "{{ .SecretKey }}";

		this->DEBUG = true;

		this->ALLOWED_HOSTS = {"127.0.0.1", "::1"};

		this->INSTALLED_APPS = {
			this->app<MainAppConfig>()
		};

		this->MIDDLEWARE = {
			this->middleware<middleware::SecurityMiddleware>(),
			this->middleware<middleware::CommonMiddleware>(),
			this->middleware<middleware::XFrameOptionsMiddleware>(),
			this->middleware<middleware::CookieMiddleware>()
		};

		this->TEMPLATES_ENV = render::env::DefaultEnvironment::Config{
			.dirs = std::vector<std::string>{
				core::path::join(this->BASE_DIR, "templates")
			},
			.use_app_dirs		= true,
			.apps				= this->INSTALLED_APPS,
			.debug				= this->DEBUG,
			.logger				= this->LOGGER.get(),
			.auto_escape		= true,
			.libraries = std::vector<std::shared_ptr<render::lib::ILibrary>>{
				this->library<render::lib::BuiltinLibrary>()
			}
		}.make_env();

		this->MEDIA_ROOT = core::path::join(this->BASE_DIR, "media");
		this->MEDIA_URL = "/media/";

		this->STATIC_ROOT = core::path::join(this->BASE_DIR, "static");
		this->STATIC_URL = "/static/";

		this->DATA_UPLOAD_MAX_MEMORY_SIZE = 20971520;
	}

	// Implement in local_settings.cpp!
	void override() final;
};
`

	Template_ProjectName_LocalSettingsCpp = `#include "./settings.h"


void Settings::override()
{
	// TODO: override default settings
}
`
)
