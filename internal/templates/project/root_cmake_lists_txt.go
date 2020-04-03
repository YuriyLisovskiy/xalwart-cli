package project

const Template_Root_CMakeListsTxt = `cmake_minimum_required(VERSION {{ .CMakeMinimumVersion }})
set(CMAKE_CXX_STANDARD {{.CMakeCPPStandard}})
project({{.ProjectName}})

set(BINARY "${{ "{" }}CMAKE_PROJECT_NAME{{ "}" }}-main")

# Setup application sources.
file(GLOB_RECURSE SOURCES LIST_DIRECTORIES true ./*/*.h ./*/*.cpp)
set(SOURCES ${{ "{" }}SOURCES{{ "}" }})

add_executable(${{ "{" }}BINARY{{ "}" }} ${{ "{" }}SOURCES{{ "}" }} "./main.cpp")

# Link {{ .FrameworkName | title }} library.
find_library({{ .FrameworkName | upper }}_LIBRARY {{ .FrameworkName | lower }})
if({{ .FrameworkName | upper }}_LIBRARY)
    target_link_libraries(${{ "{" }}BINARY{{ "}" }} ${{ "{" }}{{ .FrameworkName | upper }}_LIBRARY{{ "}" }})
else({{ .FrameworkName | upper }}_LIBRARY_NOT_FOUND)
    message(FATAL_ERROR "{{ .FrameworkName | title }} Framework is not found. Make sure if it is installed properly.")
endif({{ .FrameworkName | upper }}_LIBRARY)
`
