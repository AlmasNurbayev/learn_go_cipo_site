package parser

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	parser "cipo_cite_server/internal/parser/XMLtypes"
	"cipo_cite_server/internal/parser/partParsers"
	priceVids "cipo_cite_server/internal/repository/price_vids"
	"cipo_cite_server/internal/repository/products"
	"cipo_cite_server/internal/repository/qnt_price_registry"
	"cipo_cite_server/internal/repository/registrators"
	"cipo_cite_server/internal/repository/sizes"
	"cipo_cite_server/internal/repository/stores"
	"errors"
	"slices"
	"strconv"
)

func OfferParser(p *Parser, mainStruct *parser.OfferType, filePath string, newPath string) error {

	//utils.SaveStructToJSONFile(mainStruct, "offer.json", P.Log)
	p.Log.Info("Starting offer parsing")

	// получаем из XML и записываем регистратор
	registrator_id, err := parseAndSaveRegistratorOffer(mainStruct, filePath, newPath)
	if err != nil {
		p.Log.Error("Error parse or saving registrator:", err)
		return err
	}

	// получаем из XML и записываем размеры
	err = parseAndSaveSizes(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving sizes:", err)
		return err
	}

	// получаем из XML и записываем ТипыЦен
	err = parseAndSavePrices(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving prices:", err)
		return err
	}

	// получаем из XML и записываем Склады
	err = parseAndSaveStores(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving stores:", err)
		return err
	}

	err = parseAndSaveQntPrices(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving offers:", err)
		return err
	}

	return nil
}

func parseAndSaveRegistratorOffer(mainStruct *XMLTypes.OfferType,
	filePath string, newPath string) (int64, error) {
	registrator, err := partParsers.RegistratorParserFromOffer(mainStruct, filePath, newPath, P.Log)
	if err != nil {
		P.Log.Error("Error pasrsing registrator:", err)
		return 0, err
	}
	registerRepo := registrators.NewRepository(P.Tx)
	registrator_id, err := registerRepo.Create(*registrator)
	if err != nil {
		P.Log.Error("Error inserting registrators:", err)
		return 0, err
	}
	P.Log.Info("added registrator with id: " + strconv.Itoa(int(registrator_id)))
	return registrator_id, nil
}

func parseAndSaveSizes(mainStruct *XMLTypes.OfferType, registrator_id int64) error {
	NewSizes := partParsers.SizesParser(mainStruct, registrator_id)
	sizesRepo := sizes.NewRepository(P.Tx)

	// берем из базы имеющие записи и проверяем на дубликаты
	existSizes, err := sizesRepo.List()
	if err != nil {
		P.Log.Error("Error selecting sizes:", err)
		return err
	}
	P.Log.Info("exist sizes: " + strconv.Itoa(len(*existSizes)))
	for _, val := range NewSizes {
		indexDuplicated := slices.IndexFunc(*existSizes, func(item models.Sizes) bool {
			return item.Id_1c == val.Id_1c
		})
		if indexDuplicated != -1 {
			P.Log.Debug("Duplicated and updated sizes: " + val.Id_1c)
			val.Id = (*existSizes)[indexDuplicated].Id
			if _, err := sizesRepo.Update(val); err != nil {
				P.Log.Error("Error updating sizes:", err)
				return err
			}
			continue
		}
		if _, err := sizesRepo.Create(val); err != nil {
			P.Log.Error("Error inserting sizes:", err)
			return err
		}
	}
	return nil

}

func parseAndSavePrices(mainStruct *XMLTypes.OfferType, registrator_id int64) error {
	NewPrices := partParsers.PriceVidsParser(mainStruct, registrator_id)
	pricesRepo := priceVids.NewRepository(P.Tx)

	// берем из базы имеющие записи и проверяем на дубликаты
	existPrices, err := pricesRepo.List()
	if err != nil {
		P.Log.Error("Error selecting price_vids:", err)
		return err
	}
	P.Log.Info("exist price_vids: " + strconv.Itoa(len(*existPrices)))
	for _, val := range NewPrices {
		indexDuplicated := slices.IndexFunc(*existPrices, func(item models.PriceVids) bool {
			return item.Id_1c == val.Id_1c
		})
		if indexDuplicated != -1 {
			P.Log.Debug("Duplicated and updated price_vids: " + val.Id_1c)
			val.Id = (*existPrices)[indexDuplicated].Id
			if _, err := pricesRepo.Update(val); err != nil {
				P.Log.Error("Error updating price_vids:", err)
				return err
			}
			continue
		}
		if _, err := pricesRepo.Create(val); err != nil {
			P.Log.Error("Error inserting price_vids:", err)
			return err
		}
	}

	return nil

}

func parseAndSaveStores(mainStruct *XMLTypes.OfferType, registrator_id int64) error {
	NewStores := partParsers.StoresParser(mainStruct, registrator_id)
	storesRepo := stores.NewRepository(P.Tx)

	// берем из базы имеющие записи и проверяем на дубликаты
	existStores, err := storesRepo.List()
	if err != nil {
		P.Log.Error("Error selecting stores:", err)
		return err
	}
	P.Log.Info("exist stores: " + strconv.Itoa(len(*existStores)))
	for _, val := range NewStores {
		indexDuplicated := slices.IndexFunc(*existStores, func(item models.Stores) bool {
			return item.Id_1c == val.Id_1c
		})
		if indexDuplicated != -1 {
			P.Log.Debug("Duplicated and updated stores: " + val.Id_1c)
			val.Id = (*existStores)[indexDuplicated].Id
			if _, err := storesRepo.Update(val); err != nil {
				P.Log.Error("Error updating price_vids:", err)
				return err
			}
			continue
		}
		if _, err := storesRepo.Create(val); err != nil {
			P.Log.Error("Error inserting stores:", err)
			return err
		}
	}

	return nil

}

func parseAndSaveQntPrices(mainStruct *XMLTypes.OfferType, registrator_id int64) error {

	// получаем значения справочников из БД
	registratorRepo := registrators.NewRepository(P.Tx)
	existsRegistrator, err := registratorRepo.GetById(registrator_id)
	if err != nil {
		P.Log.Error("Error selecting registrators:", err)
		return err
	}
	if len(*existsRegistrator) == 0 {
		return errors.New("registrator not found by id " + strconv.Itoa(int(registrator_id)))
	}

	// получаем значения справочников из БД
	storesRepo := stores.NewRepository(P.Tx)
	existsStores, err := storesRepo.List()
	if err != nil {
		P.Log.Error("Error selecting stores:", err)
		return err
	}

	pricesRepo := priceVids.NewRepository(P.Tx)
	existPrices, err := pricesRepo.List()
	if err != nil {
		P.Log.Error("Error selecting price_vids:", err)
		return err
	}

	sizesRepo := sizes.NewRepository(P.Tx)
	existSizes, err := sizesRepo.List()
	if err != nil {
		P.Log.Error("Error selecting sizes:", err)
		return err
	}

	productsRepo := products.NewRepository(P.Tx)
	existProducts, err := productsRepo.List()
	if err != nil {
		P.Log.Error("Error selecting products:", err)
		return err
	}

	// получаем новые значения
	NewQntPrices, err := partParsers.QntPrices(mainStruct, registrator_id, (*existsRegistrator)[0].Operation_date, *existsStores, *existPrices, *existSizes, *existProducts)
	if err != nil {
		P.Log.Error("Error parsing qnt_prices:", err)
		return err
	}
	qntPricesRepo := qnt_price_registry.NewRepository(P.Tx)
	for _, val := range NewQntPrices {
		_, err := qntPricesRepo.Create(val)
		if err != nil {
			P.Log.Error("Error inserting qnt_price_registry:", err)
			return err
		}
	}
	P.Log.Info("qnt_prices add count: " + strconv.Itoa(len(NewQntPrices)))
	return nil

}
