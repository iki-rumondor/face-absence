package request

type CreateAbsence struct {
	ScheduleID uint `json:"schedule_id" valid:"required~field schedule_id is not found"`
}
