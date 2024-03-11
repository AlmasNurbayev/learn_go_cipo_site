package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Envs struct {
	DB_HOST     string `env:"DB_HOST"`
	DB_PORT     int    `env:"DB_PORT"`
	DB_USERNAME string `env:"DB_USERNAME"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
	DB_DATABASE string `env:"DB_DATABASE"`
	CONFIG_PATH string `env:"CONFIG_PATH"`
}

type Config struct {
	Env    string `yaml:"env" env-required:"true"`
	Server struct {
		Addr         string        `yaml:"addr" env-required:"true"`
		Http_port    int           `yaml:"http_port" env-required:"true"`
		Ssl_port     int           `yaml:"ssl_port" env-required:"true"`
		Timeout      time.Duration `yaml:"timeout" env-required:"true"`
		Idle_timeout time.Duration `yaml:"idle_timeout" env-required:"true"`
	} `yaml:"server"`
	Parser struct {
		Classificator_name string `yaml:"classificator_name" env-required:"true"`
		Offer_name         string `yaml:"offer_name" env-required:"true"`
		ImageFolder_name   string `yaml:"imageFolder_name" env-required:"true"`
		Default_user_id    int64  `yaml:"default_user_id" env-required:"true"`
	} `yaml:"parser"`
}

type MultiConfig struct {
	Envs
	Config
}

func MustLoad() *MultiConfig {

	var config Config
	var envs Envs

	//
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("cannot load env: %s", err)
		os.Exit(1)
	}

	errEnv := cleanenv.ReadEnv(&envs)
	if errEnv != nil {
		log.Printf("cannot read env: %s", errEnv)
		os.Exit(1)
	}

	if _, err := os.Stat(os.Getenv("CONFIG_PATH")); os.IsNotExist(err) {
		log.Printf("config file does not exist: %s", envs.CONFIG_PATH)
		os.Exit(1)
	}

	if errConfig := cleanenv.ReadConfig(os.Getenv("CONFIG_PATH"), &config); errConfig != nil {
		log.Printf("cannot read config file: %s", errConfig)
		os.Exit(1)
	}

	return &MultiConfig{Envs: envs, Config: config}
}
