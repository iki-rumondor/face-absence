package application

import "github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"

var (
	INTERNAL_ERROR = &response.Error{
		Code:    500,
		Message: "ServiceError: Silahkan Hubungi Developper",
	}
)