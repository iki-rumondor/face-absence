package response

import "time"

type Teacher struct {
	Uuid          string    `json:"uuid"`
	Nip           *string    `json:"nip"`
	Nuptk         *string    `json:"nuptk"`
	StatusPegawai string    `json:"status_pegawai"`
	JK            string    `json:"jk"`
	TempatLahir   string    `json:"tempat_lahir"`
	TanggalLahir  string    `json:"tanggal_lahir"`
	NoHp          string    `json:"no_hp"`
	Jabatan       string    `json:"jabatan"`
	TotalJtm      string    `json:"total_jtm"`
	Alamat        string    `json:"alamat"`
	User          *UserData `json:"user"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type TeacherData struct {
	Uuid          string    `json:"uuid"`
	Nip           string   `json:"nip"`
	Nuptk         string   `json:"nuptk"`
	StatusPegawai string    `json:"status_pegawai"`
	JK            string    `json:"jk"`
	TempatLahir   string    `json:"tempat_lahir"`
	TanggalLahir  string    `json:"tanggal_lahir"`
	NoHp          string    `json:"no_hp"`
	Jabatan       string    `json:"jabatan"`
	TotalJtm      string    `json:"total_jtm"`
	Alamat        string    `json:"alamat"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
