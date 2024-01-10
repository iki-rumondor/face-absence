package response

import "time"

type UserData struct {
	Nama      string    `json:"nama"`
	Username  string    `json:"username"`
	Avatar    *string   `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
