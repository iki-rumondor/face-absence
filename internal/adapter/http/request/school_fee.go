package request

type SchoolFee struct {
	Date           string `json:"date" valid:"required~field date tidak ditemukan, date~format date harus yyyy-mm-dd"`
	Status         string `json:"status" valid:"required~field status tidak ditemukan"`
	Month          string `json:"month" valid:"required~field month tidak ditemukan"`
	StudentUuid    string `json:"student_uuid" valid:"required~field student_uuid tidak ditemukan"`
	SchoolYearUuid string `json:"school_year_uuid" valid:"required~field school_year_uuid tidak ditemukan"`
}

type UpdateNominal struct {
	Nominal int `json:"nominal" valid:"required~field nominal tidak ditemukan, numeric~nominal harus angka"`
}
