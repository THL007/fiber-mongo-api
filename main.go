package main

import(
	"github.com/gofiber/fiber/v2"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"
)

func main() {
	app := fiber.New()

	// run database
	configs.ConnectDB()

	// routes
	routes.UserRoute(app)
	// dunno

	// app.Get("/", func(c *fiber.Ctx) error{
	// 	return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	// })

	app.Listen(":6000")
}