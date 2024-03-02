package parser

import (
	parser "cipo_cite_server/internal/parser/XMLtypes"
	"cipo_cite_server/internal/parser/importParser"
	"cipo_cite_server/internal/parser/offerParser"
	"cipo_cite_server/internal/utils"
	"encoding/xml"
	"log/slog"
	"os"
)

func ReadAndParse(files *[]InputFilesT, log *slog.Logger) {
	utils.PrintAsJSON(files)

	for _, fileItem := range *files {
		file, err := os.Open(fileItem.PathFile)
		if err != nil {
			log.Error("Ошибка открытия файла:", err)
			return
		}
		defer file.Close()
		switch fileItem.TypeFile {
		case "classificator":
			var temp parser.ImportType               // создаем экземпляр структуры
			xmlStruct := temp.КоммерческаяИнформация // создаем подчиненный узел
			decoder := xml.NewDecoder(file)
			if err := decoder.Decode(&xmlStruct); err != nil {
				log.Error("Ошибка декодирования XML:", err)
				return
			}
			// передаем полный тип, чтобы не выделять подчиненный узел в парсере
			importParser.Parser(&parser.ImportType{КоммерческаяИнформация: xmlStruct}, log)
		case "offer":
			var temp parser.OfferType
			xmlStruct := temp.КоммерческаяИнформация
			decoder := xml.NewDecoder(file)
			if err := decoder.Decode(&xmlStruct); err != nil {
				log.Error("Ошибка декодирования XML:", err)
				return
			}
			// передаем полный тип, чтобы не выделять подчиненный узел в парсере
			offerParser.Parser(&parser.OfferType{КоммерческаяИнформация: xmlStruct}, log)
		}
	}
}
