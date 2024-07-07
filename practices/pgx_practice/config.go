package pgxPractice

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type CfgServer struct {
	Address     string        `yaml:"address" env-default:"localhost:5000"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type CfgPsql struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5500"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"root"`
	DbName   string `yaml:"dbname" env-default:"database"`
}

type Config struct {
	Env string `yaml:"env", env-default:"dev"`
	Server     CfgServer `yaml:"http_server"`
	Postgres   CfgPsql        `yaml:"postgres"`
}

var configInstance *Config

func NewConfig() *Config {
	if configInstance != nil {
		return configInstance
	}

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	configInstance = &cfg

	return configInstance

}
