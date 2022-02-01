package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.Createuser)

	app.Post("/user/:userId", controllers.GetAUser)

	app.Put("/user/:userId", controllers.EditAUser)

	app.Delete("user/:userId", controllers.DeleteAUser)

	app.Get("/users", controllers.GetAllUser)
}