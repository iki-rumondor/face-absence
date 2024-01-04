package response

import "time"

type StudentUser struct {
	Uuid         string    `json:"uuid"`
	Nama         string    `json:"nama"`
	Username     string    `json:"username"`
	NIS          string    `json:"nis"`
	JK           string    `json:"jk"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir string    `json:"tanggal_lahir"`
	Alamat       string    `json:"alamat"`
	UserID       uint      `json:"user_id"`
	ClassID      uint      `json:"class_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type FailedStudent struct {
	Nama        string `json:"nama"`
	Description string `json:"description"`
	Error       string `json:"error"`
}
