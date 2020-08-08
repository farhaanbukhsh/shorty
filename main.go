package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/farhaanbukhsh/shorty/config"
	"github.com/farhaanbukhsh/shorty/server"
	"github.com/farhaanbukhsh/shorty/storage"
)

func main() {
	configFile := flag.String("config", "./config/config.json", "Config file having information")
	flag.Parse()
	fileData, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("Not able to read file")
	}

	var cfg config.Config
	err = json.Unmarshal(fileData, &cfg)
	if err != nil {
		log.Fatalf("Config cannot be un-marshalled")
	}
	svc, err := storage.New(cfg.Database)
	if err != nil {
		log.Fatalf("Cannot connect to database, %s", err)
	}
	server.StartServer(cfg, svc)
	defer svc.Close()
}
