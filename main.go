package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ledo01/shorten/api"
	"github.com/ledo01/shorten/repository/redis"
	"github.com/ledo01/shorten/shorten"
)

func main() {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	repo, err := redis.NewRedisRepository(config.Database.URL)
	if err != nil {
		panic(err)
	}
	service := shorten.NewRedirectService(repo)
	handler := api.NewHandler(service)

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/:code", handler.Get)
	app.Post("/", handler.Post)

	app.Listen(":" + config.Server.Port)
}
