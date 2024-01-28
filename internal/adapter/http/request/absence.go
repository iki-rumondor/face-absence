package request

type CreateAbsence struct {
	StudentUuid  string `form:"student_uuid"`
	ScheduleUuid string `form:"schedule_uuid"`
}
