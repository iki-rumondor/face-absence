package response

import "time"

type SchoolFee struct {
	Uuid       string              `json:"uuid"`
	Date       string              `json:"tanggal"`
	Nominal    int                 `json:"nominal"`
	Month      string              `json:"month"`
	Status     string              `json:"status"`
	Student    *StudentResponse    `json:"student"`
	SchoolYear *SchoolYearResponse `json:"school_year"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}
