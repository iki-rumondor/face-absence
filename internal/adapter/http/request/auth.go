package request

type Login struct {
	Email    string `json:"email" valid:"required~kindly provide your email to continue the journey, email"`
	Password string `json:"password" valid:"required~please make sure to provide password in the request "`
}

type JWT struct{
	Token string `json:"token" valid:"required~the provided JWT token is not valid"`
}