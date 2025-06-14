package mocks

import (
	"errors"

	"github.com/myapp/GoToDoApp/models"
)

type MockTodoRepository struct {
	todos  map[uint]*models.Todo
	nextID uint
}

func NewMockTodoRepository() *MockTodoRepository {
	return &MockTodoRepository{
		todos:  make(map[uint]*models.Todo),
		nextID: 1,
	}
}

func (r *MockTodoRepository) Create(todo *models.Todo) error {
	todo.ID = r.nextID
	r.todos[todo.ID] = todo
	r.nextID++
	return nil
}

func (r *MockTodoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo
	for _, todo := range r.todos {
		todos = append(todos, *todo)
	}
	return todos, nil
}

func (r *MockTodoRepository) FindByID(id uint) (*models.Todo, error) {
	if todo, exists := r.todos[id]; exists {
		return todo, nil
	}
	return nil, errors.New("todo not found")
}

func (r *MockTodoRepository) Update(todo *models.Todo) error {
	if _, exists := r.todos[todo.ID]; !exists {
		return errors.New("todo not found")
	}
	r.todos[todo.ID] = todo
	return nil
}

func (r *MockTodoRepository) Delete(todo *models.Todo) error {
	if _, exists := r.todos[todo.ID]; !exists {
		return errors.New("todo not found")
	}
	delete(r.todos, todo.ID)
	return nil
} 