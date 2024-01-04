package application

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
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
		return nil, &response.Error{
			Code:    500,
			Message: "Failed to open file",
		}
	}
	defer f.Close()

	// Get all the rows in the Siswa.
	rows, err := f.GetRows("Siswa")
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Failed to get all rows",
		}
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	var failedStudent []response.FailedStudent

	for i := 1; i < len(rows); i++ {
		cols := rows[i]

		classID, err := strconv.Atoi(cols[7])
		if err != nil {
			failedStudent = append(failedStudent, response.FailedStudent{
				Nama:        cols[7],
				Description: "class ID in column 7 is not a number",
				Error:       err.Error(),
			})
			continue
		}

		user, err := s.Repo.CreateUser(&domain.User{
			Nama:     cols[0],
			Username: cols[1],
			Password: cols[2],
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
			Uuid:         uuid.NewString(),
			NIS:          cols[2],
			JK:           cols[3],
			TempatLahir:  cols[4],
			TanggalLahir: cols[5],
			Alamat:       cols[6],
			ClassID:      uint(classID),
			UserID:       user.ID,
		}

		if err := s.Repo.SaveStudent(&student); err != nil {
			failedStudent = append(failedStudent, response.FailedStudent{
				Nama:        cols[0],
				Description: "failed create student",
				Error:       err.Error(),
			})
			s.Repo.DeleteUser(user)
			continue
		}

	}

	return &failedStudent, nil
}

func (s *StudentService) GetAllStudents() (*[]response.StudentUser, error) {

	students, err := s.Repo.FindAllStudents()
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Failed to get all users",
		}
	}

	var res = []response.StudentUser{}
	for _, student := range *students {
		res = append(res, response.StudentUser{
			Uuid:         student.Uuid,
			Nama:         student.User.Nama,
			Username:     student.User.Username,
			JK:           student.JK,
			NIS:          student.NIS,
			TempatLahir:  student.TempatLahir,
			TanggalLahir: student.TanggalLahir,
			Alamat:       student.Alamat,
			UserID:       student.UserID,
			CreatedAt:    student.CreatedAt,
			UpdatedAt:    student.UpdatedAt,
		})
	}

	return &res, nil
}

func (s *StudentService) StudentsPagination(urlPath string, pagination *domain.Pagination) (*domain.Pagination, error) {

	page, err := s.Repo.PaginationStudents(pagination)
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Failed to get all users: " + err.Error(),
		}
	}

	page.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, page.Limit, 0)
	page.LastPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, page.Limit, page.TotalPages)

	if page.Page > 0 {
		page.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, page.Limit, page.Page-1)
	}

	if page.Page < page.TotalPages {
		page.NextPage = fmt.Sprintf("%s?limit=%d&page=%d", urlPath, page.Limit, page.Page+1)
	}

	if page.Page > page.TotalPages {
		page.PreviousPage = ""
	}



	return page, nil
}

func (s *StudentService) GetStudent(uuid string) (*response.StudentUser, error) {

	student, err := s.Repo.FindStudent(uuid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &response.Error{
			Code:    404,
			Message: fmt.Sprintf("Student with uuid %s is not found", uuid),
		}
	}

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Failed to get student: " + err.Error(),
		}
	}

	res := response.StudentUser{
		Uuid:         student.Uuid,
		Nama:         student.User.Nama,
		Username:     student.User.Username,
		JK:           student.JK,
		NIS:          student.NIS,
		TempatLahir:  student.TempatLahir,
		TanggalLahir: student.TanggalLahir,
		Alamat:       student.Alamat,
		UserID:       student.UserID,
		CreatedAt:    student.CreatedAt,
		UpdatedAt:    student.UpdatedAt,
	}

	return &res, nil
}

func (s *StudentService) UpdateStudent(student *domain.Student, user *domain.User) error {
	studentInDB, err := s.Repo.FindStudent(student.Uuid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Student with uuid %s is not found", student.Uuid),
			}
		}
		return &response.Error{
			Code:    500,
			Message: "Failed to find student: " + err.Error(),
		}
	}

	user.ID = studentInDB.UserID
	student.ID = studentInDB.ID

	if err := s.Repo.UpdateStudent(student, user); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Failed to update student: " + err.Error(),
		}
	}

	return nil
}

func (s *StudentService) DeleteStudent(uuid string) error {

	studentInDB, err := s.Repo.FindStudent(uuid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Student with uuid %s is not found", uuid),
			}
		}
		return &response.Error{
			Code:    500,
			Message: "Failed to find student: " + err.Error(),
		}
	}

	if err := s.Repo.DeleteStudent(studentInDB.UserID); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Failed to delete student: " + err.Error(),
		}
	}

	return nil
}
