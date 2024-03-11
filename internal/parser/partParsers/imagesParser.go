package partParsers

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

// ищет в структуре товары, их узлы картинок и возвращат слайс картинок
func ImageParser(receiveStruct *XMLTypes.ImportType,
	registrator_id int64,
	products []models.Products, newPath string,
) ([]models.ImageRegistry, error) {

	mainStruct := (*receiveStruct)
	time := time.Now()
	var images []models.ImageRegistry

	root := mainStruct.КоммерческаяИнформация.Каталог.Товары.Товар

	for productIndex := 0; productIndex < len(root); productIndex++ {
		if root[productIndex].Ид == "" {
			continue
		}
		is_exists_product := slices.IndexFunc(products, func(item models.Products) bool {
			return item.Id_1c == root[productIndex].Ид
		})
		if is_exists_product == -1 {
			return nil, errors.New("in XML found product not exists in DB " + root[productIndex].Ид)
		}
		// if root[productIndex].Наименование == "Cipo туфли черный-лаковый 6700-01" {
		// 	utils.PrintAsJSON(root[productIndex])
		// }
		countImagesInProduct := 0

		root_images := root[productIndex].Картинка
		for imageIndex := 0; imageIndex < len(root_images); imageIndex++ {
			// TODO - for moved
			full_name := strings.Replace(root_images[imageIndex], "import_files", "product_images", -1)
			//full_name := root_images[imageIndex]
			// TODO - for moved
			fileInfo, err := os.Stat("assets/" + full_name)
			//fileInfo, err := os.Stat(newPath + "/" + full_name)
			if err != nil {
				return nil, errors.New("Error getting file information: " + err.Error())
			}
			var is_main bool
			if countImagesInProduct == 0 {
				is_main = true
			} else {
				is_main = false
			}
			image := models.ImageRegistry{
				Full_name: full_name,
				Name:      filepath.Base(full_name),
				Path:      filepath.Dir(full_name),
				Size:      int(fileInfo.Size()),

				Is_active:      true,
				Is_main:        is_main,
				Registrator_id: registrator_id,
				Product_id:     products[is_exists_product].Id,

				Operation_date: time,
				Changed_at:     &time,
			}
			images = append(images, image)
			countImagesInProduct = countImagesInProduct + 1
			//utils.PrintAsJSON(image)
		}

		//fmt.Println(root[productIndex].Наименование)

	}
	return images, nil
}
