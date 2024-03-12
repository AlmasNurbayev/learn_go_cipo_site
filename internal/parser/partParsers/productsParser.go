package partParsers

import (
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"slices"
	"strconv"
	"time"
)

// ищет в структуре Товары  и возвращат ее элементы
func ProductsParser(
	receiveStruct *XMLTypes.ImportType,
	registrator_id int64,
	existsProductGroups []models.ProductsGroup,
	existsProductsVids []models.ProductVids,
	existsProductDesc []models.ProductsDesc,
	existsVid []models.Vids,
) []models.Products {

	mainStruct := (*receiveStruct)
	var products []models.Products

	root := mainStruct.КоммерческаяИнформация.Каталог.Товары.Товар

	for i := 0; i < len(root); i++ {

		var product_folder string = ""
		var product_vid_id *int64
		var description *string
		var product_desc struct {
			Material_inside   *string
			Material_podoshva *string
			Material_up       *string
			Sex               *int
			Main_color        *string
			Is_public_web     bool
			Product_group_id  int64
			Vid_id            *int64
		}
		var base_ed string
		time := time.Now()

		if root[i].Группы.Ид != "" {
			product_folder = root[i].Группы.Ид
		}
		if root[i].БазоваяЕдиница.НаименованиеПолное != "" {
			base_ed = root[i].БазоваяЕдиница.НаименованиеПолное
		}

		if root[i].Описание != "" {
			description = &root[i].Описание
		}

		root_rekv := root[i].ЗначенияРеквизитов.ЗначениеРеквизита
		for j := 0; j < len(root_rekv); j++ {
			if root_rekv[j].Наименование == "ВидНоменклатуры" {
				product_vid_index := slices.IndexFunc(existsProductsVids, func(item models.ProductVids) bool {
					return item.Name_1c == root_rekv[j].Значение
				})
				if product_vid_index != -1 {
					product_vid_id = &existsProductsVids[product_vid_index].Id
				}
			}
		}

		root_svoistv := root[i].ЗначенияСвойств.ЗначенияСвойства

		// разбираем перечень свойств сопоставляя со справочником Product_desc
		for k := 0; k < len(root_svoistv); k++ {
			for _, val := range existsProductDesc {
				if val.Id_1c == root_svoistv[k].Ид {
					if val.Field == "material_podoshva" {
						product_desc.Material_podoshva = &root_svoistv[k].Значение
					}
					if val.Field == "material_inside" {
						product_desc.Material_inside = &root_svoistv[k].Значение
					}
					if val.Field == "material_up" {
						product_desc.Material_up = &root_svoistv[k].Значение
					}
					if val.Field == "main_color" {
						product_desc.Main_color = &root_svoistv[k].Значение
					}
					if val.Field == "public_web" {
						if root_svoistv[k].Значение == "Да" {
							product_desc.Is_public_web = true
						} else {
							product_desc.Is_public_web = false
						}
					}
					if val.Field == "sex" {
						intSex, err := strconv.Atoi(root_svoistv[k].Значение)
						product_desc.Sex = &intSex
						if err != nil {
							product_desc.Sex = nil
						}
					}
					if val.Field == "product_group" {
						// в справочнике Product_desc ищем какое id_1c имеет свойство "ТоварнаяГруппа"
						product_group_index := slices.IndexFunc(existsProductGroups, func(item models.ProductsGroup) bool {
							return item.Id_1c == root_svoistv[k].Значение
						})
						if product_group_index != -1 {
							product_desc.Product_group_id = existsProductGroups[product_group_index].Id
							//product_desc.Product_group_id.Valid = true
						}
					}
					if val.Field == "vids" {
						// в справочнике Product_desc ищем какое id_1c имеет свойство "Виды"
						vid_index := slices.IndexFunc(existsVid, func(item models.Vids) bool {
							return item.Id_1c == root_svoistv[k].Значение
						})
						if vid_index != -1 {
							product_desc.Vid_id = &existsVid[vid_index].Id
						}
					}
				}
			}
		}

		newProduct := models.Products{
			Id_1c:             root[i].Ид,
			Name_1c:           root[i].Наименование,
			Name:              root[i].Наименование,
			Registrator_id:    registrator_id,
			Artikul:           root[i].Артикул,
			Description:       description,
			Changed_at:        &time,
			Base_ed:           base_ed,
			Product_group_id:  product_desc.Product_group_id,
			Product_folder:    product_folder,
			Product_vid_id:    product_vid_id,
			Material_podoshva: product_desc.Material_podoshva,
			Material_inside:   product_desc.Material_inside,
			Material_up:       product_desc.Material_up,
			Sex:               product_desc.Sex,
			Main_color:        product_desc.Main_color,
			Is_public_web:     product_desc.Is_public_web,
			Vid_id:            product_desc.Vid_id,
		}
		products = append(products, newProduct)
	}
	return products
}
