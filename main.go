package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/myapp/GoToDoApp/controllers"
	"github.com/myapp/GoToDoApp/models"
	"github.com/myapp/GoToDoApp/repositories"
	"github.com/myapp/GoToDoApp/validators"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// バリデーターの登録
	validators.RegisterTodoValidators()

	// データベース接続
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	// マイグレーション
	db.AutoMigrate(&models.Todo{})

	// リポジトリの初期化
	todoRepository := repositories.NewTodoRepository(db)

	// コントローラーの初期化
	todoController := controllers.NewTodoController(todoRepository)

	// Ginルーターの設定
	r := gin.Default()

	// ルーティング
	v1 := r.Group("/api/v1")
	{
		todos := v1.Group("/todos")
		{
			todos.POST("/", todoController.CreateTodo)
			todos.GET("/", todoController.GetTodos)
			todos.GET("/:id", todoController.GetTodo)
			todos.PUT("/:id", todoController.UpdateTodo)
			todos.DELETE("/:id", todoController.DeleteTodo)
		}
	}

	// サーバーの起動
	r.Run(":8080")
} 