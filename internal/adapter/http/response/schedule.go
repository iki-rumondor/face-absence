package response

import "time"

type ScheduleResponse struct {
	ID           uint      `json:"id"`
	Uuid         string    `json:"uuid"`
	Name         string    `json:"name"`
	Day          time.Time `json:"day"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	ClassID      uint      `json:"class_id"`
	SubjectID    uint      `json:"subject_id"`
	TeacherID    uint      `json:"teacher_id"`
	SchoolYearID uint      `json:"school_year_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
