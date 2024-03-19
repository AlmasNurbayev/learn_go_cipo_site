package parser

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"cipo_cite_server/internal/parser/partParsers"
	"cipo_cite_server/internal/repository/image_registry"
	"cipo_cite_server/internal/repository/product_desc"
	"cipo_cite_server/internal/repository/product_vids"
	"cipo_cite_server/internal/repository/products"
	"cipo_cite_server/internal/repository/products_groups"
	"cipo_cite_server/internal/repository/registrators"
	"cipo_cite_server/internal/repository/vids"
	"cipo_cite_server/internal/utils"
	"fmt"
	"slices"
	"sort"
	"strconv"
)

var P *Parser

// читаем большую структуру из файла import.xml и записываем в базу
func ImportParser(p *Parser, mainStruct *XMLTypes.ImportType, filePath string, newPath string) error {
	P = p
	p.Log.Info("Starting import parsing")

	// получаем из XML и записываем регистратор
	registrator_id, err := parseAndSaveRegistrator(mainStruct, filePath, newPath)
	if err != nil {
		p.Log.Error("Error parse or saving registrator:", err)
		return err
	}

	// получаем из XML и записываем продуктовые группы
	err = parseAndSaveProductGroups(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving product_groups:", err)
		return err
	}

	// получаем из XML и записываем продуктовые виды моделей
	err = parseAndSaveVids(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving vids:", err)
		return err
	}

	// получаем из XML и записываем виды продуктов
	err = parseAndSaveProductVids(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving product_vids:", err)
		return err
	}

	// получаем из БД доп. реквизиты продуктов
	productDescRepo := product_desc.NewRepository(P.Tx)
	existsProductDesc, err := productDescRepo.List()
	if err != nil {
		p.Log.Error("Error selecting product_desc:", err)
		return err
	}
	P.Log.Info("exist product_desc: " + strconv.Itoa(len(*existsProductDesc)))

	// получаем из XML и записываем продукты
	err = parseAndSaveProducts(mainStruct, registrator_id)
	if err != nil {
		p.Log.Error("Error parse or saving products:", err)
		return err
	}

	// получаем из XML и записываем картинки
	err = parseAndSaveImages(mainStruct, registrator_id, newPath)
	if err != nil {
		p.Log.Error("Error parse or saving images:", err)
		return err
	}
	return nil
}

func parseAndSaveRegistrator(mainStruct *XMLTypes.ImportType,
	filePath string, newPath string) (int64, error) {
	registrator, err := partParsers.RegistratorParserFromImport(mainStruct, filePath, newPath, P.Log)
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

func parseAndSaveProductGroups(mainStruct *XMLTypes.ImportType,
	registrator_id int64) error {
	NewProductGroups := partParsers.ProductGroupsParser(mainStruct, registrator_id)
	productGroupsRepo := products_groups.NewRepository(P.Tx)

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
			P.Log.Debug("Duplicated and updated product_groups: " + val.Id_1c)
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
	vidsRepo := vids.NewRepository(P.Tx)
	existsVids, err := vidsRepo.List()
	if err != nil {
		P.Log.Error("Error selecting vids:", err)
		return err
	}
	P.Log.Info("exist vids: " + strconv.Itoa(len(*existsVids)))
	for _, val := range NewVids {
		indexDuplicated := slices.IndexFunc(*existsVids, func(item models.Vids) bool {
			return item.Id_1c == val.Id_1c
		})
		if indexDuplicated != -1 {
			P.Log.Debug("Duplicated and updated vids: " + val.Id_1c)
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
	productVidRepo := product_vids.NewRepository(P.Tx)
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
			P.Log.Debug("Duplicated and updated product_vids: " + val.Id_1c)
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

func parseAndSaveProducts(mainStruct *XMLTypes.ImportType, registrator_id int64) error {

	// считываем и передаем все значения таблицы product_group
	productGroupsRepo := products_groups.NewRepository(P.Tx)
	existsProductGroups, err := productGroupsRepo.List()
	if err != nil {
		P.Log.Error("Error selecting product_groups:", err)
		return err
	}

	// считываем и передаем все значения таблицы product_vid
	productVidRepo := product_vids.NewRepository(P.Tx)
	existsProductsVids, err := productVidRepo.List()
	if err != nil {
		P.Log.Error("Error selecting product_vids:", err)
		return err
	}

	// считываем и передаем все значения таблицы product_desc
	productDescRepo := product_desc.NewRepository(P.Tx)
	existsProductDesc, err := productDescRepo.List()
	if err != nil {
		P.Log.Error("Error selecting product_desc:", err)
		return err
	}

	// считываем и передаем все значения таблицы vids
	vidsRepo := vids.NewRepository(P.Tx)
	existsVids, err := vidsRepo.List()
	if err != nil {
		P.Log.Error("Error selecting vids:", err)
		return err
	}

	// считываем все уже имющиеся записи в products
	productsRepo := products.NewRepository(P.Tx)
	existsProducts, err := productsRepo.List()
	if err != nil {
		P.Log.Error("Error selecting products: ", err)
		return err
	}

	// парсим все продукты из XML
	NewProducts := partParsers.ProductsParser(mainStruct, registrator_id,
		*existsProductGroups, *existsProductsVids, *existsProductDesc, *existsVids)

	// сортируем по первым 20 символам поля id_1c, для хронологии
	sort.Slice(NewProducts, func(i, j int) bool {
		return NewProducts[i].Id_1c[19:] < NewProducts[j].Id_1c[19:]
	})

	fmt.Println("products parsing count: ", len(NewProducts))
	fmt.Println("products exists count: ", len(*existsProducts))

	for _, val := range NewProducts {
		indexDuplicated := slices.IndexFunc(*existsProducts, func(item models.Products) bool {
			return item.Id_1c == val.Id_1c
		})
		if indexDuplicated != -1 {
			//P.Log.Debug("Duplicated and updated products: " + val.Id_1c)
			val.Id = (*existsProducts)[indexDuplicated].Id
			if _, err := productsRepo.Update(val); err != nil {
				P.Log.Error("Error updating products:", err)
				return err
			}
			continue
		}
		if _, err := productsRepo.Create(val); err != nil {
			//utils.PrintAsJSON(val)
			P.Log.Error("Error inserting products: ", err)
			return err
		}
	}

	return nil
}

func parseAndSaveImages(mainStruct *XMLTypes.ImportType, registrator_id int64, newPath string) error {
	productsRepo := products.NewRepository(P.Tx)
	existsProducts, err := productsRepo.List()
	if err != nil {
		P.Log.Error("Error selecting products: ", err)
		return err
	}
	NewImages, err := partParsers.ImageParser(mainStruct, registrator_id, *existsProducts, newPath)
	if err != nil {
		P.Log.Error("Error parsing images: ", err)
		return err
	}
	fmt.Println("image parsing count: ", len(NewImages))

	// считываем все уже имющиеся записи в products
	imagesRepo := image_registry.NewRepository(P.Tx)
	existsImages, err := imagesRepo.List()
	if err != nil {
		P.Log.Error("Error selecting images: ", err)
		return err
	}

	fmt.Println("image exists count: ", len(*existsImages))

	for _, val := range NewImages {
		indexDuplicated := slices.IndexFunc(*existsImages, func(item models.ImageRegistry) bool {
			return item.Name == val.Name
		})
		if indexDuplicated != -1 {
			//P.Log.Debug("Duplicated and updated products: " + val.Id_1c)
			if _, err := imagesRepo.Update(val); err != nil {
				P.Log.Error("Error updating images:", err)
				return err
			}
			continue
		}
		if _, err := imagesRepo.Create(val); err != nil {
			utils.PrintAsJSON(val)
			P.Log.Error("Error inserting images: ", err)
			return err
		}
	}

	return nil
}
