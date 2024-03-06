package partParsers

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"database/sql"
	"time"
)

// ищет в структуре вложенную структуру "ТоварнаяГруппа" и возвращат ее элементы
func ProductVidsParser(receiveStruct *XMLTypes.ImportType, registrator_id int64) []models.ProductVids {

	mainStruct := (*receiveStruct)
	var productVids []models.ProductVids

	root := mainStruct.КоммерческаяИнформация.Классификатор.Группы.Группа

	for i := 0; i < len(root); i++ {
		productVid := models.ProductVids{
			Id_1c:          root[i].Ид,
			Name_1c:        root[i].Наименование,
			Registrator_id: registrator_id,
			Changed_at: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		}
		productVids = append(productVids, productVid)
	}
	return productVids
}
