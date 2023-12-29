package request

type CreateClass struct {
	Name      string `json:"name" valid:"required~field name tidak ditemukan"`
	TeacherID uint   `json:"teacher_id" valid:"required~field teacher_id tidak ditemukan"`
}

type UpdateClass struct {
	Name      string `json:"name" valid:"required~field name tidak ditemukan"`
	TeacherID uint   `json:"teacher_id" valid:"required~field teacher_id tidak ditemukan"`
}
