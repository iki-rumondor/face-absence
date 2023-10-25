package response

type Message struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
