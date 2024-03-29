package moved

import (
	"cipo_cite_server/internal/config"
	"log/slog"
	"os"
	"time"
)

// находим в папке Input файлы, создаем папку Webdata_дата_время,
// перемещаем туда offers0_1.xml и import0_1.xml
func MovedInputFiles(cfg config.Config, log *slog.Logger) (*MovedInputFilesT, error) {

	currentTime := time.Now()
	folderName := "input/webdata_" + currentTime.Format("2006_01_02_15_04_05")
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(folderName, 0755)
		if err != nil {
			log.Error("Error if create folder:", err)
			return nil, err
		}
	}

	filesName := []InputFilesT{
		{"classificator", cfg.Parser.Classificator_name},
		{"offer", cfg.Parser.Offer_name},
		{"imageFolder", cfg.Parser.ImageFolder_name},
	}

	for i, file := range filesName {
		oldPath := "input/" + file.PathFile
		newPath := folderName + "/" + file.PathFile
		filesName[i].PathFile = newPath
		if _, err := os.Stat(oldPath); err == nil {
			err := os.Rename(oldPath, newPath)
			if err != nil {
				if i == 2 {
					// если это imageFolder, то не прерываем программу
					log.Error("Error moving imageFolder:", err)
				} else {
					log.Error("Error moving file:", err)
					return nil, err
				}
			} else {
				log.Info(file.PathFile + " exists and moved successfully")
			}
		} else {
			if i == 2 {
				// если это imageFolder, то не прерываем программу
				log.Error(file.PathFile + " does not exist")
			} else {
				log.Error(file.PathFile + " does not exist")
				return nil, err
			}
		}
	}
	return &MovedInputFilesT{Files: filesName, NewPath: folderName}, nil
}

type InputFilesT struct {
	TypeFile string
	PathFile string
}

type MovedInputFilesT struct {
	Files   []InputFilesT
	NewPath string
}
