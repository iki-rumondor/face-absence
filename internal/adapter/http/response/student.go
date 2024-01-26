package response

import "time"

type StudentResponse struct {
	Nama         string             `json:"nama"`
	Uuid         string             `json:"uuid"`
	NIS          string             `json:"nis"`
	JK           string             `json:"jk"`
	TempatLahir  string             `json:"tempat_lahir"`
	TanggalLahir string             `json:"tanggal_lahir"`
	Alamat       string             `json:"alamat"`
	Class        *ClassData         `json:"class"`
	Absence      *[]AbsenceResponse `json:"absence"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

type TeacherStudents struct {
	Nama         string    `json:"nama"`
	Uuid         string    `json:"uuid"`
	NIS          string    `json:"nis"`
	JK           string    `json:"jk"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir string    `json:"tanggal_lahir"`
	Alamat       string    `json:"alamat"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type FailedStudent struct {
	Nama        string `json:"nama"`
	Description string `json:"description"`
	Error       string `json:"error"`
}
