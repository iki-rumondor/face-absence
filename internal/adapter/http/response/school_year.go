package response

import "time"

type SchoolYearResponse struct {
	ID        uint      `json:"id"`
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
