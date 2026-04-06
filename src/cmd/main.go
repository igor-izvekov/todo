package main

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/igor-izvekov/todo/pkg/database"
	"github.com/igor-izvekov/todo/pkg/migrations"
)

var (
	http_server = "localhost:8000"
)

func run_http_server() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Logger())

	router.LoadHTMLGlob("frontend/" + "*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	taskGroup := router.Group("/tasks")

	{
		taskGroup.POST("/", handlers.CreateTask)
		taskGroup.GET("/", handlers.GetTasks)
		taskGroup.GET("/:id", handlers.GetTask)
		taskGroup.PUT("/:id", handlers.UpdateTask)
		taskGroup.DELETE("/:id", handlers.DeleteTask)
		taskGroup.PATCH("/:id/complete", handlers.CompleteTask)
	}

	if err := router.Run(http_server); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

func main() {
	if err := database.Connect("todo.db"); err != nil {
		panic(err)
	}

	db := database.GetDB()
	if err := migrations.AutoMigrate(db); err != nil {
		panic(err)
	}
	run_http_server()
}