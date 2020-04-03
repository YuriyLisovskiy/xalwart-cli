package config

import (
	"strings"
	"text/template"
	"wasp-cli/internal/templates/project"
)

var DefaultFunctions = template.FuncMap {
	"title": strings.Title,
	"upper": strings.ToUpper,
	"lower": strings.ToLower,
}

var ProjectTemplates = []ProjectFile{
	{
		Name:     "settings.h",
		Path:     "PROJECT_NAME",
		TemplateStr: project.Template_ProjectName_SettingsH,
	},
	{
		Name:     "local_settings.cpp",
		Path:     "PROJECT_NAME",
		TemplateStr: project.Template_ProjectName_LocalSettingsCpp,
	},
	{
		Name:     "app.h",
		Path:     "main_app",
		TemplateStr: project.Template_MainApp_AppH,
	},
	{
		Name:     "views.h",
		Path:     "main_app",
		TemplateStr: project.Template_MainApp_ViewsH,
	},
	{
		Name:     "index.html",
		Path:     "templates",
		TemplateStr: project.Template_Templates_IndexHtml,
	},
	{
		Name:     "CMakeLists.txt",
		Path:     "",
		TemplateStr: project.Template_Root_CMakeListsTxt,
	},
	{
		Name:     "main.cpp",
		Path:     "",
		TemplateStr: project.Template_Root_MainCpp,
	},
}
