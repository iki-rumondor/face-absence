package response

import "time"

type ScheduleResponse struct {
	Uuid       string              `json:"uuid"`
	Day        string              `json:"day"`
	Start      string              `json:"start"`
	End        string              `json:"end"`
	Class      *ClassData          `json:"class"`
	Subject    *SubjectResponse    `json:"subject"`
	SchoolYear *SchoolYearResponse `json:"school_year"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}

type ScheduleResponseForStudent struct {
	Uuid       string              `json:"uuid"`
	Day        string              `json:"day"`
	Start      string              `json:"start"`
	End        string              `json:"end"`
	Class      *ClassData          `json:"class"`
	Subject    *SubjectResponse    `json:"subject"`
	SchoolYear *SchoolYearResponse `json:"school_year"`
	Absence    *AbsenceResponse    `json:"absence"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}
