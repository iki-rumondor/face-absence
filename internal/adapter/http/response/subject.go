package response

import "time"

type SubjectResponse struct {
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	Teacher   *Teacher  `json:"teacher"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}