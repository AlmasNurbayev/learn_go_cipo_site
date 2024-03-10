package partParsers

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"time"
)

// ищет в структуре вложенную структуру "Размер2" и возвращает ее элементы
func SizesParser(receiveStruct *XMLTypes.OfferType, registrator_id int64) []models.Sizes {

	mainStruct := (*receiveStruct)
	var sizes []models.Sizes
	time := time.Now()

	root := mainStruct.КоммерческаяИнформация.Классификатор.Свойства.Свойство

	// TODO - в этой ветке стуктуры единственный массив, может быть не проверять Ид
	if root.Ид == "a001d8e3-a3b3-11ed-b0d2-50ebf624c538" {

		for j := 0; j < len(root.ВариантыЗначений.Справочник); j++ {
			//var vid T
			size := models.Sizes{
				Id_1c:          root.ВариантыЗначений.Справочник[j].ИдЗначения,
				Name_1c:        root.ВариантыЗначений.Справочник[j].Значение,
				Registrator_id: registrator_id,
				Changed_at:     &time,
			}
			sizes = append(sizes, size)
		}
	}
	return sizes
}
