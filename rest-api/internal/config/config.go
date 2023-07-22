package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	PORT   string `koanf:"PORT"`
	DB_URL string `koanf:"DB_URL"`
}

var k = koanf.New(".")

var env *string = flag.String("env", "development", "App environment - |development|production|")

func Load() Config {
	var config Config
	flag.Parse()

	if err := k.Load(file.Provider(fmt.Sprintf("./internal/config/%v.env", *env)), dotenv.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := k.Unmarshal("", &config); err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	return config
}
