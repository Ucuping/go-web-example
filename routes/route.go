package routes

import (
	"github.com/Ucuping/go-web-example/configs"
	"github.com/Ucuping/go-web-example/controllers"
	"github.com/Ucuping/go-web-example/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func Route(app *fiber.App) {

	app.Use(cors.New(configs.CorsConfig), helmet.New(configs.HelmetConfig), compress.New(configs.CompressConfig))

	// app.Get("/auth/login", controllers.AuthIndex)
	// // app.Get("/auth/logout", controllers.Logout)
	// post := app.Group("/posts", middlewares.AuthMiddleware)

	// post.Get("/", controllers.FetchPost)

	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)
	app.Get("/api/auth/logout", controllers.Logout)

	api := app.Group("/api", middlewares.AuthMiddleware)

	api.Get("/auth/verify", controllers.Verify)

	api.Get("/posts", controllers.FetchPost)
	api.Post("/posts", controllers.CreatePost)
	api.Get("/posts/:id", controllers.ShowPost)
	api.Post("/posts/:id", controllers.UpdatePost)
	api.Delete("/posts/:id", controllers.DeletePost)
}
