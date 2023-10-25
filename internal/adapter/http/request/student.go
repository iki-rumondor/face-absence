package request

type Student struct {
	ID uint
	UserID   uint
	Nama     string `json:"nama" valid:"required~field nama tidak ditemukan"`
	NIS      string `json:"nis" valid:"required~field nis tidak ditemukan"`
	Kelas    string `json:"kelas" valid:"required~field kelas tidak ditemukan"`
	JK       string `json:"jk" valid:"required"`
	Semester string `json:"semester" valid:"required"`
}
