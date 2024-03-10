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
	Tx      *sqlx.Tx
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
		panic(err)
	}
	p.Sqlx = postgresStore.Dbx
	p.Tx = postgresStore.Tx
}

func (p *Parser) Run() {
	//init moved files
	// movedFiles, err := moved.MovedInputFiles(p.Cfg.Config, p.Log)
	// if err != nil {
	// 	p.Log.Error("Error moving input files:", err)
	// 	os.Exit(1)
	// }
	// if moved.CopyImages("assets", movedFiles.NewPath, p.Log) != nil {
	// 	p.Log.Error("Error copying images: ", err)
	// 	os.Exit(1)
	// }
	movedFiles := moved.MovedInputFilesT{
		Files: []moved.InputFilesT{
			{TypeFile: "classificator", PathFile: "input/import0_1.xml"},
			{TypeFile: "offer", PathFile: "input/offers0_1.xml"},
		},
		NewPath: "input",
	}
	utils.PrintAsJSON(movedFiles)

	// init parser
	for _, fileItem := range movedFiles.Files {
		file, err := os.Open(fileItem.PathFile)
		if err != nil {
			p.Log.Error("Ошибка открытия файла:", err)
			os.Exit(1)
		}
		defer file.Close()

		switch fileItem.TypeFile {
		case "classificator":
			var temp XMLtypes.ImportType             // создаем экземпляр структуры
			xmlStruct := temp.КоммерческаяИнформация // создаем подчиненный узел для декодирования
			decoder := xml.NewDecoder(file)
			decoder.Strict = false
			if err := decoder.Decode(&xmlStruct); err != nil {
				p.Log.Error("Ошибка декодирования XML:", err)
				panic(err)
			}
			// передаем полный тип, чтобы не выделять подчиненный узел в парсере
			err := ImportParser(p, &XMLtypes.ImportType{КоммерческаяИнформация: xmlStruct}, fileItem.PathFile, movedFiles.NewPath)
			if err != nil {
				p.Log.Error("Error import and rollback all db changes:", err)
				err = p.Tx.Rollback()
				if err != nil {
					p.Log.Error("Error rollback all db changes:", err)
				}
				p.Sqlx.Close()
				panic(err)
			}
		case "offer":
			var temp XMLtypes.OfferType
			xmlStruct := temp.КоммерческаяИнформация
			decoder := xml.NewDecoder(file)
			if err := decoder.Decode(&xmlStruct); err != nil {
				p.Log.Error("Ошибка декодирования XML:", err)
				panic(err)
			}
			// передаем полный тип, чтобы не выделять подчиненный узел в парсере
			err := OfferParser(p, &XMLtypes.OfferType{КоммерческаяИнформация: xmlStruct}, fileItem.PathFile, movedFiles.NewPath)
			if err != nil {
				p.Log.Error("Error offer and rollback all db changes:", err)
				err = p.Tx.Rollback()
				if err != nil {
					p.Log.Error("Error rollback all db changes:", err)
				}
				p.Sqlx.Close()
				panic(err)
			}
		}
	}

	// TODO: graceful shutdown
	err := p.Tx.Commit()
	if err != nil {
		p.Log.Error("Error commit all db changes:", err)
	}
	p.Sqlx.Close()

	p.Log.Info("DB shutdown: " + p.Cfg.Envs.DB_DATABASE)
}
