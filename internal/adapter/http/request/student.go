package request

type CreateStudent struct {
	Nama         string `json:"nama" valid:"required~field nama tidak ditemukan"`
	NIS          string `json:"nis" valid:"required~field nis tidak ditemukan"`
	JK           string `json:"jk" valid:"required~field jk tidak ditemukan"`
	TempatLahir  string `json:"tempat_lahir" valid:"required~field tempat_lahir tidak ditemukan"`
	TanggalLahir string `json:"tanggal_lahir" valid:"required~field tanggal_lahir tidak ditemukan"`
	Alamat       string `json:"alamat" valid:"required~field alamat tidak ditemukan"`
	TanggalMasuk string `json:"tanggal_masuk" valid:"required~field tanggal_masuk tidak ditemukan, date~format tanggal DD-MM-YYYY"`
	ClassUuid    string `json:"class_uuid" valid:"required~field class_uuid tidak ditemukan"`
}

type UpdateStudent struct {
	Nama         string `json:"nama" valid:"required~field nama tidak ditemukan"`
	NIS          string `json:"nis" valid:"required~field nis tidak ditemukan"`
	JK           string `json:"jk" valid:"required~field jk tidak ditemukan"`
	TempatLahir  string `json:"tempat_lahir" valid:"required~field tempat_lahir tidak ditemukan"`
	TanggalLahir string `json:"tanggal_lahir" valid:"required~field tanggal_lahir tidak ditemukan"`
	Alamat       string `json:"alamat" valid:"required~field alamat tidak ditemukan"`
	TanggalMasuk string `json:"tanggal_masuk" valid:"required~field tanggal_masuk tidak ditemukan, date~format tanggal DD-MM-YYYY"`
	ClassUuid    string `json:"class_uuid" valid:"required~field class_uuid tidak ditemukan"`
}

type ImportStudents struct {
	ClassUuid string `json:"class_uuid" binding:"required~field class_uuid tidak ditemukan"`
}
