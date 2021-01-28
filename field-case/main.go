package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/stoewer/go-strcase"
)

func getTags() {
	type S struct {
		FB0 string `json:"foo_bar_0"`
		F1  string `json:"foo_bar_1"`
		F2  string `json:"fooBar2"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if alias, ok := field.Tag.Lookup("json"); ok {
			if alias == "" {
				fmt.Println("(blank)")
			} else {
				fmt.Println(alias, strcase.SnakeCase(alias), strcase.LowerCamelCase(alias))
			}
		} else {
			fmt.Println("(not specified)")
		}
	}
}

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

type FooStruct struct {
	F1 int `json:"foo_bar_1"`
	F2 int `json:"fooBar2"`
	F3 int `json:"foo"`
	F4 int
}

func main() {
	data := FooStruct{1, 1, 1, 5}
	bs1, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(bs1))

	parsedData := ToMapWithMultiCases(data)
	bs, _ := json.MarshalIndent(parsedData, "", "  ")
	fmt.Println(string(bs))
}
