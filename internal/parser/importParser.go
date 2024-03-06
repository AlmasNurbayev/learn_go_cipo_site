package parser

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"cipo_cite_server/internal/parser/partParsers"
	"cipo_cite_server/internal/repository/product_desc"
	"cipo_cite_server/internal/repository/product_vids"
	"cipo_cite_server/internal/repository/products_groups"
	"cipo_cite_server/internal/repository/registrators"
	"cipo_cite_server/internal/repository/vids"
	"slices"
	"strconv"
)

var P *Parser

// читаем большую структуру из файла import.xml и записываем в базу
func ImportParser(p *Parser, mainStruct *XMLTypes.ImportType, filePath string, newPath string) {
	P = p
	// получаем из XML и записываем регистратор
	registrator_id, err := parseAndSaveRegistrator(mainStruct, filePath, newPath)
	if err != nil {
		p.Log.Error("Error parse or saving registrator:", err)
		return
	}

	// получаем из XML и записываем продуктовые группы
	err = parseAndSaveProductGroups(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving product_groups:", err)
		return
	}

	// получаем из XML и записываем продуктовые виды моделей
	err = parseAndSaveVids(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving vids:", err)
		return
	}

	// получаем из XML и записываем виды продуктов
	err = parseAndSaveProductVids(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving product_vids:", err)
		return
	}

	// получаем из БД доп. реквизиты продуктов
	productDescRepo := product_desc.NewRepository(P.Sqlx)
	existsProductDesc, err := productDescRepo.List()
	if err != nil {
		p.Log.Error("Error selecting product_desc:", err)
		return
	}
	P.Log.Info("exist product_desc: " + strconv.Itoa(len(*existsProductDesc)))

}

func parseAndSaveRegistrator(mainStruct *XMLTypes.ImportType,
	filePath string, newPath string) (int64, error) {
	registrator, err := partParsers.RegistratorParser(mainStruct, filePath, newPath, P.Log)
	if err != nil {
		P.Log.Error("Error pasrsing registrator:", err)
		return 0, err
	}
	registerRepo := registrators.NewRepository(P.Sqlx)
	registrator_id, err := registerRepo.Create(*registrator)
	if err != nil {
		P.Log.Error("Error inserting registrators:", err)
		return 0, err
	}
	P.Log.Info("added registrator with id: " + strconv.Itoa(int(registrator_id)))
	return registrator_id, nil
}

func parseAndSaveProductGroups(mainStruct *XMLTypes.ImportType,
	registrator_id int64) error {
	NewProductGroups := partParsers.ProductGroupsParser(mainStruct, registrator_id)
	productGroupsRepo := products_groups.NewRepository(P.Sqlx)

	// берем из базы имеющие записи и проверяем на дубликаты
	existsProductGroups, err := productGroupsRepo.List()
	if err != nil {
		P.Log.Error("Error selecting product_groups:", err)
		return err
	}
	P.Log.Info("exist product_groups: " + strconv.Itoa(len(*existsProductGroups)))
	for _, val := range NewProductGroups {
		indexDuplicated := slices.IndexFunc(*existsProductGroups, func(item models.ProductsGroup) bool {
			return item.Id_1c == val.Id_1c
		})
		if indexDuplicated != -1 {
			P.Log.Warn("Duplicated and updated product_groups: " + val.Id_1c)
			val.Id = (*existsProductGroups)[indexDuplicated].Id
			if _, err := productGroupsRepo.Update(val); err != nil {
				P.Log.Error("Error updating vids:", err)
				return err
			}
			continue
		}
		if _, err := productGroupsRepo.Create(val); err != nil {
			P.Log.Error("Error inserting product_groups:", err)
			return err
		}
	}
	return nil
}

func parseAndSaveVids(mainStruct *XMLTypes.ImportType,
	registrator_id int64) error {
	NewVids := partParsers.VidParser(mainStruct, registrator_id)
	//utils.PrintAsJSON(NewVids)

	// берем из базы имеющие записи и проверяем на дубликаты
	vidsRepo := vids.NewRepository(P.Sqlx)
	existsVids, err := vidsRepo.List()
	if err != nil {
		P.Log.Error("Error selecting vids:", err)
		return err
	}
	P.Log.Info("exist vids: " + strconv.Itoa(len(*existsVids)))
	for _, val := range NewVids {
		indexDuplicated := slices.IndexFunc(*existsVids, func(item models.VidsModel) bool {
			return item.Id_1c == val.Id_1c
		})
		if indexDuplicated != -1 {
			P.Log.Warn("Duplicated and updated vids: " + val.Id_1c)
			val.Id = (*existsVids)[indexDuplicated].Id
			if _, err := vidsRepo.Update(val); err != nil {
				P.Log.Error("Error updating vids:", err)
				return err
			}
			continue
		}
		if _, err := vidsRepo.Create(val); err != nil {
			P.Log.Error("Error inserting vids:", err)
			return err
		}
	}
	return nil
}

func parseAndSaveProductVids(mainStruct *XMLTypes.ImportType,
	registrator_id int64) error {
	NewProductsVids := partParsers.ProductVidsParser(mainStruct, registrator_id)
	//utils.PrintAsJSON(NewVids)

	// берем из базы имеющие записи и проверяем на дубликаты
	productVidRepo := product_vids.NewRepository(P.Sqlx)
	existsProductsVids, err := productVidRepo.List()
	if err != nil {
		P.Log.Error("Error selecting product_vids:", err)
		return err
	}
	P.Log.Info("exist product_vids: " + strconv.Itoa(len(*existsProductsVids)))
	for _, val := range NewProductsVids {
		indexDuplicated := slices.IndexFunc(*existsProductsVids, func(item models.ProductVids) bool {
			return item.Id_1c == val.Id_1c
		})
		if indexDuplicated != -1 {
			P.Log.Warn("Duplicated and updated product_vids: " + val.Id_1c)
			val.Id = (*existsProductsVids)[indexDuplicated].Id
			if _, err := productVidRepo.Update(val); err != nil {
				P.Log.Error("Error updating product_vids:", err)
				return err
			}

			continue
		}
		if _, err := productVidRepo.Create(val); err != nil {
			P.Log.Error("Error inserting vids:", err)
			return err
		}
	}
	return nil
}
