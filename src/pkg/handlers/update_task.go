package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/igor-izvekov/todo/pkg/database"
    "github.com/igor-izvekov/todo/pkg/models"
)

type UpdateTaskForm struct {
    UserID int    `json:"userID"`
    Title  string `json:"title"`
}

func UpdateTask(c *gin.Context) {
    taskID := c.Param("id")
    if taskID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Не указан ID задачи"})
        return
    }

    var form UpdateTaskForm
    if err := c.BindJSON(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Некорректный формат данных: " + err.Error(),
        })
        return
    }

    db := database.GetDB()

    var task models.Task
    if err := db.First(&task, taskID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
        return
    }

    if task.UserID != form.UserID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Нет прав на обновление этой задачи"})
        return
    }

    task.Title = form.Title

    if err := db.Save(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить задачу"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Задача обновлена",
        "task":    task,
    })
}
