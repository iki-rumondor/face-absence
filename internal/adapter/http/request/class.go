package request

type CreateClass struct {
	Name        string `json:"name" valid:"required~field name tidak ditemukan"`
	TeacherUuid string `json:"teacher_uuid" valid:"required~field teacher_uuid tidak ditemukan"`
}

type UpdateClass struct {
	Name        string `json:"name" valid:"required~field name tidak ditemukan"`
	TeacherUuid string `json:"teacher_uuid" valid:"required~field teacher_uuid tidak ditemukan"`
}

