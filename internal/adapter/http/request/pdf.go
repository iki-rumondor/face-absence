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
	Semester        string            `json:"semester"`
	Class           string            `json:"kelas"`
	Month           string            `json:"bulan"`
	JumlahHari      int               `json:"jumlah_hari"`
	SchoolYear      string            `json:"tahun_ajaran"`
	StudentsAbsence []StudentsAbsence `json:"siswa"`
}

type StudentsAbsence struct {
	Nama        string    `json:"nama"`
	Absences    []Absence `json:"absensi"`
	JumlahSakit int       `json:"jml_s"`
	JumlahAlpha int       `json:"jml_a"`
	JumlahIzin  int       `json:"jml_i"`
}

type Absence struct {
	Tanggal int    `json:"tanggal"`
	Status  string `json:"absensi"`
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
	Status  string `json:"status"`
}
