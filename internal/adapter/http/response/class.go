package response

import "time"

type ClassResponse struct {
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	TeacherID uint      `json:"teacher_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
