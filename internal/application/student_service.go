package application

import (
	"github.com/google/uuid"
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

	var students domain.ListOfStudent

	for _, row := range rows {
		students.Students = append(students.Students, domain.Student{
			Uuid:  uuid.NewString(),
			Nama:  row[1],
			NIS:   row[2],
			Kelas: row[3],
		})
	}

	if err := s.Repo.SaveList(&students); err != nil {
		return err
	}

	return nil
}

func (s *StudentService) GetAllStudents() (*domain.ListOfStudent, error) {

	students, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *StudentService) GetStudent(uuid string) (*domain.Student, error) {

	student, err := s.Repo.Find(uuid)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (s *StudentService) UpdateStudent(student *domain.Student) error {

	if err := s.Repo.Save(student); err != nil {
		return err
	}

	return nil
}

func (s *StudentService) DeleteStudent(student *domain.Student) error {

	if err := s.Repo.Delete(student); err != nil {
		return err
	}

	return nil
}
