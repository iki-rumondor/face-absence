package response

import "time"

type ScheduleResponse struct {
	ID           uint      `json:"id"`
	Uuid         string    `json:"uuid"`
	Name         string    `json:"name"`
	Day          string    `json:"day"`
	Start        string    `json:"start"`
	End          string    `json:"end"`
	ClassID      uint      `json:"class_id"`
	SubjectID    uint      `json:"subject_id"`
	TeacherID    uint      `json:"teacher_id"`
	SchoolYearID uint      `json:"school_year_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
