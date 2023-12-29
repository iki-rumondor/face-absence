package request

type CreateSubject struct {
	Name      string `json:"name" valid:"required~field name tidak ditemukan"`
}

type UpdateSubject struct {
	Name      string `json:"name" valid:"required~field name tidak ditemukan"`
}
