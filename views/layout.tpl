{{if .LoginUser}}
		{{template "authlogin/nav.tpl" .}}
{{end}}
{{.AdminContent}}
