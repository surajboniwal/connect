package config

import (
	"fmt"
	"log"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	PORT    string `koanf:"PORT"`
	DB_URL  string `koanf:"DB_URL"`
	APP_ENV string
}

var k = koanf.New(".")

func Load(env string) Config {
	var config Config = Config{
		APP_ENV: env,
	}

	if err := k.Load(file.Provider(fmt.Sprintf("./internal/config/%v.env", env)), dotenv.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := k.Unmarshal("", &config); err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	return config
}
