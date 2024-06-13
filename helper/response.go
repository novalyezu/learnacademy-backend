package helper

import "github.com/go-playground/validator/v10"

type ResponseDto struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	NextPage    int `json:"next_page"`
	PrevPage    int `json:"prev_page"`
	TotalPage   int `json:"total_page"`
	TotalData   int `json:"total_data"`
}

func WrapperResponse(code int, status string, message string, data interface{}) ResponseDto {
	return ResponseDto{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func ValidationErrorResponse(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			errors = append(errors, e.Field()+": field is required")
		case "email":
			errors = append(errors, e.Field()+": invalid email")
		case "password":
			errors = append(errors, e.Field()+": password min 6, uppercase, lowercase, number, special char")
		default:
			errors = append(errors, e.Error())
		}
	}
	return errors
}
