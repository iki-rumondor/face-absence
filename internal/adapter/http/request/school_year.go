package request

type CreateSchoolYear struct {
	Name      string `json:"name" valid:"required~field name tidak ditemukan"`
}

type UpdateSchoolYear struct {
	Name      string `json:"name" valid:"required~field name tidak ditemukan"`
}
