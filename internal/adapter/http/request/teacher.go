package request

type CreateTeacher struct {
	Nuptk         *string `json:"nuptk"`
	Nip           *string `json:"nip"`
	StatusPegawai string  `json:"status_pegawai" valid:"required~field status_pegawai tidak ditemukan"`
	JK            string  `json:"jk" valid:"required~field jk tidak ditemukan"`
	TempatLahir   string  `json:"tempat_lahir" valid:"required~field tempat_lahir tidak ditemukan"`
	TanggalLahir  string  `json:"tanggal_lahir" valid:"required~field nama tidak ditemukan"`
	NoHp          string  `json:"no_hp" valid:"required~field no_hp tidak ditemukan"`
	Jabatan       string  `json:"jabatan" valid:"required~field jabatan tidak ditemukan"`
	TotalJtm      string  `json:"total_jtm" valid:"required~field total_jtm tidak ditemukan"`
	Alamat        string  `json:"alamat" valid:"required~field alamat tidak ditemukan"`
	Nama          string  `json:"nama" valid:"required~field nama tidak ditemukan"`
	Username      string  `json:"username" valid:"required~field username tidak ditemukan"`
}

type UpdateTeacher struct {
	Uuid          string
	Nuptk         *string `json:"nuptk"`
	Nip           *string `json:"nip"`
	StatusPegawai string  `json:"status_pegawai" valid:"required~field status_pegawai tidak ditemukan"`
	JK            string  `json:"jk" valid:"required~field jk tidak ditemukan"`
	TempatLahir   string  `json:"tempat_lahir" valid:"required~field tempat_lahir tidak ditemukan"`
	TanggalLahir  string  `json:"tanggal_lahir" valid:"required~field nama tidak ditemukan"`
	NoHp          string  `json:"no_hp" valid:"required~field no_hp tidak ditemukan"`
	Jabatan       string  `json:"jabatan" valid:"required~field jabatan tidak ditemukan"`
	TotalJtm      string  `json:"total_jtm" valid:"required~field total_jtm tidak ditemukan"`
	Alamat        string  `json:"alamat" valid:"required~field alamat tidak ditemukan"`
	Nama          string  `json:"nama" valid:"required~field nama tidak ditemukan"`
	Username      string  `json:"username" valid:"required~field username tidak ditemukan"`
}
