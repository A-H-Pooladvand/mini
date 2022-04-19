package structs

import (
	"reflect"
)

func FieldNames(i any) (fields []string) {
	val := reflect.ValueOf(i).Elem()

	for i := 0; i < val.NumField(); i++ {
		fields = append(fields, val.Type().Field(i).Name)
	}

	return fields
}
