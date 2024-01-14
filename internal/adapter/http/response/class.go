package response

import "time"

type ClassResponse struct {
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Teacher   *Teacher  `json:"teacher"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClassData struct {
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClassOption struct {
	Uuid string `json:"value"`
	Name string `json:"label"`
}
