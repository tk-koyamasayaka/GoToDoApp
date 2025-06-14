package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/myapp/GoToDoApp/dto"
	"github.com/myapp/GoToDoApp/mappers"
	"github.com/myapp/GoToDoApp/repositories"
	"github.com/myapp/GoToDoApp/responses"
	"github.com/myapp/GoToDoApp/validators"
)

type TodoController struct {
	repository repositories.TodoRepository
}

func NewTodoController(repository repositories.TodoRepository) *TodoController {
	return &TodoController{repository: repository}
}

// CreateTodo creates a new todo
func (tc *TodoController) CreateTodo(c *gin.Context) {
	var req dto.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.SendBadRequest(c, validators.TranslateError(err))
		return
	}

	todo := mappers.CreateRequestToTodo(&req)
	if err := tc.repository.Create(todo); err != nil {
		responses.SendInternalServerError(c, err)
		return
	}

	responses.SendCreated(c, mappers.TodoToResponse(todo), "Todoが正常に作成されました")
}

// GetTodos returns all todos
func (tc *TodoController) GetTodos(c *gin.Context) {
	todos, err := tc.repository.FindAll()
	if err != nil {
		responses.SendInternalServerError(c, err)
		return
	}

	responses.NewSuccessResponse(mappers.TodosToResponses(todos), "Todoリストを取得しました").Send(c)
}

// GetTodo returns a specific todo
func (tc *TodoController) GetTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.SendBadRequest(c, "IDの形式が不正です")
		return
	}

	todo, err := tc.repository.FindByID(uint(id))
	if err != nil {
		responses.SendNotFound(c, "指定されたTodoが見つかりません")
		return
	}

	responses.NewSuccessResponse(mappers.TodoToResponse(todo), "Todoを取得しました").Send(c)
}

// UpdateTodo updates a todo
func (tc *TodoController) UpdateTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.SendBadRequest(c, "IDの形式が不正です")
		return
	}

	todo, err := tc.repository.FindByID(uint(id))
	if err != nil {
		responses.SendNotFound(c, "指定されたTodoが見つかりません")
		return
	}

	var req dto.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.SendBadRequest(c, validators.TranslateError(err))
		return
	}

	mappers.UpdateRequestToTodo(todo, &req)
	if err := tc.repository.Update(todo); err != nil {
		responses.SendInternalServerError(c, err)
		return
	}

	responses.NewSuccessResponse(mappers.TodoToResponse(todo), "Todoが正常に更新されました").Send(c)
}

// DeleteTodo deletes a todo
func (tc *TodoController) DeleteTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		responses.SendBadRequest(c, "IDの形式が不正です")
		return
	}

	todo, err := tc.repository.FindByID(uint(id))
	if err != nil {
		responses.SendNotFound(c, "指定されたTodoが見つかりません")
		return
	}

	if err := tc.repository.Delete(todo); err != nil {
		responses.SendInternalServerError(c, err)
		return
	}

	responses.NewSuccessResponse(nil, "Todoが正常に削除されました").Send(c)
} 