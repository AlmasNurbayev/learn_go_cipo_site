package partParsers

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"cipo_cite_server/internal/utils"
	"errors"
	"slices"
	"time"
)

// ищет в структуре вложенную структуру "Предложения" и возвращает ее элементы
func QntPrices(receiveStruct *XMLTypes.OfferType, registrator_id int64, operation_date time.Time,
	existsStores []models.Stores, existPrices []models.PriceVids, existSizes []models.Sizes,
	existProducts []models.Products) ([]models.QntPriceRegistry, error) {

	mainStruct := (*receiveStruct)
	var qntPrices = make([]models.QntPriceRegistry, 0, 3000)
	time := time.Now()

	root := mainStruct.КоммерческаяИнформация.ПакетПредложений.Предложения.Предложение

	for j := 0; j < len(root); j++ {

		// получаем ссылку на product
		productIndex := slices.IndexFunc(existProducts, func(item models.Products) bool {
			// так как в JSON ид задан дважды с разделителем, сначала вытаскиваем подстроку
			subString := utils.GetSubstringIfSymbolExists(root[j].Ид, "#")
			return item.Id_1c == subString
		})
		if productIndex == -1 {
			return nil, errors.New("Not found product in DB " + root[j].Ид)
		}
		product := existProducts[productIndex]

		// получаем ссылку на размер
		rootSv := root[j].ЗначенияСвойств.ЗначенияСвойства
		var size models.Sizes
		for k := 0; k < len(rootSv); k++ {
			if rootSv[k].Ид != "a001d8e3-a3b3-11ed-b0d2-50ebf624c538" {
				continue
			}
			sizeIndex := slices.IndexFunc(existSizes, func(item models.Sizes) bool {
				return item.Id_1c == rootSv[k].Значение
			})
			if sizeIndex == -1 {
				return nil, errors.New("not found size in DB " + rootSv[k].Ид)
			}
			size = existSizes[sizeIndex]
		}
		if size.Id == 0 {
			return nil, errors.New("not found size in DB")
		}

		// получаем ссылку на склад
		rootStore := root[j].Склад
		var store models.Stores
		var qnt float32
		for k := 0; k < len(rootStore); k++ {
			storeIndex := slices.IndexFunc(existsStores, func(item models.Stores) bool {
				return item.Id_1c == rootStore[k].ИдСклада
			})
			if storeIndex == -1 {
				return nil, errors.New("not found store in DB " + rootStore[k].ИдСклада)
			}
			store = existsStores[storeIndex]
			qnt = rootStore[k].КоличествоНаСкладе
		}
		if store.Id == 0 {
			return nil, errors.New("not found store in DB")
		}

		// получаем ссылку на вид цены и цену
		rootPriceVid := root[j].Цены.Цена
		var price float32
		var priceVid models.PriceVids
		for k := 0; k < len(rootPriceVid); k++ {
			priceIndex := slices.IndexFunc(existPrices, func(item models.PriceVids) bool {
				// если тип цены совпадает и он активен
				return item.Id_1c == rootPriceVid[k].ИдТипаЦены && item.Is_active
			})
			if priceIndex != -1 {
				priceVid = existPrices[priceIndex]
				price = rootPriceVid[k].ЦенаЗаЕдиницу
			}
		}
		if priceVid.Id == 0 {
			return nil, errors.New("not found price_vid in DB")
		}
		if price == 0 {
			return nil, errors.New("not found price in DB")
		}

		qntPrice := models.QntPriceRegistry{
			Product_name:       product.Name,
			Product_group_id:   product.Product_group_id,
			Vid_modeli_id:      *product.Vid_id,
			Product_created_at: &product.Created_at,
			Operation_date:     operation_date,
			Qnt:                qnt,
			Sum:                price,
			Store_id:           store.Id,
			Price_vid_id:       priceVid.Id,
			Size_id:            size.Id,
			Size_name_1c:       &size.Name_1c,
			Product_id:         product.Id,
			Registrator_id:     registrator_id,
			Changed_at:         &time,
		}
		//utils.PrintAsJSON(qntPrice)
		qntPrices = append(qntPrices, qntPrice)
	}
	return qntPrices, nil
}
