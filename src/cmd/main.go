package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/igor-izvekov/todo/pkg/auth"
	"github.com/igor-izvekov/todo/pkg/database"
	"github.com/igor-izvekov/todo/pkg/handlers"
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
	router.Static("frontend/", "styles.css")

	taskGroup := router.Group("/tasks")

	{
		taskGroup.POST("/", handlers.CreateTask)
		taskGroup.GET("/", handlers.GetTasks)
		taskGroup.GET("/:id", handlers.GetTask)
		taskGroup.PUT("/:id", handlers.UpdateTask)
		taskGroup.DELETE("/:id", handlers.DeleteTask)
		taskGroup.PATCH("/:id/complete", handlers.CompleteTask)
	}

	router.GET("/", auth.HandleHome)
	router.GET("/login", auth.HandleLogin)
	router.GET("/auth/callback", auth.HandleCallback)

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
