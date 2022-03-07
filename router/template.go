package router

import "text/template"

var (
	defaultFuncs = template.FuncMap{
		"defTitle": func(ip interface{}) string {
			v, ok := ip.(string)

			if !ok || (ok && v == "") {
				return "Go on Plan 9"
			}
			return v
		},
	}
	templateFiles = []string{
		"./web/templates/base.gohtml",
	}
)

func templLayout(files ...string) []string {
	return append(templateFiles, files...)
}
