package response

import "time"

type StudentUser struct {
	ID        uint      `json:"id"`
	Uuid      string    `json:"uuid"`
	Nama      string    `json:"nama"`
	Username  string    `json:"username"`
	NIS       string    `json:"nis"`
	JK        string    `json:"jk"`
	Kelas     string    `json:"kelas"`
	Semester  string    `json:"semester"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FailedStudent struct {
	Nama        string `json:"nama"`
	Description string `json:"description"`
	Error       string `json:"error"`
}
