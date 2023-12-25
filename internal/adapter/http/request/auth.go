package request

type Login struct {
	Username string `json:"username" valid:"required~please make sure to provide username in the request"`
	Password string `json:"password" valid:"required~please make sure to provide password in the request "`
	Role     string `json:"role" valid:"required~please make sure to provide role in the request "`
}

type JWT struct {
	Token string `json:"token" valid:"required~the provided JWT token is not valid"`
}
