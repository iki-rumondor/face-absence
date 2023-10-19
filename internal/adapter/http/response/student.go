package response

type Student struct{
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}