package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	Id          uint
	Title       string
	Description string
}

func main() {
	dsn := "root:Ntare@1995@tcp(127.0.0.1:3306)/posts_ms?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Post{})

	app := fiber.New()

	// app.Use(cors.New())

	app.Get("/api/posts", func(c *fiber.Ctx) error {
		var posts []Post
		db.Find(&posts)
		return c.JSON(posts)
	})

	app.Listen(":3000")
}
