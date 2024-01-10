package response

import "time"

type StudentResponse struct {
	Uuid         string     `json:"uuid"`
	NIS          string     `json:"nis"`
	JK           string     `json:"jk"`
	TempatLahir  string     `json:"tempat_lahir"`
	TanggalLahir string     `json:"tanggal_lahir"`
	Alamat       string     `json:"alamat"`
	User         *UserData  `json:"user"`
	Class        *ClassData `json:"class"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type FailedStudent struct {
	Nama        string `json:"nama"`
	Description string `json:"description"`
	Error       string `json:"error"`
}
