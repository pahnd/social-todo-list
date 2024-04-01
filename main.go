package main

import (
	"log"
	"os"
	"social-todo-list/middleware"
	ginitem "social-todo-list/module/item/transport/gin"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load(".env")
	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	log.Println("DB Connection:", db)

	r := gin.Default()
	r.Use(middleware.Recover())

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}

	if err := r.Run(":3000"); err != nil {
		log.Fatalln(err)
	}
}
