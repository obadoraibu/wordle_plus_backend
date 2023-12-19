package main

import (
	"flag"
	"github.com/obadoraibu/wordle_plus_backend.git/internal/app"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "mainCfgPath", "configs/main.yml", "path to the main config file")
}

func main() {
	flag.Parse()

	if err := app.Run(configPath); err != nil {
		log.Fatal("cannot run the app")
	}
}
