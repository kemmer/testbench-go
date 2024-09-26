package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigDatabase struct {
	DBString string `yaml:"db_string" env:"DB_STRING" env-default:"aaaaaaaaaaaaaaaaaaaaa"`
}

func main() {
	var cfg ConfigDatabase

	err := cleanenv.ReadConfig("./cmd/cleanenv-lesson/config.yaml", &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", cfg)
}
