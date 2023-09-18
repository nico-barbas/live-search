package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	brand string = "Boogle"
)

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("search.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": brand,
		}, "layouts/main")
	})

	app.Post("/search", func(c *fiber.Ctx) error {
		// query := c.FormValue("query")
		return c.JSON("hello world")
	})

	log.Fatal(app.Listen(":5000"))
}
