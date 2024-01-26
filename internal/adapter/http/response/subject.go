package response

import "time"

type SubjectResponse struct {
	Uuid      string     `json:"uuid"`
	Name      string     `json:"name"`
	Teachers  *[]Teacher `json:"teachers"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
