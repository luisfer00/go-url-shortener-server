package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luisfer00/url-shortener/pkg/controllers"
)

func RegisterURLRoutes(rg fiber.Router) {
	rg.Get("/:slug", controllers.GetURLController)
	rg.Post("/", controllers.InsertURLController)
}