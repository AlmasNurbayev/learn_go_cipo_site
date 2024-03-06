package partParsers

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"database/sql"
	"time"
)

// ищет в структуре вложенную структуру "ТоварнаяГруппа" и возвращат ее элементы
func VidParser(receiveStruct *XMLTypes.ImportType, registrator_id int64) []models.VidsModel {

	mainStruct := (*receiveStruct)
	var vids []models.VidsModel

	root := mainStruct.КоммерческаяИнформация.Классификатор.Свойства.Свойство

	for i := 0; i < len(root); i++ {
		if root[i].Наименование == "ВидМодели" {

			for j := 0; j < len(root[i].ВариантыЗначений.Справочник); j++ {
				//var vid T
				vid := models.VidsModel{
					Id_1c:          root[i].ВариантыЗначений.Справочник[j].ИдЗначения,
					Name_1c:        root[i].ВариантыЗначений.Справочник[j].Значение,
					Registrator_id: registrator_id,
					Changed_at: sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					},
				}
				vids = append(vids, vid)
			}
		}
	}
	return vids
}
