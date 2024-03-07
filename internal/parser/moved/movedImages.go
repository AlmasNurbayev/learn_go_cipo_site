package moved

import (
	"log/slog"
	"os"

	cp "github.com/otiai10/copy"
)

// перемещаем папку "import_files" в другую папку для загруженного пакета
func CopyImages(assetsFolder string, newPath string, log *slog.Logger) error {
	//utils.PrintAsJSON(newPath)
	oldPath := newPath + "/import_files"
	newPath = "assets/product_images"
	if _, err := os.Stat(oldPath); err == nil {
		err := cp.Copy(oldPath, newPath)
		if err != nil {
			log.Error("Error copying images:", err)
			return err
		} else {
			log.Info(oldPath + " exists and copied successfully to " + newPath)
		}
	} else {
		log.Error(oldPath + " does not exist")
		return err
	}
	return nil
}
