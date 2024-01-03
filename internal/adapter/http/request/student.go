package request

type UpdateStudent struct {
	Nama         string `json:"nama" valid:"required~field nama tidak ditemukan"`
	NIS          string `json:"nis" valid:"required~field nis tidak ditemukan"`
	Username     string `json:"username" valid:"required~field username tidak ditemukan"`
	JK           string `json:"jk" valid:"required~field jk tidak ditemukan"`
	TempatLahir  string `json:"tempat_lahir" valid:"required~field tempat_lahir tidak ditemukan"`
	TanggalLahir string `json:"tanggal_lahir" valid:"required~field tanggal_lahir tidak ditemukan"`
	Alamat       string `json:"alamat" valid:"required~field alamat tidak ditemukan"`
	ClassID      uint   `json:"class_id" valid:"required~field class_id tidak ditemukan"`
}
