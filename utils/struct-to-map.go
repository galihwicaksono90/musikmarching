package utils

import "reflect"

func structToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(obj)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		result[typ.Field(i).Name] = field.Interface()
	}
	return result
} 
