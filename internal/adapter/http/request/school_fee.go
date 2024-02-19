package request

type SchoolFee struct {
	Date string `json:"date" valid:"required~field date tidak ditemukan, date~format date harus yyyy-mm-dd"`
	// Nominal        int    `json:"nominal" valid:"required~field nominal tidak ditemukan, numeric~nominal harus angka"`
	Month          string `json:"month" valid:"required~field month tidak ditemukan"`
	StudentUuid    string `json:"student_uuid" valid:"required~field student_uuid tidak ditemukan"`
	SchoolYearUuid string `json:"school_year_uuid" valid:"required~field school_year_uuid tidak ditemukan"`
}
