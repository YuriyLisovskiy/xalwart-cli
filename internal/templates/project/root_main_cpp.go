package project

const Template_Root_MainCpp = `#include <{{ .FrameworkName | lower }}/apps/{{ .FrameworkName | lower }}.h>

#include "./{{.ProjectName}}/settings.h"


int main(int argc, char** argv)
{
	auto settings = std::make_shared<Settings>();
	try
	{
		auto app = {{ .FrameworkNamespace }}::apps::{{ .FrameworkName | title }}Application(settings.get());
		app.execute(argc, argv);
	}
	catch (const {{ .FrameworkNamespace }}::core::ImproperlyConfigured& exc)
	{
		if (settings->LOGGER)
		{
			settings->LOGGER->error(exc);
		}
	}

	return 0;
}
`
