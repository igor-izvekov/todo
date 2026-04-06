package handlers

import (
    "net/http"
	"strconv"

    "github.com/gin-gonic/gin"
    "github.com/igor-izvekov/todo/pkg/database"
    "github.com/igor-izvekov/todo/pkg/models"
)

func GetTasks(c *gin.Context) {
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
    var tasks []models.Task
    if err := db.Where("user_id = ?", userID).Order("id desc").Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении задач"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "tasks": tasks,
    })
}
