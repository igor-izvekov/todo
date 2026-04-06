package handlers_test

import (
    "net/http"
    "strconv"
	"testing"

    "github.com/gin-gonic/gin"
    "github.com/igor-izvekov/todo/pkg/models"
    "github.com/igor-izvekov/todo/pkg/database"
)

func TestDeleteTask(t *testing.T) {
    setupTestDB()
    task := createTestTask(1, "Удаляемая задача")

    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    c.Params = gin.Params{{Key: "id", Value: "1"}}
    c.Request = httptest.NewRequest("DELETE", "/tasks/1?userID=1", nil)

    DeleteTask(c)

    assert.Equal(t, http.StatusOK, w.Code)

    // Проверяем, что задача удалена
    var count int64
    database.DB.Model(&models.Task{}).Where("id = ?", task.ID).Count(&count)
    assert.Equal(t, int64(0), count)
}

func TestDeleteTask_Forbidden(t *testing.T) {
    setupTestDB()
    createTestTask(1, "Чужая задача")

    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    c.Params = gin.Params{{Key: "id", Value: "1"}}
    c.Request = httptest.NewRequest("DELETE", "/tasks/1?userID=2", nil)

    DeleteTask(c)
    assert.Equal(t, http.StatusForbidden, w.Code)
}