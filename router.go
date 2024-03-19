package main

import (
	"GoRsyncManager/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Routes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/get_nodes", handlers.GetNodes)
	app.Get("/get_node/:id", handlers.GetNodeById)
	app.Post("/add_node", handlers.AddNode)
	app.Delete("/del_node/:id", handlers.DelNode)
}
