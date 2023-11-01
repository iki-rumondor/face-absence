package response

import "time"

type Teacher struct {
	ID        uint      `json:"id"`
	Uuid      string    `json:"uuid"`
	Nama      string    `json:"nama"`
	Email     string    `json:"email"`
	Nip       string    `json:"nip"`
	JK        string    `json:"jenis_kelamin"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
