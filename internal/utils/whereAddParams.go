package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func WhereAddParams(selectQuery string, params map[string]interface{}) string {
	if len(params) > 0 {
		selectQuery = selectQuery + " WHERE "
	}
	var count = 0
	for key, value := range params {
		if value == nil {
			continue
		}
		count = count + 1
		if count > 1 {
			selectQuery = selectQuery + " AND "
		}
		if reflect.TypeOf(value) == reflect.TypeOf("") {
			selectQuery = selectQuery + key + " = " + fmt.Sprintf(`'%s'`, value)
		} else if isArrayOrSlice(value) {
			selectQuery = selectQuery + key + " IN " + SliceToWhereString(value)
		} else {
			selectQuery = selectQuery + key + " = " + fmt.Sprintf("%v", value)
		}
	}
	return selectQuery
}

func isArrayOrSlice(data interface{}) bool {
	dataType := reflect.TypeOf(data)
	kind := dataType.Kind()
	return kind == reflect.Array || kind == reflect.Slice
}

func SliceToWhereString(slice interface{}) string {
	// Преобразуем слайс в слайс строк
	var strSlice []string
	switch s := slice.(type) {
	case []int:
		for _, v := range s {
			strSlice = append(strSlice, fmt.Sprintf("%d", v))
		}
	case []string:
		strSlice = s
	default:
		return "Unsupported slice type"
	}

	// Объединяем строки через запятую и пробел
	return "(" + strings.Join(strSlice, ",") + ")"
}
