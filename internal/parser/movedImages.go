package parser

import (
	"cipo_cite_server/internal/utils"
	"log/slog"
	"os"
)

// перемещаем папку "import_files" в другую папку для загруженного пакета
func MovedImages(newPath string, log *slog.Logger) error {
	utils.PrintAsJSON(newPath)
	oldPath := "input/import_files"
	newPath = newPath + "/import_files"
	if _, err := os.Stat(oldPath); err == nil {
		err := os.Rename(oldPath, newPath)
		if err != nil {
			log.Error("Error moving images:", err)
			return err
		} else {
			log.Info(oldPath + " exists and moved successfully to " + newPath)
		}
	} else {
		log.Error(oldPath + " does not exist")
		return err
	}
	return nil
}
