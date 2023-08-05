package main

import (
	"os"

	"github.com/Ucuping/go-web-example/databases"
	"github.com/Ucuping/go-web-example/helpers"
	"github.com/Ucuping/go-web-example/routes"
	"github.com/gofiber/fiber/v2"
)

func init() {
	helpers.LoadEnv()
	databases.ConnectToDB()
	// migrations.DropTable()
	// migrations.MigrateTable()
	// seeders.Seeder()
}

func main() {
	// engine := mustache.New("./views", ".mustache")
	// // If you want other engine, just replace with following
	// // Create a new engine with django
	// // engine := django.New("./views", ".django")
	// engine.Reload(true)

	// app := fiber.New(fiber.Config{
	// 	Views: engine,
	// })

	app := fiber.New()

	routes.Route(app)

	app.Listen(":" + os.Getenv("PORT"))
}
