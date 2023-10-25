package response

type FailedResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SuccessResponse struct{
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
