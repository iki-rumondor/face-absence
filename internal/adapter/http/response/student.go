package response

import "time"

type StudentResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Student struct {
	ID        uint      `json:"id"`
	Uuid      string    `json:"uuid"`
	Nama      string    `json:"nama"`
	NIS       string    `json:"nis"`
	Kelas     string    `json:"kelas"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
