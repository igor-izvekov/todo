package handlers

import (
    "net/http"
	"strconv"

    "github.com/gin-gonic/gin"
    "github.com/igor-izvekov/todo/pkg/database"
    "github.com/igor-izvekov/todo/pkg/models"
)

func CompleteTask(c *gin.Context) {
    taskID := c.Param("id")
    if taskID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Не указан ID задачи"})
        return
    }

    userIDStr := c.Query("userID")
    if userIDStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Не указан userID"})
        return
    }
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат userID"})
        return
    }

    db := database.GetDB()
    var task models.Task
    if err := db.First(&task, taskID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
        return
    }

    if task.UserID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Нет прав на изменение этой задачи"})
        return
    }

    task.Completed = true
    if err := db.Save(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении статуса"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Задача отмечена как выполненная",
        "task":    task,
    })
}
