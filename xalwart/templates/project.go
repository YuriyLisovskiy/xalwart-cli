package templates

import "github.com/YuriyLisovskiy/xalwart-cli/xalwart/core"

var ProjectTemplateBox = core.NewFileTemplateBoxWithTemplates(
	[]core.Template{
		// src/
		core.NewFileTemplateWithContent(
			"src/main.cpp.txt", `<% .Header.CLikeCopyrightNotice %>

// STL
#include <iostream>

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>/conf/application.h>

// <% .ProjectName %>
#include "./config/settings.h"


int main(int argc, char** argv)
{
	<% .Header.FrameworkNamespace %>::conf::initialize_signal_handlers();
	std::unique_ptr<Settings> settings;
	try
	{
		settings = Settings::load();
		<% .Header.FrameworkNamespace %>::conf::Application(settings.get())
			.configure()
			.execute(argc, argv);
	}
	catch (const <% .Header.FrameworkNamespace %>::ImproperlyConfigured& exc)
	{
		if (settings && settings->LOGGER)
		{
			settings->LOGGER->error(exc);
		}
		else
		{
			std::cerr << exc.get_message() << std::endl;
		}
	}

	return 0;
}
`,
		),
		core.NewFileTemplateWithContent(
			"src/settings.local.yaml.txt", `<% .Header.NumberSignCopyrightNotice %>
# Do not push this configuration file to public repositories.

secret_key: '<% .SecretKey %>'

allowed_hosts:
  - '*'
`,
		),
		core.NewFileTemplateWithContent(
			"src/settings.yaml.txt", `<% .Header.NumberSignCopyrightNotice %>

debug: true

secret_key: 'this is set in local settings'

allowed_hosts:
  - 127.0.0.1
  - ::1

timezone:
  name: UTC

modules:
  - MainModule

middleware:
  - <% .Header.FrameworkNamespace %>::middleware::Security
  - <% .Header.FrameworkNamespace %>::middleware::Common
  - <% .Header.FrameworkNamespace %>::middleware::XFrameOptions
<% if .UseStandardORM %>
databases:
  - instance: default
    dbms: sqlite3
    file: db.sqlite3
    connections: 3
<% end %>
logger:
  use_colors: true
  levels: '*'
`,
		),

		// src/config/
		core.NewFileTemplateWithContent(
			"src/config/module.cpp.txt", `<% .Header.CLikeCopyrightNotice %>

#include "./module.h"

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>/controllers/redirect.h>

// <% .ProjectName %>
#include "./controllers/index.h"


void MainModule::configure()
{
    // Set up module here instead of in constructor

	this->set_ready();
}

void MainModule::urlpatterns()
{
    // Register URL patterns here.

    this->url<IndexController>(R"(index/?)", "index");
    this->url<<% .Header.FrameworkNamespace %>::ctrl::RedirectController>(R"(/?)", "root", "/index", false, false);
}
`,
		),
		core.NewFileTemplateWithContent(
			"src/config/module.h.txt", `<% .Header.CLikeCopyrightNotice %>

#pragma once

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>/conf/module.h>


class MainModule : public <% .Header.FrameworkNamespace %>::conf::ModuleConfig
{
public:
	explicit inline MainModule(const std::string& registration_name, <% .Header.FrameworkNamespace %>::conf::Settings* settings) :
        <% .Header.FrameworkNamespace %>::conf::ModuleConfig(registration_name, settings)
    {
    }

	void configure() override;

	void urlpatterns() override;
};
`,
		),
		core.NewFileTemplateWithContent(
			"src/config/settings.cpp.txt", `<% .Header.CLikeCopyrightNotice %>

#include "./settings.h"

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>.base/workers/threaded_worker.h><% if .UseStandardORM %>
#include <<% .Header.FrameworkName %>.orm/config/yaml.h><% end %><% if .UseStandardServer %>
#include <<% .Header.FrameworkName %>.server/http_server.h><% end %>
#include <<% .Header.FrameworkName %>/conf/loaders/yaml_loader.h>

// <% .ProjectName %>
#include "./module.h"


void Settings::register_modules()
{
	this->module<MainModule>();
}

std::unique_ptr<<% .Header.FrameworkNamespace %>::server::IServer> Settings::build_server(
    const std::function<<% .Header.FrameworkNamespace %>::net::StatusCode(
	    <% .Header.FrameworkNamespace %>::net::RequestContext*, const std::map<std::string, std::string>& /* environment */
    )>& handler,
    const <% .Header.FrameworkNamespace %>::Options& options
)
{<% if .UseStandardServer %>
	return std::make_unique<<% .Header.FrameworkNamespace %>::server::DevelopmentHTTPServer>(<% .Header.FrameworkNamespace %>::server::Context{
		.logger = this->LOGGER.get(),
		.timezone = this->TIMEZONE,
		.max_headers_count = this->LIMITS.MAX_HEADERS_COUNT,
		.max_header_length = this->LIMITS.MAX_HEADER_LENGTH,
		.timeout_seconds = options.get<time_t>("timeout_seconds"),
		.timeout_microseconds = options.get<time_t>("timeout_microseconds"),
		.socket_creation_retries_count = options.get<size_t>("retries"),
		.worker = std::make_unique<<% .Header.FrameworkNamespace %>::ThreadedWorker>(options.get<size_t>("workers")),
		.handler = handler
	});<% else %>
    // Create and return server here.
    return nullptr;<% end %>
}

std::unique_ptr<Settings> Settings::load()
{
	return <% .Header.FrameworkNamespace %>::conf::YAMLSettingsLoader<Settings>({"settings", R"(settings\.local)"})
		.with_components([](auto* loader, auto* settings)
		{
			// Registration of logger should be done first
			// to log errors during the next components' setup.
			loader->register_default_logger(settings);

			// Setup standard components which are present in
			// <% .Header.FrameworkNamespace %>::conf::Settings excluding template engine
			// and databases.
			loader->register_standard_components(settings);

			// Setup components from external libraries.<% if .UseStandardORM %>
			loader->register_component("databases", std::make_unique<<% .Header.FrameworkNamespace %>::orm::config::YAMLDatabasesComponent>(
				settings->BASE_DIR, settings->DATABASES
			));
			<% else %>
			<% end %>
			// Other custom components' setup.
		})
		.load();
}
`,
		),
		core.NewFileTemplateWithContent(
			"src/config/settings.h.txt", `<% .Header.CLikeCopyrightNotice %>

#pragma once

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>/conf/settings.h>


class Settings : public <% .Header.FrameworkNamespace %>::conf::Settings
{
public:
	void register_modules() override;

	std::unique_ptr<<% .Header.FrameworkNamespace %>::server::IServer> build_server(
	    const std::function<<% .Header.FrameworkNamespace %>::net::StatusCode(
		    <% .Header.FrameworkNamespace %>::net::RequestContext*, const std::map<std::string, std::string>& /* environment */
	    )>& handler,
	    const <% .Header.FrameworkNamespace %>::Options& options
	) override;

	static std::unique_ptr<Settings> load();
};
`,
		),

		// src/config/controllers/
		core.NewFileTemplateWithContent(
			"src/config/controllers/index.cpp.txt", `<% .Header.CLikeCopyrightNotice %>

#include "./index.h"

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>/http/response.h>


std::unique_ptr<<% .Header.FrameworkNamespace %>::http::IResponse> IndexController::get(<% .Header.FrameworkNamespace %>::http::IRequest* request) const
{
    // Implement controller logic here.

	return std::make_unique<<% .Header.FrameworkNamespace %>::http::Response>("Hello, World!");
}
`,
		),
		core.NewFileTemplateWithContent(
			"src/config/controllers/index.h.txt", `<% .Header.CLikeCopyrightNotice %>

#pragma once

// <% .Header.FrameworkName %>
#include <<% .Header.FrameworkName %>/controllers/controller.h>


class IndexController : public <% .Header.FrameworkNamespace %>::ctrl::Controller<>
{
public:
	explicit inline IndexController(const <% .Header.FrameworkNamespace %>::ILogger* logger) :
		<% .Header.FrameworkNamespace %>::ctrl::Controller<>({"get"}, logger)
	{
	}

	std::unique_ptr<<% .Header.FrameworkNamespace %>::http::IResponse> get(<% .Header.FrameworkNamespace %>::http::IRequest* request) const override;
};
`,
		),

		// root
		core.NewFileTemplateWithContent(
			".gitignore.txt", `# Prerequisites
*.d

# Compiled Object files
*.slo
*.lo
*.o
*.obj

# Precompiled Headers
*.gch
*.pch

# Compiled Dynamic libraries
*.so
*.dylib
*.dll

# Fortran module files
*.mod
*.smod

# Compiled Static libraries
*.lai
*.la
*.a
*.lib

# Executables
*.exe
*.out
*.app

# Visual Studio
.vs/
out/
CMakeSettings.json

# CLion
.idea/
cmake-build-*
build/

# Xalwart Framework
media/
settings.local.yaml

# Other
db.sqlite3

# macOS files
*.DS_Store
`,
		),
		core.NewFileTemplateWithContent(
			"CMakeLists.txt.txt", `<% .Header.NumberSignCopyrightNotice %>

cmake_minimum_required(VERSION 3.12)

set(CMAKE_CXX_STANDARD 20)

project(<% .ProjectName %>)

set(BINARY application)

set(CMAKE_CXX_FLAGS "-pthread")

set(DEFAULT_INCLUDE_PATHS "/usr/local" "/usr")

foreach(ENTRY ${DEFAULT_INCLUDE_PATHS})
    include_directories(${ENTRY}/include)
    link_directories(${ENTRY}/lib)
endforeach()

# Search for OpenSSL
find_package(OpenSSL 1.1 REQUIRED)
include_directories(${OPENSSL_INCLUDE_DIR})

# Load and filter project sources.
file(GLOB_RECURSE SOURCES LIST_DIRECTORIES true src/*.h src/*.cpp)
list(
	FILTER SOURCES
	EXCLUDE REGEX "^.*/(include|lib$|media|static|templates|cmake-build-*)/?.*"
)
foreach(entry ${SOURCES})
    if (IS_DIRECTORY ${entry})
        list(REMOVE_ITEM SOURCES ${entry})
    endif()
endforeach()

add_executable(${BINARY} ${SOURCES})
if (NOT APPLE)
    target_link_libraries(${BINARY} PUBLIC stdc++fs)
endif()
<% if .UseStandardORM %>
option(XW_USE_SQLITE3 "Link sqlite3 library and add compile XW_USE_SQLITE3 compile definition." OFF)
if (${XW_USE_SQLITE3})
    # Search for sqlite3
    find_library(
        SQLITE3 sqlite3 REQUIRED
        PATHS ${DEFAULT_INCLUDE_PATHS}
    )
    target_link_libraries(${BINARY} PUBLIC ${SQLITE3})
    add_compile_definitions(USE_SQLITE3)
endif()

if (${XW_USE_POSTGRESQL})
    # Search for PostgreSQL
    find_package(PostgreSQL REQUIRED)
    target_link_libraries(${BINARY} PUBLIC ${PostgreSQL_LIBRARIES})
    include_directories(${PostgreSQL_INCLUDE_DIRS})
    add_compile_definitions(USE_POSTGRESQL)
endif()
<% end %>
set(XALWART_LIBRARIES)
function(FIND_XALWART_LIBRARY PART)
    find_library(
        XALWART_${PART} xalwart${PART} REQUIRED
        PATHS ${DEFAULT_INCLUDE_PATHS}
    )
    set(XALWART_LIBRARIES ${XALWART_LIBRARIES} ${XALWART_${PART}} PARENT_SCOPE)
endfunction(FIND_XALWART_LIBRARY)

find_xalwart_library(.base)
find_xalwart_library(.crypto)<% if .UseStandardORM %>
find_xalwart_library(.orm)<% end %><% if .UseStandardServer %>
find_xalwart_library(.server)<% end %>
find_xalwart_library("")

target_link_libraries(${BINARY} PUBLIC ${OPENSSL_LIBRARIES} ${XALWART_LIBRARIES})
`,
		),
		core.NewFileTemplateWithContent(
			"Dockerfile.txt", `# Simple configuration for quick check the application.
#
# CMake configuration log and build log are available
# in /var/log/app after building the container as configure.log
# and build.log respectfully.

FROM <% .FrameworkBaseDockerImage %>

ENV WORK_DIR=/app

# Set the working directory.
WORKDIR $WORK_DIR

# Copy project files.
COPY CMakeLists.txt ./
COPY ./src ./src

# Create directories for build files and logs.
RUN mkdir -p ./build && mkdir -p /var/log/app

# Configure and build the application.
# Remark: to make the database changes persistent use PostgreSQL or another database with a server.
#         For using PostgreSQL add '-D XW_USE_POSTGRESQL=yes' option.
RUN cd ./build && \
    cmake -D CMAKE_C_COMPILER=clang \
          -D CMAKE_CXX_COMPILER=clang++ \
          -D CMAKE_BUILD_TYPE=Release \
          -D XW_USE_SQLITE3=yes \
          .. \
          2>&1 | tee /var/log/app/configure.log && \
    make application 2>&1 | tee /var/log/app/build.log

WORKDIR $WORK_DIR

# Copy compiled app and YAML configurations to the working directory.
RUN cp ./build/application ./ && cp ./src/*.yaml ./

# Clean up the working directory.
RUN rm -rf ./CMakeLists.txt ./src ./build

CMD ["bash"]
`,
		),
		core.NewFileTemplateWithContent(
			"README.md.txt", `## <% .ProjectName %>

### Application commands
* `+"`start-server`"+` - start a web application on the local machine:
  `+"```"+`bash
  ./application start-server
  `+"```"+`
* `+"`migrate`"+` - migrate changes to the database:
  `+"```"+`bash
  ./application migrate
  `+"```"+`

For more information about application commands, run:
`+"```"+`bash
# list all available commands
./application

# get command usage info
./application [command] --help
`+"```"+`

### Build and run
Build the Docker container and run the application:
`+"```"+`bash
sudo docker build -t <% .ProjectName | to_snake_case %>:latest .
docker run -p 8000:8000 <% .ProjectName | to_snake_case %>:latest ./application start-server --bind 0.0.0.0:8000 --workers=5
`+"```"+`
`,
		),
	},
)
