package response

import "time"

type UserData struct {
	Nama      string    `json:"nama"`
	Username  string    `json:"username"`
	Avatar    *string   `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Dashboard struct {
	Stats map[string]int64       `json:"stats"`
	Graph map[string]interface{} `json:"graph"`
}

type VerifyTokenResponse struct {
	Nama      string    `json:"nama"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
