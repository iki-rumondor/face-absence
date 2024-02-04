package request

type ClassPDFData struct {
	Name        string `json:"Nama Kelas"`
	TeacherName string `json:"Nama Guru"`
	CreatedAt   string `json:"Dibuat Pada Tanggal"`
}

type StudentPDFData struct {
	Nama         string `json:"Nama"`
	NIS          string `json:"NIS"`
	JK           string `json:"Jenis Kelamin"`
	TempatLahir  string `json:"Tempat Lahir"`
	TanggalLahir string `json:"Tanggal Lahir"`
	Alamat       string `json:"Alamat"`
	Kelas        string `json:"Kelas"`
	TanggalMasuk string `json:"Tanggal Masuk"`
}

type TeacherPDFData struct {
	Nama          string `json:"Nama"`
	Nip           string `json:"NIP"`
	Nuptk         string `json:"NUPTK"`
	StatusPegawai string `json:"Status Pegawai"`
	JK            string `json:"Jenis Kelamin"`
	TempatLahir   string `json:"Tempat Lahir"`
	TanggalLahir  string `json:"Tanggal Lahir"`
	Alamat        string `json:"Alamat"`
	NoHp          string `json:"Nomor HP"`
	Jabatan       string `json:"Jabatan"`
	TotalJtm      string `json:"Total JTM"`
}

type AbsencePDFData struct {
	Date            string            `json:"tanggal"`
	Time            string            `json:"waktu"`
	Class           string            `json:"kelas"`
	Subject         string            `json:"mapel"`
	Teacher         string            `json:"guru"`
	SchoolYear      string            `json:"tahun_ajaran"`
	StudentsAbsence []StudentsAbsence `json:"siswa"`
}

type StudentsAbsence struct {
	Nis        string `json:"nis"`
	Nama       string `json:"nama"`
	Keterangan string `json:"keterangan"`
	Waktu      string `json:"waktu_absensi"`
}

type SchoolFeePDFData struct {
	StudentName   string          `json:"nama"`
	Class         string          `json:"kelas"`
	SchoolFeeData []SchoolFeeData `json:"spp"`
}

type SchoolFeeData struct {
	Date    string `json:"tanggal"`
	Month   string `json:"bulan"`
	Nominal int    `json:"nominal"`
}
