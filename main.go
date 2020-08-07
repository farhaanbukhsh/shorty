package main

import (
	"log"

	"github.com/farhaanbukhsh/shorty/server"
	"github.com/farhaanbukhsh/shorty/storage"
)

func main() {
	svc, err := storage.New("shorty.db")
	if err != nil {
		log.Fatalf("Cannot connect to database, %s", err)
	}
	server.StartServer(":8080", svc)
	defer svc.Close()
}
