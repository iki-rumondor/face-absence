package request

type CreateSchedule struct {
	Name           string `json:"name" valid:"required~field name tidak ditemukan"`
	Day            string `json:"day" valid:"required~field day tidak ditemukan"`
	Start          string `json:"start" valid:"required~field start tidak ditemukan"`
	End            string `json:"end" valid:"required~field end tidak ditemukan"`
	ClassUuid      string `json:"class_uuid" valid:"required~field class_uuid tidak ditemukan"`
	SubjectUuid    string `json:"subject_uuid" valid:"required~field subject_uuid tidak ditemukan"`
	SchoolYearUuid string `json:"school_year_uuid" valid:"required~field school_year_uuid tidak ditemukan"`
}

type UpdateSchedule struct {
	Name           string `json:"name" valid:"required~field name tidak ditemukan"`
	Day            string `json:"day" valid:"required~field day tidak ditemukan"`
	Start          string `json:"start" valid:"required~field start tidak ditemukan"`
	End            string `json:"end" valid:"required~field end tidak ditemukan"`
	ClassUuid      string `json:"class_uuid" valid:"required~field class_uuid tidak ditemukan"`
	SubjectUuid    string `json:"subject_uuid" valid:"required~field subject_uuid tidak ditemukan"`
	SchoolYearUuid string `json:"school_year_uuid" valid:"required~field school_year_uuid tidak ditemukan"`
}
