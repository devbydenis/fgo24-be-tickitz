package utils

import (
	"fmt"
	"reflect"
)

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0.0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Struct:
		// Untuk embedded struct, check semua field-nya
		for i := 0; i < v.NumField(); i++ {
			if !isEmptyValue(v.Field(i)) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func CheckFieldValues(s interface{}) string {
	v := reflect.ValueOf(s)
	t := v.Type()

	result := ""
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Skip unexported fields
		if !value.CanInterface() {
			continue
		}

		isEmpty := isEmptyValue(value)
		if isEmpty {
			result = fmt.Sprintf("%s is required", field.Name)
			break
		}
	}

	return result
}

func ToNullString(s string) interface{} {
    if s == "" {
        return nil
    }
    return s
}