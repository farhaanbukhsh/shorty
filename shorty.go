package main

import (
	"flag"
	"log"

	"github.com/farhaanbukhsh/shorty/config"
	"github.com/farhaanbukhsh/shorty/server"
	"github.com/farhaanbukhsh/shorty/storage"
)

func main() {
	configFile := flag.String("config", "./config/config.json", "Config file having information")
	flag.Parse()
	cfg := config.ReadConfig(*configFile)
	svc, err := storage.New(cfg.Database)
	if err != nil {
		log.Fatalf("Cannot connect to database, %s", err)
	}
	server.StartServer(cfg, svc)
	defer svc.Close()
}
