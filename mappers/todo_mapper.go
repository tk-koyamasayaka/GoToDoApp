package mappers

import (
	"github.com/myapp/GoToDoApp/dto"
	"github.com/myapp/GoToDoApp/models"
)

func CreateRequestToTodo(req *dto.CreateTodoRequest) *models.Todo {
	return &models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	}
}

func UpdateRequestToTodo(todo *models.Todo, req *dto.UpdateTodoRequest) {
	if req.Title != "" {
		todo.Title = req.Title
	}
	if req.Description != "" {
		todo.Description = req.Description
	}
	if req.Completed != nil {
		todo.Completed = *req.Completed
	}
}

func TodoToResponse(todo *models.Todo) *dto.TodoResponse {
	return &dto.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}
}

func TodosToResponses(todos []models.Todo) []dto.TodoResponse {
	responses := make([]dto.TodoResponse, len(todos))
	for i, todo := range todos {
		responses[i] = *TodoToResponse(&todo)
	}
	return responses
} 