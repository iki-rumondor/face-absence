package request

type SchoolFee struct {
	Date        string `json:"date" valid:"required~field date tidak ditemukan, date~format date harus yyyy-mm-dd"`
	Nominal     int    `json:"nominal" valid:"required~field nominal tidak ditemukan, numeric~nominal harus angka"`
	StudentUuid string `json:"student_uuid" valid:"required~field student_uuid tidak ditemukan"`
}
