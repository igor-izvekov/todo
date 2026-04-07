package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/igor-izvekov/todo/pkg/handlers"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	setupTestDB()
	defer cleanTasks()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	reqBody := handlers.CreateTaskForm{
		UserID: 1,
		Title:  "Написать тесты",
	}
	jsonBytes, _ := json.Marshal(reqBody)
	c.Request = httptest.NewRequest("POST", "/tasks", bytes.NewReader(jsonBytes))
	c.Request.Header.Set("Content-Type", "application/json")

	handlers.CreateTask(c)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Задача создана", resp["message"])

	taskMap := resp["task"].(map[string]interface{})
	assert.Equal(t, float64(1), taskMap["UserID"])
	assert.Equal(t, "Написать тесты", taskMap["Title"])
	assert.Equal(t, false, taskMap["Completed"])
}
