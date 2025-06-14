package repositories

import "github.com/myapp/GoToDoApp/models"

type TodoRepository interface {
	Create(todo *models.Todo) error
	FindAll() ([]models.Todo, error)
	FindByID(id uint) (*models.Todo, error)
	Update(todo *models.Todo) error
	Delete(todo *models.Todo) error
} 