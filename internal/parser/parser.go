package parser

import (
	XMLtypes "cipo_cite_server/internal/parser/XMLtypes"
	"cipo_cite_server/internal/utils"
	"encoding/xml"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
)

// читаем файлы и обмена и записываем в базу
func ReadAndParse(moved *MovedInputFilesT, log *slog.Logger, db *sqlx.DB) error {
	utils.PrintAsJSON(moved.Files)

	for _, fileItem := range moved.Files {
		file, err := os.Open(fileItem.PathFile)
		if err != nil {
			log.Error("Ошибка открытия файла:", err)
			return err
		}
		defer file.Close()
		switch fileItem.TypeFile {
		case "classificator":
			var temp XMLtypes.ImportType             // создаем экземпляр структуры
			xmlStruct := temp.КоммерческаяИнформация // создаем подчиненный узел для декодирования
			decoder := xml.NewDecoder(file)
			if err := decoder.Decode(&xmlStruct); err != nil {
				log.Error("Ошибка декодирования XML:", err)
				return err
			}
			// передаем полный тип, чтобы не выделять подчиненный узел в парсере
			utils.PrintAsJSON(moved)
			ImportParser(log, db, &XMLtypes.ImportType{КоммерческаяИнформация: xmlStruct}, fileItem.PathFile, moved.NewPath)
		case "offer":
			var temp XMLtypes.OfferType
			xmlStruct := temp.КоммерческаяИнформация
			decoder := xml.NewDecoder(file)
			if err := decoder.Decode(&xmlStruct); err != nil {
				log.Error("Ошибка декодирования XML:", err)
				return err
			}
			// передаем полный тип, чтобы не выделять подчиненный узел в парсере
			OfferParser(&XMLtypes.OfferType{КоммерческаяИнформация: xmlStruct}, log)
		}
	}
	return nil
}
