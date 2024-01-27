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
