package main

import (
	"errors"
	"reflect"
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
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values) == 0 {
				return nil, errors.New("invalid dict call")
			}

			dict := make(map[string]interface{})

			for i := 0; i < len(values); i++ {
				key, isset := values[i].(string)
				if !isset {
					if reflect.TypeOf(values[i]).Kind() == reflect.Map {
						m := values[i].(map[string]interface{})
						for i, v := range m {
							dict[i] = v
						}
					} else {
						return nil, errors.New("dict values must be maps")
					}
				} else {
					i++
					if i == len(values) {
						return nil, errors.New("specify the key for non array values")
					}
					dict[key] = values[i]
				}

			}
			return dict, nil
		},
		"CmpString": func(str1 string, str2Ptr *string) bool {
			var str2 string
			if str2Ptr != nil {
				str2 = *str2Ptr
			}
			if str1 == str2 {
				return true
			}
			return false
		},
		"HasField": func(v interface{}, name string) bool {
			rv := reflect.ValueOf(v)
			if rv.Kind() == reflect.Ptr {
				rv = rv.Elem()
			}
			if rv.Kind() != reflect.Struct {
				return false
			}
			return rv.FieldByName(name).IsValid()
		},
	}

	for key, fn := range tmplFuncs {
		beego.AddFuncMap(key, fn)
	}
	beego.Run()
}
