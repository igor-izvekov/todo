package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/igor-izvekov/todo/pkg/database"
    "github.com/igor-izvekov/todo/pkg/models"
)

type CreateTaskForm struct {
    UserID int    `json:"userID"`
    Title  string `json:"title"`
}

func CreateTask(c *gin.Context) {
    var form CreateTaskForm
    if err := c.BindJSON(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат: " + err.Error()})
        return
    }

    if form.Title == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Название задачи не может быть пустым"})
        return
    }

    task := models.Task{
        UserID:    form.UserID,
        Title:     form.Title,
        Completed: false,
    }

    db := database.GetDB()
    if err := db.Create(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании задачи"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Задача создана",
        "task":    task,
    })
}
