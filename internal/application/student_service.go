package application

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/jung-kurt/gofpdf"
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
			Message: "Terjadi kesalahan pada sistem, silahkan hubungi developper",
		}
	}
	defer f.Close()

	// Get all the rows in the Siswa.
	rows, err := f.GetRows("Siswa")
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan pada sistem, silahkan hubungi developper",
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
				Description: "ID kelas pada kolom 7 bukan sebuah angka",
				Error:       err.Error(),
			})
			continue
		}

		user, err := s.Repo.CreateUser(&domain.User{
			Nama:     cols[0],
			Username: cols[1],
			Password: cols[1],
		})

		if err != nil {
			failedStudent = append(failedStudent, response.FailedStudent{
				Nama:        cols[0],
				Description: "Gagal menambah data user",
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
				Description: "Gagal menambah data santri",
				Error:       err.Error(),
			})
			s.Repo.DeleteUser(user)
			continue
		}

	}

	return &failedStudent, nil
}

func (s *StudentService) CreateStudent(student *domain.Student, user *domain.User) error {

	if err := s.Repo.CreateStudentUser(student, user); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Data santri gagal ditambahkan",
		}
	}

	return nil
}

func (s *StudentService) GetAllStudents() (*[]response.StudentUser, error) {

	students, err := s.Repo.FindAllStudents()
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Data santri gagal didapatkan",
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
			Message: "Data santri gagal didapatkan",
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
			Message: fmt.Sprintf("Santri dengan uuid %s tidak ditemukan", uuid),
		}
	}

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Data santri gagal didapatkan",
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
				Message: fmt.Sprintf("Santri dengan uuid %s tidak ditemukan", student.Uuid),
			}
		}
		return &response.Error{
			Code:    500,
			Message: "Gagal mendapatkan data santri",
		}
	}

	user.ID = studentInDB.UserID
	student.ID = studentInDB.ID

	if err := s.Repo.UpdateStudent(student, user); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Gagal memperbarui data santri",
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
				Message: fmt.Sprintf("Santri dengan uuid %s tidak ditemukan", uuid),
			}
		}
		return &response.Error{
			Code:    500,
			Message: "Gagal mendapatkan data santri",
		}
	}

	if err := s.Repo.DeleteStudent(studentInDB.UserID); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return &response.Error{
				Code:    500,
				Message: "Data ini tidak dapat dihapus karena berelasi dengan data lain",
			}
		}
		return &response.Error{
			Code:    500,
			Message: "Gagal menghapus data santri",
		}
	}

	return nil
}

func (s *StudentService) CreateStudentPDF(filePath string) error {
	data, err := s.Repo.FindAllStudents()
	if err != nil {
		return &response.Error{
			Code:    500,
			Message: "Gagal mendapatkan data santri",
		}
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Tambahkan header
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Data Seluruh Santri")

	// Tambahkan data
	pdf.SetFont("Arial", "", 8)

	pdf.Ln(15)
	pdf.SetFillColor(200, 220, 255)
	pdf.SetDrawColor(0, 0, 0)

	type Cell struct {
		Name  string
		Width float64
	}

	headerCells := []Cell{
		{
			Name:  "Nama",
			Width: 50,
		},
		{
			Name:  "NIS",
			Width: 30,
		},
		{
			Name:  "Jenis Kelamin",
			Width: 25,
		},
		{
			Name:  "Tempat, Tanggal Lahir",
			Width: 45,
		},
		{
			Name:  "Alamat",
			Width: 30,
		},
	}

	// Fungsi untuk menambahkan baris data
	addRow := func(cells ...string) {
		for i, cell := range cells {
			pdf.CellFormat(headerCells[i].Width, 10, cell, "1", 0, "", false, 0, "")
		}
		pdf.Ln(10)
	}

	// Tambahkan header
	for _, headerCell := range headerCells {
		pdf.CellFormat(headerCell.Width, 10, headerCell.Name, "1", 0, "", true, 0, "")
	}

	pdf.Ln(10)

	for _, entry := range *data {
		birthInfo := fmt.Sprintf("%s, %s", entry.TempatLahir, entry.TanggalLahir)
		addRow(entry.User.Nama, entry.NIS, entry.JK, birthInfo, entry.Alamat)
	}

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan, silahkan hubungi developper",
		}
	}

	return nil
}

func (s *StudentService) CreatePdfHistory(history *domain.PdfDownloadHistory) error {
	result, err := s.Repo.FindLatestHistory()
	if err == nil {
		if err := os.Remove("internal/assets/temp/" + result.Name); err != nil {
			fmt.Println(err.Error())
		}
	}

	if err := s.Repo.CreatePdfHistory(history); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return nil

}
