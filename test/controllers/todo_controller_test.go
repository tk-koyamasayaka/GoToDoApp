package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/myapp/GoToDoApp/controllers"
	"github.com/myapp/GoToDoApp/dto"
	"github.com/myapp/GoToDoApp/models"
	"github.com/myapp/GoToDoApp/test/mocks"
	"github.com/stretchr/testify/assert"
)

func setupTest() (*gin.Engine, *mocks.MockTodoRepository) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockRepo := mocks.NewMockTodoRepository()
	controller := controllers.NewTodoController(mockRepo)

	v1 := r.Group("/api/v1")
	todos := v1.Group("/todos")
	{
		todos.POST("/", controller.CreateTodo)
		todos.GET("/", controller.GetTodos)
		todos.GET("/:id", controller.GetTodo)
		todos.PUT("/:id", controller.UpdateTodo)
		todos.DELETE("/:id", controller.DeleteTodo)
	}

	return r, mockRepo
}

func TestCreateTodo(t *testing.T) {
	r, _ := setupTest()

	t.Run("正常系: Todoの作成", func(t *testing.T) {
		reqBody := dto.CreateTodoRequest{
			Title:       "テストTodo",
			Description: "テスト説明",
			Completed:   false,
		}
		body, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/todos/", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.True(t, response["success"].(bool))
		assert.NotNil(t, response["data"])
	})

	t.Run("異常系: 不正なリクエスト（タイトルなし）", func(t *testing.T) {
		reqBody := dto.CreateTodoRequest{
			Description: "テスト説明",
			Completed:   false,
		}
		body, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/todos/", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.False(t, response["success"].(bool))
		assert.NotNil(t, response["errors"])
	})
}

func TestGetTodos(t *testing.T) {
	r, mockRepo := setupTest()

	t.Run("正常系: 空のTodoリスト取得", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/todos/", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.True(t, response["success"].(bool))
		assert.NotNil(t, response["data"])
	})

	t.Run("正常系: Todoが存在する場合のリスト取得", func(t *testing.T) {
		todo := &models.Todo{
			Title:       "テストTodo",
			Description: "テスト説明",
			Completed:   false,
		}
		mockRepo.Create(todo)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/todos/", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.True(t, response["success"].(bool))
		data := response["data"].([]interface{})
		assert.Equal(t, 1, len(data))
	})
}

func TestGetTodo(t *testing.T) {
	r, mockRepo := setupTest()

	t.Run("正常系: 存在するTodoの取得", func(t *testing.T) {
		todo := &models.Todo{
			Title:       "テストTodo",
			Description: "テスト説明",
			Completed:   false,
		}
		mockRepo.Create(todo)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/todos/1", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.True(t, response["success"].(bool))
		assert.NotNil(t, response["data"])
	})

	t.Run("異常系: 存在しないTodoの取得", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/todos/999", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.False(t, response["success"].(bool))
	})

	t.Run("異常系: 不正なID形式", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/todos/invalid", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.False(t, response["success"].(bool))
	})
}

func TestUpdateTodo(t *testing.T) {
	r, mockRepo := setupTest()

	t.Run("正常系: Todoの更新", func(t *testing.T) {
		todo := &models.Todo{
			Title:       "テストTodo",
			Description: "テスト説明",
			Completed:   false,
		}
		mockRepo.Create(todo)

		updateReq := dto.UpdateTodoRequest{
			Title:       "更新後のTodo",
			Description: "更新後の説明",
			Completed:   new(bool),
		}
		body, _ := json.Marshal(updateReq)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/api/v1/todos/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.True(t, response["success"].(bool))
		data := response["data"].(map[string]interface{})
		assert.Equal(t, "更新後のTodo", data["title"])
	})

	t.Run("異常系: 存在しないTodoの更新", func(t *testing.T) {
		updateReq := dto.UpdateTodoRequest{
			Title: "更新後のTodo",
		}
		body, _ := json.Marshal(updateReq)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/api/v1/todos/999", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.False(t, response["success"].(bool))
	})
}

func TestDeleteTodo(t *testing.T) {
	r, mockRepo := setupTest()

	t.Run("正常系: Todoの削除", func(t *testing.T) {
		todo := &models.Todo{
			Title:       "テストTodo",
			Description: "テスト説明",
			Completed:   false,
		}
		mockRepo.Create(todo)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/v1/todos/1", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.True(t, response["success"].(bool))
	})

	t.Run("異常系: 存在しないTodoの削除", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/v1/todos/999", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.False(t, response["success"].(bool))
	})
} 