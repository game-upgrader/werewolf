package server

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"uwwolf/app/handler"
	"uwwolf/app/middleware"
	"uwwolf/app/types"
	"uwwolf/config"
)

func StartAPI() {
	app := fiber.New()

	app.Post("/api/v1/init", middleware.Validation[types.GameSetting], handler.CreateGame)

	log.Fatal(app.Listen(":" + strconv.Itoa(config.App.HttpPort)))
}
