package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/luisfer00/url-shortener/pkg/routes"
)

func main() {
	app := fiber.New()

	PORT := os.Getenv("PORT")
	addr := fmt.Sprintf(":%v", PORT)
	
	app.Use(cors.New())
	urlRoutes := app.Group("/api/url")
	routes.RegisterURLRoutes(urlRoutes)

	log.Fatalln(app.Listen(addr))
}