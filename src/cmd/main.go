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
	http_server = "localhost:8080"
)

func run_http_server() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Logger())

	router.LoadHTMLGlob("frontend/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.GET("/todo", func(c *gin.Context) {
		c.HTML(200, "todo.html", nil)
	})

	api := router.Group("/api")
	{
		api.POST("/auth/login", auth.HandleLoginOrRegister)

		taskGroup := api.Group("/tasks")
		taskGroup.Use(auth.AuthMiddleware())
		{
			taskGroup.POST("/", handlers.CreateTask)
			taskGroup.GET("/", handlers.GetTasks)
			taskGroup.GET("/:id", handlers.GetTask)
			taskGroup.PUT("/:id", handlers.UpdateTask)
			taskGroup.DELETE("/:id", handlers.DeleteTask)
			taskGroup.PATCH("/:id/complete", handlers.CompleteTask)
		}
	}

	log.Printf("Сервер запущен на http://%s", http_server)
	if err := router.Run(http_server); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}

func main() {
	if err := database.Connect("todo.db"); err != nil {
		panic("Ошибка подключения к БД: " + err.Error())
	}

	db := database.GetDB()
	if err := migrations.AutoMigrate(db); err != nil {
		panic("Ошибка миграции: " + err.Error())
	}

	run_http_server()
}
