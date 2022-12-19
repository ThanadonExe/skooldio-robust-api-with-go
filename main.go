package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/thanadonexe/skooldio-robust-api-with-go/auth"
	"github.com/thanadonexe/skooldio-robust-api-with-go/todo"
)

func main() {
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

	r.GET("/tokenz", auth.AccessToken("==Signature=="))
	protected := r.Group("", auth.Protect([]byte("==Signature==")))
	handler := todo.NewTodoHandler(db)
	protected.POST("/todos", handler.NewTask)
	r.Run()
}
