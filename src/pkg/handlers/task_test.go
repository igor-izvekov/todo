package handlers_test

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/igor-izvekov/todo/pkg/database"
	"github.com/igor-izvekov/todo/pkg/models"
)

func setupTestDB() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	database.DB = db
	db.AutoMigrate(&models.Task{})
}

func cleanTasks() {
	database.DB.Exec("DELETE FROM tasks")
}

func createTestTask(userID int, title string) models.Task {
	task := models.Task{
		UserID:    userID,
		Title:     title,
		Completed: false,
	}
	database.DB.Create(&task)
	return task
}
