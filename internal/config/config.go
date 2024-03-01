package config

import (
	"log"
	"os"

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
		Http_port    int    `yaml:"http_port" env-required:"true"`
		Ssl_port     int    `yaml:"ssl_port" env-required:"true"`
		Timeout      string `yaml:"timeout" env-required:"true"`
		Idle_timeout string `yaml:"idle_timeout" env-required:"true"`
	} `yaml:"server"`
	Parser struct {
		Classificator_name string `yaml:"classificator_name" env-required:"true"`
		Offer_name         string `yaml:"offer_name" env-required:"true"`
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
	godotenv.Load(".env")
	errEnv := cleanenv.ReadEnv(&envs)
	if errEnv != nil {
		log.Fatalf("cannot read env: %s", errEnv)
	}

	if _, err := os.Stat(os.Getenv("CONFIG_PATH")); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", envs.CONFIG_PATH)
	}

	if errConfig := cleanenv.ReadConfig(os.Getenv("CONFIG_PATH"), &config); errConfig != nil {
		log.Fatalf("cannot read config file: %s", errConfig)
	}

	return &MultiConfig{Envs: envs, Config: config}
}
