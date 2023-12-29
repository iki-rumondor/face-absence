package request

type CreateSchedule struct {
	Name         string `json:"name" valid:"required~field name tidak ditemukan"`
	Day          string `json:"day" valid:"required~field day tidak ditemukan"`
	Start        string `json:"start" valid:"required~field start tidak ditemukan"`
	End          string `json:"end" valid:"required~field end tidak ditemukan"`
	ClassID      uint   `json:"class_id" valid:"required~field class_id tidak ditemukan"`
	SubjectID    uint   `json:"subject_id" valid:"required~field subject_id tidak ditemukan"`
	TeacherID    uint   `json:"teacher_id" valid:"required~field teacher_id tidak ditemukan"`
	SchoolYearID uint   `json:"school_year_id" valid:"required~field school_year_id tidak ditemukan"`
}

type UpdateSchedule struct {
	Name         string `json:"name" valid:"required~field name tidak ditemukan"`
	Day          string `json:"day" valid:"required~field day tidak ditemukan"`
	Start        string `json:"start" valid:"required~field start tidak ditemukan"`
	End          string `json:"end" valid:"required~field end tidak ditemukan"`
	ClassID      uint   `json:"class_id" valid:"required~field class_id tidak ditemukan"`
	SubjectID    uint   `json:"subject_id" valid:"required~field subject_id tidak ditemukan"`
	TeacherID    uint   `json:"teacher_id" valid:"required~field teacher_id tidak ditemukan"`
	SchoolYearID uint   `json:"school_year_id" valid:"required~field school_year_id tidak ditemukan"`
}
