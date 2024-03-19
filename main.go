package main

import (
	"GoRsyncManager/configs"
	"log"

	_ "GoRsyncManager/docs"

	"github.com/gofiber/fiber/v2"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @schemes http
func main() {
	app := fiber.New()
	Routes(app)
	configs.ConnectDB()
	log.Fatal(app.Listen(":3000"))
}
