package response

import "time"

type SchoolFee struct {
	Uuid      string           `json:"uuid"`
	Date      string           `json:"tanggal"`
	Nominal   int              `json:"nominal"`
	Student   *StudentResponse `json:"student"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
