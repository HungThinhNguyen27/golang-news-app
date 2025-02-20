package configs

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

// type Postgres struct {
// 	Host     string `yaml:"HOST" env-required:"true"`
// 	UserName string `yaml:"USERNAME" env-required:"true"`
// 	PassWord string `yaml:"PASSWORD" env-required:"true"`
// 	DB       string `yaml:"DB" env-required:"true"`
// 	PORT     int    `yaml:"PORT" env-required:"true"`
// }

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-required:"true" `
	HTTPServer `yaml:"http_server"`
	// Postgres    `yaml:"POSTGRES"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("yaml", "", "path to the config file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path isn't set")

		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exist: %s", configPath)
	}
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can't read config file: %s", err.Error())
	}
	return &cfg
}
