package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/igor-izvekov/todo/pkg/database"
	"github.com/igor-izvekov/todo/pkg/handlers"
	"github.com/igor-izvekov/todo/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTask(t *testing.T) {
	setupTestDB()
	task := createTestTask(1, "Удаляемая задача")

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = httptest.NewRequest("DELETE", "/tasks/1?userID=1", nil)

	handlers.DeleteTask(c)

	assert.Equal(t, http.StatusOK, w.Code)

	// Проверяем, что задача удалена
	var count int64
	database.DB.Model(&models.Task{}).Where("id = ?", task.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}
