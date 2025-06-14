package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// カスタムバリデーションルールを登録
func RegisterTodoValidators() {
	if _, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// ここにカスタムバリデーションルールを追加できます
		// 例: v.RegisterValidation("custom_rule", customValidationFunc)
	}
}

// バリデーションエラーをわかりやすいメッセージに変換
func TranslateError(err error) map[string]string {
	errs := make(map[string]string)

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrs {
			switch e.Tag() {
			case "required":
				errs[e.Field()] = "このフィールドは必須です"
			case "min":
				errs[e.Field()] = "最小長さを満たしていません"
			case "max":
				errs[e.Field()] = "最大長さを超えています"
			default:
				errs[e.Field()] = "無効な値です"
			}
		}
		return errs
	}

	errs["error"] = err.Error()
	return errs
} 