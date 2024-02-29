package utils

import (
	"fmt"
	"strings"
)

// Добавляем в строку SQL-запроса переданные параметры
// Пример: передали "SELECT * FROM STORE" и map[string]interface{}{"id": 19, "name_1c": "name"}
// получили "SELECT * FROM STORE WHERE id = 19 AND name_1c = 'name'"
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

		switch value.(type) {
		case int:
			selectQuery = selectQuery + key + " = " + fmt.Sprintf("%v", value)
		case string:
			selectQuery = selectQuery + key + " = " + fmt.Sprintf(`'%s'`, value)
		case []int, []string:
			selectQuery = selectQuery + key + " IN " + SliceToWhereString(value)
		default:
			selectQuery = selectQuery + key + " = " + fmt.Sprintf("%v", value)
		}
	}
	return selectQuery
}

// UNUSED
// Проверяем является ли переданный аргумент массивом или слайсом
// func isArrayOrSlice(data interface{}) bool {
// 	dataType := reflect.TypeOf(data)
// 	kind := dataType.Kind()
// 	return kind == reflect.Array || kind == reflect.Slice
// }

// Конвертируем слайс с string/Int в строку для передачи в Where
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
	// Объединяем строки через запятую
	return "(" + strings.Join(strSlice, ",") + ")"
}
