package parser

import (
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"cipo_cite_server/internal/parser/partParsers"
	"cipo_cite_server/internal/storage/operations"
	"cipo_cite_server/internal/utils"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

// читаем большую структуру из файла import.xml и записываем в базу
func ImportParser(log *slog.Logger, db *sqlx.DB, mainStruct *XMLTypes.ImportType, filePath string, newPath string) {
	registrator, err := partParsers.RegistratorParser(mainStruct, filePath, newPath, log)
	if err != nil {
		log.Error("Error pasrsing registrator:", err)
	}
	registrator_id, err := operations.Insert(
		db, log, "registrators",
		registrator,
		[]string{"Id", "Created_at"})

	if err != nil {
		log.Error("Error inserting registrators:", err)
		return
	}

	product_groups := partParsers.ProductGroupsParser(mainStruct, registrator_id)
	utils.PrintAsJSON(product_groups)
	for _, val := range product_groups {
		if _, err := operations.Insert(
			db, log, "product_groups",
			val,
			[]string{"Id", "Created_at"}); err != nil {
			log.Error("Error inserting product_groups:", err)
			return
		}
	}
	//utils.SaveStructToJSONFile(*mainStruct, "import.json", log)
}
