package utils

import (
	"reflect"

	"github.com/go-basic/uuid"
)

// Struct2Map changes struct to map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	data := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func UUID() string {
	uuid := uuid.New()
	return uuid
}
