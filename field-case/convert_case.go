package main

import (
	"reflect"

	"github.com/stoewer/go-strcase"
)

func ToMapWithMultiCases(data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	st := reflect.ValueOf(data)

	for i := 0; i < st.NumField(); i++ {
		valueField := st.Field(i)
		typeField := st.Type().Field(i)

		if tagName, ok := typeField.Tag.Lookup("json"); ok {
			camelCase := strcase.LowerCamelCase(tagName)
			snakeCase := strcase.SnakeCase(tagName)

			res[camelCase] = valueField.Interface()
			res[snakeCase] = valueField.Interface()
		} else {
			res[typeField.Name] = valueField.Interface()
		}
	}

	return res
}
