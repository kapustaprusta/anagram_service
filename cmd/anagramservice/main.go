package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/kapustaprusta/anagram_service/internal/app/anagramservice"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/anagramservice.toml", "path to configuration file")
}

func main() {
	flag.Parse()

	config := anagramservice.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("start listening at", config.BindAddr)
	if err := anagramservice.Start(config); err != nil {
		log.Fatal(err)
	}
}
