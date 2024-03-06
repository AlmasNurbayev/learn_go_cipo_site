package operations1

import (
	"cipo_cite_server/internal/storage/postgres"
	"fmt"
	"log/slog"
	"reflect"
	"slices"
	"strings"
)

// вставка в БД, передаем:
// 1. db - экземпляр клиента БД,
// 2. log - экземпляр логера,
// 3. tableName - имя таблицы в БД,
// 4. data - структура для вставки,
// 5. skippedFields - поля, которые нужно пропустить при вставке,
// возвращаем: id новой записи или ошибку
func Insert(dbStore *postgres.PostgresStore,
	log *slog.Logger,
	tableName string,
	data interface{},
	skippedFields []string,
) (int64, error) {

	dataType := reflect.TypeOf(data)

	// Проверяем, что data является структурой
	if dataType.Kind() != reflect.Struct {
		log.Error("ожидалась структура, получен %v", dataType.Kind())
		return 0, fmt.Errorf("ожидалась структура, получен %v", dataType.Kind())
	}

	// Создаем пустой слайс для значений структуры
	values := make([]interface{}, 0)

	// Создаем строку для форматирования запроса
	query := "INSERT INTO " + tableName + " ("

	// Получаем количество полей в структуре
	numFields := dataType.NumField()

	// Добавляем имена полей в запрос
	for i := 0; i < numFields; i++ {
		field := dataType.Field(i)
		if slices.Contains(skippedFields, field.Name) {
			continue
		}
		query += strings.ToLower(field.Name[:1]) + field.Name[1:] + ","
		values = append(values, reflect.ValueOf(data).Field(i).Interface())
	}

	// Удаляем последнюю запятую и добавляем закрывающую скобку
	query = query[:len(query)-1] + ") VALUES ("

	// Добавляем плейсхолдеры для значений
	for i := 0; i < numFields; i++ {
		field := dataType.Field(i)
		if slices.Contains(skippedFields, field.Name) {
			continue
		}
		query += ":" + strings.ToLower(field.Name[:1]) + field.Name[1:] + ","
	}

	// Удаляем последнюю запятую и добавляем закрывающую скобку
	query = query[:len(query)-1] + ") RETURNING id"
	fmt.Println("query", query)

	rows, err := dbStore.Dbx.NamedQuery(query, data)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var res int64

	for rows.Next() {
		err := rows.Scan(&res)
		if err != nil {
			return 0, err
		}
		//utils.PrintAsJSON(res)
	}
	return res, nil
}
