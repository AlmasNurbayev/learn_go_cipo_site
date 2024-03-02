package importParser

import (
	parser "cipo_cite_server/internal/parser/XMLtypes"
	"encoding/json"
	"log/slog"
	"os"
)

func Parser(mainStruct *parser.ImportType, log *slog.Logger) {
	// parser.Parser()
	//utils.PrintAsJSON((*mainStruct))

	jsonData, err := json.MarshalIndent((*mainStruct), "", "  ")
	if err != nil {
		log.Error("Ошибка маршалинга в JSON:", err)
		return
	}

	// Запись JSON данных в файл
	err = os.WriteFile("import.json", jsonData, 0755)
	if err != nil {
		log.Error("Ошибка записи в файл:", err)
		return
	}

	log.Debug("Структура успешно сохранена в файл 'import.json'")
}
