package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/thanadonexe/skooldio-robust-api-with-go/auth"
	"github.com/thanadonexe/skooldio-robust-api-with-go/todo"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Configuration file not found")
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&todo.Todo{})
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	signature := os.Getenv("SIGNATURE")
	r.GET("/tokenz", auth.AccessToken(signature))
	protected := r.Group("", auth.Protect([]byte(signature)))
	handler := todo.NewTodoHandler(db)
	protected.POST("/todos", handler.NewTask)
	r.Run()
}
