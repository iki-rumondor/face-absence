package response

import (
	"bytes"
	"mime/multipart"
	"time"
)

type FormAbsence struct {
	RequestBody *bytes.Buffer
	Writer      *multipart.Writer
}

type AbsenceResponse struct {
	Uuid      string            `json:"uuid"`
	Status    string            `json:"status"`
	Student   *StudentResponse  `json:"student"`
	Schedule  *ScheduleResponse `json:"schedule"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
