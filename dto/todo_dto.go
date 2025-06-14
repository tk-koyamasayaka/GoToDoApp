package dto

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=100"`
	Description string `json:"description" binding:"max=500"`
	Completed   bool   `json:"completed"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title" binding:"omitempty,min=1,max=100"`
	Description string `json:"description" binding:"max=500"`
	Completed   *bool  `json:"completed" binding:"omitempty"`
}

type TodoResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
} 