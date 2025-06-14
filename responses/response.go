package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse は共通のレスポンス構造体
type APIResponse struct {
	Status  int         `json:"-"`          // HTTPステータスコード（JSONには含めない）
	Success bool        `json:"success"`    // 処理の成功/失敗
	Message string      `json:"message"`    // 処理結果のメッセージ
	Data    interface{} `json:"data"`       // レスポンスデータ
	Errors  interface{} `json:"errors"`     // エラー情報
}

// NewSuccessResponse 成功レスポンスを作成
func NewSuccessResponse(data interface{}, message string) *APIResponse {
	return &APIResponse{
		Status:  http.StatusOK,
		Success: true,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
}

// NewErrorResponse エラーレスポンスを作成
func NewErrorResponse(status int, message string, errors interface{}) *APIResponse {
	return &APIResponse{
		Status:  status,
		Success: false,
		Message: message,
		Data:    nil,
		Errors:  errors,
	}
}

// Send レスポンスを送信
func (r *APIResponse) Send(c *gin.Context) {
	c.JSON(r.Status, r)
}

// HTTP Status Codeに応じたヘルパー関数
func SendCreated(c *gin.Context, data interface{}, message string) {
	response := &APIResponse{
		Status:  http.StatusCreated,
		Success: true,
		Message: message,
		Data:    data,
	}
	response.Send(c)
}

func SendBadRequest(c *gin.Context, errors interface{}) {
	response := NewErrorResponse(
		http.StatusBadRequest,
		"リクエストが不正です",
		errors,
	)
	response.Send(c)
}

func SendNotFound(c *gin.Context, message string) {
	response := NewErrorResponse(
		http.StatusNotFound,
		message,
		nil,
	)
	response.Send(c)
}

func SendInternalServerError(c *gin.Context, err error) {
	response := NewErrorResponse(
		http.StatusInternalServerError,
		"内部サーバーエラーが発生しました",
		err.Error(),
	)
	response.Send(c)
} 