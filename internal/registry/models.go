package registry

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: domain.User{}},
		{Model: domain.Student{}},
		{Model: domain.Teacher{}},
		{Model: domain.Class{}},
		{Model: domain.Subject{}},
		{Model: domain.SchoolYear{}},
		{Model: domain.Schedule{}},
		{Model: domain.Admin{}},
		{Model: domain.Absence{}},
		{Model: domain.PdfDownloadHistory{}},
		{Model: domain.TeacherSubjects{}},
		{Model: domain.SchoolFee{}},
	}
}
