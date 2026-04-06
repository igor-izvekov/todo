package handlers

import (
    "net/http"
	"strconv"

    "github.com/gin-gonic/gin"
    "github.com/igor-izvekov/todo/pkg/database"
    "github.com/igor-izvekov/todo/pkg/models"
)

func GetTask(c *gin.Context) {
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
        c.JSON(http.StatusForbidden, gin.H{"error": "Нет доступа к этой задаче"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "task": task,
    })
}
