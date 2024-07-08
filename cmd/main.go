package main

import (
	"github.com/eunice99x/go-commerce/internal/api"
	"github.com/eunice99x/go-commerce/internal/database"
	"log"
)

func main() {
	database.Connect()
	apiCfg := &api.Config{

		Port: 8080,
	}

	_api, err := api.New(apiCfg)
	if err != nil {
		log.Fatalf("api: %v", err)
	}

	if err := _api.Listen(); err != nil {
		log.Fatalf("listen: %v", err)
	}
}
