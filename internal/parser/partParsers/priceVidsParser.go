package partParsers

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"time"
)

// ищет в структуре вложенную структуру "ТипыЦен" и возвращает ее элементы
func PriceVidsParser(receiveStruct *XMLTypes.OfferType, registrator_id int64) []models.PriceVids {

	mainStruct := (*receiveStruct)
	var prices []models.PriceVids
	time := time.Now()

	root := mainStruct.КоммерческаяИнформация.ПакетПредложений.ТипыЦен.ТипЦены

	// TODO - в этой ветке стуктуры единственный массив, может быть не проверять Ид
	for j := 0; j < len(root); j++ {
		//var vid T
		priceVid := models.PriceVids{
			Id_1c:              root[j].Ид,
			Name_1c:            root[j].Наименование,
			Is_active:          true,
			Active_change_date: time,
			Registrator_id:     registrator_id,
			Changed_at:         &time,
		}
		prices = append(prices, priceVid)
	}
	return prices
}
