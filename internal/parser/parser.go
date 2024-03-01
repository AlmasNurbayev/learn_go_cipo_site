package parser

import (
	parser "cipo_cite_server/internal/parser/XMLtypes"
	"cipo_cite_server/internal/utils"
	"encoding/xml"
	"log/slog"
	"os"
)

func Parser(files *[]InputFilesT, log *slog.Logger) {
	utils.PrintAsJSON(files)

	for _, fileItem := range *files {
		file, err := os.Open(fileItem.PathFile)
		if err != nil {
			log.Error("Ошибка открытия файла:", err)
			return
		}
		defer file.Close()
		if fileItem.TypeFile == "classificator" {
			var temp parser.ImportType
			xmlStruct := temp.КоммерческаяИнформация
			decoder := xml.NewDecoder(file)
			if err := decoder.Decode(&xmlStruct); err != nil {
				log.Error("Ошибка декодирования XML:", err)
				return
			}
			utils.PrintAsJSON(xmlStruct)
		}
	}

}
