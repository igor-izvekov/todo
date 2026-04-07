package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/igor-izvekov/todo/pkg/handlers"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	setupTestDB()
	defer cleanTasks()

	createTestTask(1, "Задача 1")
	createTestTask(1, "Задача 2")
	createTestTask(2, "Чужая задача")

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("GET", "/tasks?userID=1", nil)

	handlers.GetTasks(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	tasks := resp["tasks"].([]interface{})
	assert.Len(t, tasks, 2)

	firstTask := tasks[0].(map[string]interface{})
	assert.Equal(t, "Задача 2", firstTask["Title"])
}
