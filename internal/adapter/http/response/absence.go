package response

import (
	"bytes"
	"mime/multipart"
)

type FormAbsence struct {
	RequestBody *bytes.Buffer
	Writer *multipart.Writer
}