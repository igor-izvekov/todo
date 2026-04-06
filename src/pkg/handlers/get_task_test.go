package handlers_test

import (
    "net/http"
    "strconv"
	"testing"

    "github.com/gin-gonic/gin"
    "github.com/igor-izvekov/todo/pkg/models"
    "github.com/igor-izvekov/todo/pkg/database"
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

    GetTask(c)

    assert.Equal(t, http.StatusOK, w.Code)

    var resp map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &resp)
    taskResp := resp["task"].(map[string]interface{})
    assert.Equal(t, float64(task.ID), taskResp["ID"])
    assert.Equal(t, task.Title, taskResp["Title"])
}

func TestGetTask_NotFound(t *testing.T) {
    setupTestDB()
    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    c.Params = gin.Params{{Key: "id", Value: "999"}}
    c.Request = httptest.NewRequest("GET", "/tasks/999?userID=1", nil)

    GetTask(c)
    assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetTask_Forbidden(t *testing.T) {
    setupTestDB()
    task := createTestTask(1, "Моя задача")

    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    c.Params = gin.Params{{Key: "id", Value: "1"}}
    // Пытается получить чужой пользователь
    c.Request = httptest.NewRequest("GET", "/tasks/1?userID=2", nil)

    GetTask(c)
    assert.Equal(t, http.StatusForbidden, w.Code)
}
