package config

import (
	"strings"
	"text/template"
)

var DefaultFunctions = template.FuncMap {
//	"title": strings.Title,
	"upper": strings.ToUpper,
//	"lower": strings.ToLower,
}
