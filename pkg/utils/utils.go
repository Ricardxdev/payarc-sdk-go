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

// StructToQuery converts a struct with `query` tags into a map of query
// parameters. Fields tagged with `query:"-"` or without a tag are skipped.
// Pointer fields that are nil are skipped, and `omitempty` skips zero values.
// Slices are joined with commas and booleans are encoded as 1/0.
func StructToQuery(input interface{}) map[string]string {
	params := map[string]string{}
	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return params
		}
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return params
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		queryTag := field.Tag.Get("query")
		if queryTag == "" || queryTag == "-" {
			continue
		}

		tagParts := strings.Split(queryTag, ",")
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

		switch fieldValue.Kind() {
		case reflect.Slice:
			items := make([]string, 0, fieldValue.Len())
			for j := 0; j < fieldValue.Len(); j++ {
				items = append(items, fmt.Sprintf("%v", fieldValue.Index(j).Interface()))
			}
			if len(items) > 0 {
				params[tagName] = strings.Join(items, ",")
			}
		case reflect.Bool:
			if fieldValue.Bool() {
				params[tagName] = "1"
			} else {
				params[tagName] = "0"
			}
		default:
			params[tagName] = fmt.Sprintf("%v", fieldValue.Interface())
		}
	}
	return params
}
