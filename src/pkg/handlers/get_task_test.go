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

func TestGetTask(t *testing.T) {
	setupTestDB()
	defer cleanTasks()

	task := createTestTask(1, "Уникальная задача")

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = httptest.NewRequest("GET", "/tasks/1?userID=1", nil)

	handlers.GetTask(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	taskResp := resp["task"].(map[string]interface{})
	assert.Equal(t, float64(task.ID), taskResp["ID"])
	assert.Equal(t, task.Title, taskResp["Title"])
}
