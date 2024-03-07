package partParsers

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"time"
)

// ищет в структуре вложенную структуру "ТоварнаяГруппа" и возвращат ее элементы
func ProductGroupsParser(receiveStruct *XMLTypes.ImportType, registrator_id int64) []models.ProductsGroup {

	mainStruct := (*receiveStruct)
	var productGroups []models.ProductsGroup

	root := mainStruct.КоммерческаяИнформация.Классификатор.Свойства.Свойство
	// children, ok := findInStructRecursive(root, "ТоварнаяГруппа")
	//  if !ok {
	//  	return nil, errors.New("not found root nodes for product groups")
	//  }

	for i := 0; i < len(root); i++ {
		if root[i].Наименование == "ТоварнаяГруппа" {
			time := time.Now()

			for j := 0; j < len(root[i].ВариантыЗначений.Справочник); j++ {
				productGroup := models.ProductsGroup{
					Id_1c:          root[i].ВариантыЗначений.Справочник[j].ИдЗначения,
					Name_1c:        root[i].ВариантыЗначений.Справочник[j].Значение,
					Registrator_id: registrator_id,
					Changed_at:     &time,
				}
				productGroups = append(productGroups, productGroup)
			}
		}
	}
	return productGroups
}
