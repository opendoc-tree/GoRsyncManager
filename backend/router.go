package main

import (
	"GoRsyncManager/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Routes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/get_hosts", handlers.GetHosts)
	app.Post("/get_files", handlers.GetFilesFromHost)
	app.Post("/add_host", handlers.AddHost)
	app.Delete("/del_host/:id", handlers.DelHost)
}
