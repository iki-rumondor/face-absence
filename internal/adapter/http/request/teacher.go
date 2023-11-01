package request

type CreateTeacher struct {
	Nama string `json:"nama" valid:"required~field nama tidak ditemukan"`
	Email string `json:"email" valid:"required~field email tidak ditemukan, email"`
	NIP  string `json:"nip" valid:"required~field nip tidak ditemukan"`
	JK   string `json:"jenis_kelamin" valid:"required~field jenis kelamin tidak ditemukan"`
}
