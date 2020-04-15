package main

import (
	"smarthead/internal/platform/web/weberror"
	_ "smarthead/routers"
	"text/template"

	"github.com/astaxie/beego"
)

func main() {
	tmplFuncs := template.FuncMap{
		"ValidationFieldClass": func(err interface{}, fieldName string) string {
			if err == nil {
				return ""
			}
			verr, ok := err.(*weberror.Error)
			if !ok || len(verr.Fields) == 0 {
				return ""
			}

			for _, e := range verr.Fields {
				if e.Field == fieldName || e.FormField == fieldName {
					return "is-invalid"
				}
			}
			return "is-valid"
		},
	}

	for key, fn := range tmplFuncs {
		beego.AddFuncMap(key, fn)
	}
	beego.Run()
}
