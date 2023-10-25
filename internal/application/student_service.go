package application

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
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

func (s *StudentService) ImportStudents(pathFile string) (*[]response.FailedStudent, error) {

	f, err := excelize.OpenFile(pathFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Get all the rows in the Siswa.
	rows, err := f.GetRows("Siswa")
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	var failedStudent []response.FailedStudent
	fmt.Println(len(rows))

	for i := 1; i < len(rows); i++ {
		cols := rows[i]

		user, err := s.Repo.CreateUser(&domain.User{
			Uuid:     uuid.NewString(),
			Nama:     cols[0],
			Email:    cols[1],
			Password: cols[2],
			RoleID:   1,
		})

		if err != nil {
			failedStudent = append(failedStudent, response.FailedStudent{
				Nama:        cols[0],
				Description: "failed create user",
				Error:       err.Error(),
			})
			continue
		}

		student := domain.Student{
			NIS:      cols[2],
			JK:       cols[3],
			Kelas:    cols[4],
			Semester: cols[5],
			UserID:   user.ID,
		}

		if err := s.Repo.SaveStudent(&student); err != nil {
			failedStudent = append(failedStudent, response.FailedStudent{
				Nama:        cols[0],
				Description: "failed create student",
				Error:       err.Error(),
			})
			continue
		}

	}

	return &failedStudent, nil
}

func (s *StudentService) GetAllStudentUsers() (*[]response.StudentUser, error) {

	users, err := s.Repo.FindAllStudentUsers()
	if err != nil {
		return nil, err
	}

	var students []response.StudentUser
	for _, user := range *users {
		students = append(students, response.StudentUser{
			ID:        user.ID,
			Uuid:      user.Uuid,
			Nama:      user.Nama,
			Email:     user.Email,
			JK:        user.Student.JK,
			NIS:       user.Student.NIS,
			Kelas:     user.Student.Kelas,
			Semester:  user.Student.Semester,
			Role:      user.Role.Name,
			CreatedAt: user.Student.CreatedAt,
			UpdatedAt: user.Student.UpdatedAt,
		})
	}

	return &students, nil
}

func (s *StudentService) GetStudentUser(uuid string) (*response.StudentUser, error) {

	user, err := s.Repo.FindStudentUser(uuid)
	if err != nil {
		return nil, err
	}

	student := response.StudentUser{
		ID:        user.ID,
		Uuid:      user.Uuid,
		Nama:      user.Nama,
		Email:     user.Email,
		JK:        user.Student.JK,
		NIS:       user.Student.NIS,
		Kelas:     user.Student.Kelas,
		Semester:  user.Student.Semester,
		Role:      user.Role.Name,
		CreatedAt: user.Student.CreatedAt,
		UpdatedAt: user.Student.UpdatedAt,
	}

	return &student, nil
}

func (s *StudentService) GetStudent(uuid string) (*domain.Student, error) {
	student, err := s.Repo.FindStudent(uuid)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentService) UpdateStudentUser(req *request.Student) error {

	student := &domain.Student{
		ID:       req.ID,
		NIS:      req.NIS,
		JK:       req.JK,
		Kelas:    req.Kelas,
		Semester: req.Semester,
		UserID:   req.UserID,
	}

	user := &domain.User{
		ID:   req.UserID,
		Nama: req.Nama,
	}

	if err := s.Repo.UpdateStudentUser(student, user); err != nil {
		return err
	}

	return nil
}

func (s *StudentService) DeleteStudentUser(student *domain.Student) error {

	user := &domain.User{
		ID: student.UserID,
	}

	if err := s.Repo.DeleteStudentUser(student, user); err != nil {
		return err
	}

	return nil
}
