package partParsers

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"time"
)

// ищет в структуре вложенную структуру "Склады" и возвращает ее элементы
func StoresParser(receiveStruct *XMLTypes.OfferType, registrator_id int64) []models.Stores {

	mainStruct := (*receiveStruct)
	var stores []models.Stores
	time := time.Now()

	root := mainStruct.КоммерческаяИнформация.ПакетПредложений.Склады.Склад

	for j := 0; j < len(root); j++ {
		//var vid T
		store := models.Stores{
			Id_1c:          root[j].Ид,
			Name_1c:        root[j].Наименование,
			Registrator_id: registrator_id,
			Changed_at:     &time,
		}
		stores = append(stores, store)
	}

	return stores
}
