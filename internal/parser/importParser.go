package parser

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"cipo_cite_server/internal/parser/partParsers"
	"cipo_cite_server/internal/repository/products_groups"
	"cipo_cite_server/internal/repository/registrators"
	"os"
	"slices"
	"strconv"
)

// читаем большую структуру из файла import.xml и записываем в базу
func ImportParser(p *Parser, mainStruct *XMLTypes.ImportType, filePath string, newPath string) {
	registrator, err := partParsers.RegistratorParser(mainStruct, filePath, newPath, p.Log)
	if err != nil {
		p.Log.Error("Error pasrsing registrator:", err)
	}
	registerRepo := registrators.NewRepository(p.Sqlx)
	registrator_id, err := registerRepo.Create(*registrator)
	if err != nil {
		p.Log.Error("Error inserting registrators:", err)
		os.Exit(1)
	}
	p.Log.Info("added registrator with id: " + strconv.Itoa(int(registrator_id)))
	product_groups := partParsers.ProductGroupsParser(mainStruct, registrator_id)
	productGroupsRepo := products_groups.NewRepository(p.Sqlx)

	existsProductGroups, err := productGroupsRepo.List()
	if err != nil {
		p.Log.Error("Error selecting product_groups:", err)
		return
	}
	p.Log.Info("exist product_groups: " + strconv.Itoa(len(*existsProductGroups)))

	for _, val := range product_groups {
		isDuplicated := slices.ContainsFunc(*existsProductGroups, func(item models.ProductsGroup) bool {
			return item.Id_1c == val.Id_1c
		})
		if isDuplicated {
			p.Log.Warn("Duplicated product_groups: " + val.Id_1c)
			continue
		}
		if _, err := productGroupsRepo.Create(val); err != nil {
			p.Log.Error("Error inserting product_groups:", err)
		}
	}

}
