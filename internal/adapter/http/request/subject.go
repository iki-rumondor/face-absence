package request

type CreateSubject struct {
	Name         string   `json:"name" valid:"required~field name tidak ditemukan"`
	TeachersUuid []string `json:"teachers_uuid" valid:"required~field teacher_uuid tidak ditemukan"`
}

type UpdateSubject struct {
	Name string `json:"name" valid:"required~field name tidak ditemukan"`
	TeachersUuid []string `json:"teachers_uuid" valid:"required~field teacher_uuid tidak ditemukan"`
}
