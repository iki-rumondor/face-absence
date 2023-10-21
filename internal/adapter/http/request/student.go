package request

type Student struct{
	Nama string `json:"nama" valid:"required~field nama tidak ditemukan"`
	NIS string `json:"nis" valid:"required~field nis tidak ditemukan"`
	Kelas string `json:"kelas" valid:"required~field kelas tidak ditemukan"`
	// JK uint	`json:"jk" valid:"required"`
	// Semester uint `json:"semester" valid:"required"`
}