package application

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/xuri/excelize/v2"
)

type StudentService struct {
	Repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) *StudentService {
	return &StudentService{
		Repo: repo,
	}
}

func (s *StudentService) CreateStudent(pathFile string) error {

	f, err := excelize.OpenFile(pathFile)
	if err != nil {
		return err
	}
	defer f.Close()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	var students []*domain.Student

	for _, row := range rows {
		students = append(students, &domain.Student{
			Uuid:  "asdas",
			Nama:  row[1],
			NIM:   row[2],
			Kelas: row[3],
		})
	}

	if err := s.Repo.Save(students); err != nil {
		return err
	}

	return nil
}
