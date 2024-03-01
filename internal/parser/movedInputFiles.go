package parser

import (
	"cipo_cite_server/internal/config"
	"log/slog"
	"os"
	"time"
)

// находим в папке Input файлы, создаем папку Webdata_дата_время,
// перемещаем туда offers0_1.xml и import0_1.xml
func MovedInputFiles(cfg config.Config, log *slog.Logger) []moovedInputFiles {

	currentTime := time.Now()
	folderName := "input/webdata_" + currentTime.Format("2006_01_02_15_04_05")
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(folderName, 0755)
		if err != nil {
			log.Error("Error if create folder:", err)
			os.Exit(1)
		}
	}

	filesName := []moovedInputFiles{
		{"classificator", cfg.Parser.Classificator_name},
		{"offer", cfg.Parser.Offer_name},
	}

	for _, file := range filesName {
		oldPath := "input/" + file.pathFile
		newPath := folderName + "/" + file.pathFile
		file.pathFile = newPath
		if _, err := os.Stat(oldPath); err == nil {
			err := os.Rename(oldPath, newPath)
			if err != nil {
				log.Error("Error moving file:", err)
			} else {
				log.Debug(file.pathFile + " exists and moved successfully")
			}
		} else {
			log.Error(file.pathFile + " does not exist")
			os.Exit(1)
		}
	}
	return filesName
}

type moovedInputFiles struct {
	typeFile string
	pathFile string
}
