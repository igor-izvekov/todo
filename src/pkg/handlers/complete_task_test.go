package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/igor-izvekov/todo/pkg/database"
	"github.com/igor-izvekov/todo/pkg/handlers"
	"github.com/igor-izvekov/todo/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestCompleteTask(t *testing.T) {
	setupTestDB()
	task := createTestTask(1, "Завершить задачу")

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = httptest.NewRequest("PATCH", "/tasks/1/complete?userID=1", nil)

	handlers.CompleteTask(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedTask models.Task
	database.DB.First(&updatedTask, task.ID)
	assert.True(t, updatedTask.Completed)

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Задача отмечена как выполненная", resp["message"])
}
