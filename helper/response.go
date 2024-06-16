package helper

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
