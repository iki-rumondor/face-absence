package request

type ChangePassword struct {
	NewPassword     string `json:"new_password" valid:"required~field password tidak ditemukan"`
	ConfirmPassword string `json:"confirm_password" valid:"required~field konfirmasi password tidak ditemukan"`
}
