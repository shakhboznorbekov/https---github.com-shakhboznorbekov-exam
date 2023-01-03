package main

import (
	"context"
	"log"

	"github.com/shakhboznorbekov/token/user_api/api"
	"github.com/shakhboznorbekov/token/user_api/config"
	"github.com/shakhboznorbekov/token/user_api/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	storage, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.CloseDB()

	api.SetUpApi(&cfg, r, storage)

	log.Printf("Listening port %v...\n", cfg.HTTPPort)
	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}
