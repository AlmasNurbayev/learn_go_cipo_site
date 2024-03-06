package parser

import (
	"cipo_cite_server/internal/config"
	"cipo_cite_server/internal/logger"
	XMLtypes "cipo_cite_server/internal/parser/XMLtypes"
	"cipo_cite_server/internal/parser/moved"
	"cipo_cite_server/internal/storage/postgres"
	"cipo_cite_server/internal/utils"
	"encoding/xml"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
)

type Parser struct {
	Version string
	Cfg     *config.MultiConfig
	Sqlx    *sqlx.DB
	Log     *slog.Logger
}

func New(version string) *Parser {
	return &Parser{
		Version: version,
	}
}

func (p *Parser) Init() {

	p.Version = "v0.1.0"
	p.Cfg = config.MustLoad()
	p.Log = logger.InitLogger(p.Cfg.Config.Env)
	p.Log.Info("starting parser on env: " + p.Cfg.Config.Env)
	p.Log.Debug("debug message is enabled")
	var postgresStore = postgres.NewStore()
	postgresStore, err := postgresStore.Init(p.Cfg.Envs, p.Log)
	if err != nil {
		p.Log.Error("Error init postgresStore:", err)
		os.Exit(1)
	}
	p.Sqlx = postgresStore.Dbx
}

func (p *Parser) Run() {
	// init moved files
	// result, err := moved.MovedInputFiles(Cfg.Config, p.Log)
	// if err != nil {
	// 	log.Error("Error moving input files:", err)
	// 	os.Exit(1)
	// }
	// if moved.MovedImages(result.NewPath, p.Log) != nil {
	// 	log.Error("Error moving images:", err)
	// 	os.Exit(1)
	// }
	result := moved.MovedInputFilesT{
		Files: []moved.InputFilesT{
			{TypeFile: "classificator", PathFile: "input/import0_1.xml"},
			{TypeFile: "offer", PathFile: "input/offers0_1.xml"},
		},
		NewPath: "newPath",
	}

	// init parser
	p.ReadAndParse(&result)

	// TODO: graceful shutdown
	p.Sqlx.Close()
	p.Log.Info("DB shutdown: " + p.Cfg.Envs.DB_DATABASE)
}

// читаем файлы и обмена и записываем в базу
func (p *Parser) ReadAndParse(moved *moved.MovedInputFilesT) error {
	utils.PrintAsJSON(moved.Files)

	for _, fileItem := range moved.Files {
		file, err := os.Open(fileItem.PathFile)
		if err != nil {
			p.Log.Error("Ошибка открытия файла:", err)
			return err
		}
		defer file.Close()

		switch fileItem.TypeFile {
		case "classificator":
			var temp XMLtypes.ImportType             // создаем экземпляр структуры
			xmlStruct := temp.КоммерческаяИнформация // создаем подчиненный узел для декодирования
			decoder := xml.NewDecoder(file)
			if err := decoder.Decode(&xmlStruct); err != nil {
				p.Log.Error("Ошибка декодирования XML:", err)
				return err
			}
			// передаем полный тип, чтобы не выделять подчиненный узел в парсере
			utils.PrintAsJSON(moved)
			ImportParser(p, &XMLtypes.ImportType{КоммерческаяИнформация: xmlStruct}, fileItem.PathFile, moved.NewPath)
		case "offer":
			var temp XMLtypes.OfferType
			xmlStruct := temp.КоммерческаяИнформация
			decoder := xml.NewDecoder(file)
			if err := decoder.Decode(&xmlStruct); err != nil {
				p.Log.Error("Ошибка декодирования XML:", err)
				return err
			}
			// передаем полный тип, чтобы не выделять подчиненный узел в парсере
			OfferParser(&XMLtypes.OfferType{КоммерческаяИнформация: xmlStruct}, p)
		}
	}
	return nil
}
