package partParsers

import (
	"cipo_cite_server/internal/config"
	"cipo_cite_server/internal/models"
	XMLTypes "cipo_cite_server/internal/parser/XMLtypes"
	"log/slog"
	"time"
)

// ищет в структуре вложенную структуру "Классификатор", "Каталог" и возвращат ее поля
// а также сведения о загружаемом файле
func RegistratorParserFromImport(receiveStruct *XMLTypes.ImportType, filePath string, newPath string, log *slog.Logger) (*models.RegistratorsModel, error) {
	// parser.Parser()
	Cfg := config.MustLoad()

	mainStruct := (*receiveStruct)
	var registratorStruct models.RegistratorsModel

	registratorStruct.Name_folder = newPath
	registratorStruct.Name_file = filePath
	registratorStruct.User_id = Cfg.Config.Parser.Default_user_id
	registratorStruct.Id_catalog = mainStruct.КоммерческаяИнформация.Каталог.Ид
	registratorStruct.Id_class = mainStruct.КоммерческаяИнформация.Классификатор.Ид
	registratorStruct.Name_catalog = mainStruct.КоммерческаяИнформация.Каталог.Наименование
	registratorStruct.Name_class = mainStruct.КоммерческаяИнформация.Классификатор.Наименование
	registratorStruct.Operation_date = time.Now()
	registratorStruct.Changed_at = time.Now()
	registratorStruct.Ver_schema = mainStruct.КоммерческаяИнформация.ВерсияСхемы
	if mainStruct.КоммерческаяИнформация.Каталог.СодержитТолькоИзменения == "false" {
		registratorStruct.Is_only_change = false
	} else {
		registratorStruct.Is_only_change = true
	}

	layout := "2006-01-02T15:04:05"
	var time, err = time.Parse(layout, mainStruct.КоммерческаяИнформация.ДатаФормирования)
	if err != nil {
		log.Error("Error parsing time:", err)
		return nil, err
	}
	registratorStruct.Date_schema = time
	//utils.PrintAsJSON(registratorStruct)

	return &registratorStruct, nil
}

func RegistratorParserFromOffer(receiveStruct *XMLTypes.OfferType, filePath string, newPath string, log *slog.Logger) (*models.RegistratorsModel, error) {
	// parser.Parser()
	Cfg := config.MustLoad()

	mainStruct := (*receiveStruct)
	var registratorStruct models.RegistratorsModel

	registratorStruct.Name_folder = newPath
	registratorStruct.Name_file = filePath
	registratorStruct.User_id = Cfg.Config.Parser.Default_user_id
	registratorStruct.Id_catalog = mainStruct.КоммерческаяИнформация.Классификатор.Наименование
	registratorStruct.Id_class = mainStruct.КоммерческаяИнформация.Классификатор.Ид
	registratorStruct.Name_catalog = mainStruct.КоммерческаяИнформация.ПакетПредложений.Наименование
	registratorStruct.Name_class = mainStruct.КоммерческаяИнформация.ПакетПредложений.Ид
	registratorStruct.Operation_date = time.Now()
	registratorStruct.Changed_at = time.Now()
	registratorStruct.Ver_schema = mainStruct.КоммерческаяИнформация.ВерсияСхемы
	if mainStruct.КоммерческаяИнформация.ПакетПредложений.СодержитТолькоИзменения == "false" {
		registratorStruct.Is_only_change = false
	} else {
		registratorStruct.Is_only_change = true
	}

	layout := "2006-01-02T15:04:05"
	var time, err = time.Parse(layout, mainStruct.КоммерческаяИнформация.ДатаФормирования)
	if err != nil {
		log.Error("Error parsing time:", err)
		return nil, err
	}
	registratorStruct.Date_schema = time
	//utils.PrintAsJSON(registratorStruct)

	return &registratorStruct, nil
}
