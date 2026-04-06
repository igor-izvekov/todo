package handlers_test

import (
    "net/http"
    "strconv"
	"testing"

    "github.com/gin-gonic/gin"
    "github.com/igor-izvekov/todo/pkg/models"
    "github.com/igor-izvekov/todo/pkg/database"
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

    GetTasks(c)

    assert.Equal(t, http.StatusOK, w.Code)

    var resp map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &resp)
    tasks := resp["tasks"].([]interface{})
    assert.Len(t, tasks, 2)

    // Проверим, что первая задача имеет правильный заголовок
    firstTask := tasks[0].(map[string]interface{})
    assert.Equal(t, "Задача 2", firstTask["Title"]) // порядок id desc
}
