package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/amirhossein-programmer/fiber-api/cmd/web/pkg/database"
)

func main() {
	database.ConnectDb()
	engine := html.New("../../UI/html/pages", ".html")
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "fiber",
		AppName:       "My App v1.0.1",
		Views:         engine,
		ViewsLayout:   "../../UI/html/templates/",
	})
	app.Static("/static/", "../../UI/static")
	SetupRoutes(app)
	log.Fatal(app.Listen(":4000"))
}
