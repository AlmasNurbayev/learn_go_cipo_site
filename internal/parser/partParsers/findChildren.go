package partParsers

import "reflect"

// рекурсивный поиск значения в структуре - проблема в типизации результата
func findInStructRecursive(data interface{}, value interface{}) (any, bool) {
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			if _, ok := findInStructRecursive(field.Interface(), value); ok {
				return field.Interface(), true
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i)
			if _, ok := findInStructRecursive(elem.Interface(), value); ok {
				return elem.Interface(), true
			}
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			elem := val.MapIndex(key)
			if _, ok := findInStructRecursive(elem.Interface(), value); ok {
				return elem.Interface(), true
			}
		}
	default:
		if reflect.DeepEqual(val.Interface(), value) {
			return val.Interface(), true
		}
	}

	return nil, false
}
