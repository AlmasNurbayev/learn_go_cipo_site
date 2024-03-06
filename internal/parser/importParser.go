package parser

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"cipo_cite_server/internal/parser/partParsers"
	"cipo_cite_server/internal/repository/products_groups"
	"cipo_cite_server/internal/repository/registrators"
	"cipo_cite_server/internal/repository/vids"
	"log/slog"
	"slices"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// читаем большую структуру из файла import.xml и записываем в базу
func ImportParser(p *Parser, mainStruct *XMLTypes.ImportType, filePath string, newPath string) {

	// получаем из XML и записываем регистратор
	registrator_id, err := parseAndSaveRegistrator(mainStruct, filePath, newPath, p.Sqlx, p.Log)
	if err != nil {
		p.Log.Error("Error parse or saving registrator:", err)
		return
	}

	// получаем из XML и записываем продуктовые группы
	err = parseAndSaveProductGroups(mainStruct, registrator_id, p.Sqlx, p.Log)
	if err != nil {
		p.Log.Error("Error parse or saving product_groups:", err)
		return
	}

	err = parseAndSaveVids(mainStruct, registrator_id, p.Sqlx, p.Log)
	if err != nil {
		p.Log.Error("Error parse or saving vids:", err)
		return
	}

}

func parseAndSaveRegistrator(mainStruct *XMLTypes.ImportType,
	filePath string, newPath string,
	db *sqlx.DB, log *slog.Logger) (int64, error) {
	registrator, err := partParsers.RegistratorParser(mainStruct, filePath, newPath, log)
	if err != nil {
		log.Error("Error pasrsing registrator:", err)
		return 0, err
	}
	registerRepo := registrators.NewRepository(db)
	registrator_id, err := registerRepo.Create(*registrator)
	if err != nil {
		log.Error("Error inserting registrators:", err)
		return 0, err
	}
	log.Info("added registrator with id: " + strconv.Itoa(int(registrator_id)))
	return registrator_id, nil
}

func parseAndSaveProductGroups(mainStruct *XMLTypes.ImportType,
	registrator_id int64, db *sqlx.DB,
	log *slog.Logger) error {
	NewProductGroups := partParsers.ProductGroupsParser(mainStruct, registrator_id)
	productGroupsRepo := products_groups.NewRepository(db)

	// берем из базы имеющие записи и проверяем на дубликаты
	existsProductGroups, err := productGroupsRepo.List()
	if err != nil {
		log.Error("Error selecting product_groups:", err)
		return err
	}
	log.Info("exist product_groups: " + strconv.Itoa(len(*existsProductGroups)))
	for _, val := range NewProductGroups {
		isDuplicated := slices.ContainsFunc(*existsProductGroups, func(item models.ProductsGroup) bool {
			return item.Id_1c == val.Id_1c
		})
		if isDuplicated {
			log.Warn("Duplicated product_groups: " + val.Id_1c)
			continue
		}
		if _, err := productGroupsRepo.Create(val); err != nil {
			log.Error("Error inserting product_groups:", err)
			return err
		}
	}
	return nil
}

func parseAndSaveVids(mainStruct *XMLTypes.ImportType,
	registrator_id int64, db *sqlx.DB,
	log *slog.Logger) error {
	NewVids := partParsers.VidParser(mainStruct, registrator_id)
	//utils.PrintAsJSON(NewVids)

	// берем из базы имеющие записи и проверяем на дубликаты
	vidsRepo := vids.NewRepository(db)
	existsVids, err := vidsRepo.List()
	if err != nil {
		log.Error("Error selecting vids:", err)
		return err
	}
	log.Info("exist vids: " + strconv.Itoa(len(*existsVids)))
	for _, val := range NewVids {
		isDuplicated := slices.ContainsFunc(*existsVids, func(item models.VidsModel) bool {
			return item.Id_1c == val.Id_1c
		})
		if isDuplicated {
			log.Warn("Duplicated vids: " + val.Id_1c)
			continue
		}
		if _, err := vidsRepo.Create(val); err != nil {
			log.Error("Error inserting vids:", err)
			return err
		}
	}
	return nil
}
