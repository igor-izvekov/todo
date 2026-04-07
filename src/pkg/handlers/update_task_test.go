package handlers_test

import (
	"bytes"
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

func TestUpdateTask(t *testing.T) {
	setupTestDB()
	defer cleanTasks()

	task := createTestTask(1, "Старое название")

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	updateForm := handlers.UpdateTaskForm{
		UserID: 1,
		Title:  "Новое название",
	}
	jsonBytes, _ := json.Marshal(updateForm)

	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = httptest.NewRequest("PUT", "/tasks/1", bytes.NewReader(jsonBytes))
	c.Request.Header.Set("Content-Type", "application/json")

	handlers.UpdateTask(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	updatedTask := resp["task"].(map[string]interface{})
	assert.Equal(t, "Новое название", updatedTask["Title"])

	var dbTask models.Task
	database.DB.First(&dbTask, task.ID)
	assert.Equal(t, "Новое название", dbTask.Title)
}
