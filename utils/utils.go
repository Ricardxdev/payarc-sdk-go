package utils

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func StructToForm(input interface{}) url.Values {
	form := url.Values{}
	val := reflect.ValueOf(input)
	typ := val.Type()

	if val.Kind() != reflect.Struct {
		fmt.Println(typ)
		return form
	}
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		formTag := field.Tag.Get("form")
		if formTag == "" || formTag == "-" {
			continue
		}

		tagParts := strings.Split(formTag, ",")
		tagName := tagParts[0]
		omitEmpty := len(tagParts) > 1 && tagParts[1] == "omitempty"

		fieldValue := val.Field(i)
		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				continue
			}
			fieldValue = fieldValue.Elem()
		}

		if omitEmpty && fieldValue.IsZero() {
			continue
		}

		if stringer, ok := fieldValue.Interface().(fmt.Stringer); ok {
			form.Set(tagName, stringer.String())
			continue
		}

		form.Set(tagName, fmt.Sprintf("%v", fieldValue.Interface()))
	}
	return form
}
