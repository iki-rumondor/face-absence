package application

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
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

func (s *StudentService) CreateStudent(student *domain.Student) error {

	if err := s.Repo.CreateStudent(student); err != nil {
		return INTERNAL_ERROR
	}

	return nil
}

func (s *StudentService) GetAllStudents() (*[]response.StudentResponse, error) {

	students, err := s.Repo.FindAllStudents()
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Data santri gagal didapatkan",
		}
	}

	var res = []response.StudentResponse{}
	for _, student := range *students {
		res = append(res, response.StudentResponse{
			Nama:         student.Nama,
			Uuid:         student.Uuid,
			JK:           student.JK,
			NIS:          student.NIS,
			TempatLahir:  student.TempatLahir,
			TanggalLahir: student.TanggalLahir,
			Alamat:       student.Alamat,
			TanggalMasuk: student.TanggalMasuk,
			Image:        student.Image,
			Class: &response.ClassData{
				Uuid:      student.Class.Uuid,
				Name:      student.Class.Name,
				CreatedAt: student.Class.CreatedAt,
				UpdatedAt: student.Class.UpdatedAt,
			},
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
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

func (s *StudentService) GetStudent(uuid string) (*domain.Student, error) {

	student, err := s.Repo.FindStudent(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Santri dengan uuid %s tidak ditemukan", uuid),
			}
		}
		return nil, INTERNAL_ERROR
	}

	return student, nil
}

func (s *StudentService) UpdateStudent(uuid string, body *request.UpdateStudent) error {
	student, err := s.GetStudent(uuid)
	if err != nil {
		return err
	}

	class, err := s.GetClassBy("uuid", body.ClassUuid)
	if err != nil {
		return err
	}

	model := domain.Student{
		ID:           student.ID,
		Nama:         body.Nama,
		NIS:          body.NIS,
		JK:           body.JK,
		TempatLahir:  body.TempatLahir,
		TanggalLahir: body.TanggalLahir,
		TanggalMasuk: body.TanggalMasuk,
		Alamat:       body.Alamat,
		ClassID:      class.ID,
	}

	if err := s.Repo.UpdateStudent(&model); err != nil {
		return INTERNAL_ERROR
	}

	return nil
}

func (s *StudentService) UpdateStudentImage(uuid string, imagePath string) error {

	if err := s.Repo.UpdateStudentImage(uuid, imagePath); err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Santri dengan uuid %s tidak ditemukan", uuid),
			}
		}
		return INTERNAL_ERROR
	}

	return nil
}

func (s *StudentService) DeleteStudent(uuid string) error {

	student, err := s.GetStudent(uuid)
	if err != nil {
		return err
	}

	model := domain.Student{
		ID: student.ID,
	}

	if err := s.Repo.DeleteStudent(&model); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return &response.Error{
				Code:    403,
				Message: "Data ini tidak dapat dihapus karena berelasi dengan data lain",
			}
		}
		return INTERNAL_ERROR
	}

	return nil
}

func (s *StudentService) CreateStudentsPDF() ([]byte, error) {

	students, err := s.GetAllStudents()
	if err != nil {
		return nil, err
	}

	if len(*students) == 0 {
		return nil, &response.Error{
			Code:    404,
			Message: "Data Kelas Masih Kosong",
		}
	}

	var data []*request.StudentPDFData

	for _, item := range *students {
		data = append(data, &request.StudentPDFData{
			Nama:         item.Nama,
			NIS:          item.NIS,
			JK:           item.JK,
			TempatLahir:  item.TempatLahir,
			TanggalLahir: item.TanggalLahir,
			Alamat:       item.Alamat,
			TanggalMasuk: item.TanggalMasuk,
			Kelas:        item.Class.Name,
		})
	}

	resp, err := s.Repo.GetStudentsPDF(data)
	if err != nil {
		log.Println(err.Error())
		return nil, INTERNAL_ERROR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, INTERNAL_ERROR
	}

	return pdfData, nil
}

func (s *StudentService) ImportStudents(pathFile string, body *request.ImportStudents) error {

	f, err := excelize.OpenFile(pathFile)
	if err != nil {
		log.Println("Failed to open file")
		return INTERNAL_ERROR
	}
	defer f.Close()

	rows, err := f.GetRows("Siswa")
	if err != nil {
		log.Println("Failed to get rows Siswa")
		return INTERNAL_ERROR
	}

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered from panic:", r)
	// 	}
	// }()

	var students []domain.Student

	for i := 1; i < len(rows); i++ {
		cols := rows[i]

		students = append(students, domain.Student{
			Uuid:         uuid.NewString(),
			Nama:         cols[0],
			NIS:          cols[1],
			JK:           cols[2],
			TempatLahir:  cols[3],
			TanggalLahir: cols[4],
			Alamat:       cols[5],
			TanggalMasuk: cols[6],
		})
	}

	if err := s.Repo.CreateBatchStudents(&students, body.ClassUuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.Error{
				Code:    404,
				Message: "Kelas Tidak Ditemukan",
			}
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return &response.Error{
				Code:    400,
				Message: "Terdeteksi Duplikasi Data, Periksa Kembali NIS Siswa",
			}
		}
		return INTERNAL_ERROR
	}

	return nil
}

// func (s *StudentService) CreateStudentPDF(filePath string) error {
// 	data, err := s.Repo.FindAllStudents()
// 	if err != nil {
// 		return &response.Error{
// 			Code:    500,
// 			Message: "Gagal mendapatkan data santri",
// 		}
// 	}

// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.AddPage()

// 	// Tambahkan header
// 	pdf.SetFont("Arial", "B", 16)
// 	pdf.Cell(40, 10, "Data Seluruh Santri")

// 	// Tambahkan data
// 	pdf.SetFont("Arial", "", 8)

// 	pdf.Ln(15)
// 	pdf.SetFillColor(200, 220, 255)
// 	pdf.SetDrawColor(0, 0, 0)

// 	type Cell struct {
// 		Name  string
// 		Width float64
// 	}

// 	headerCells := []Cell{
// 		{
// 			Name:  "Nama",
// 			Width: 50,
// 		},
// 		{
// 			Name:  "NIS",
// 			Width: 30,
// 		},
// 		{
// 			Name:  "Jenis Kelamin",
// 			Width: 25,
// 		},
// 		{
// 			Name:  "Tempat, Tanggal Lahir",
// 			Width: 45,
// 		},
// 		{
// 			Name:  "Alamat",
// 			Width: 30,
// 		},
// 	}

// 	// Fungsi untuk menambahkan baris data
// 	addRow := func(cells ...string) {
// 		for i, cell := range cells {
// 			pdf.CellFormat(headerCells[i].Width, 10, cell, "1", 0, "", false, 0, "")
// 		}
// 		pdf.Ln(10)
// 	}

// 	// Tambahkan header
// 	for _, headerCell := range headerCells {
// 		pdf.CellFormat(headerCell.Width, 10, headerCell.Name, "1", 0, "", true, 0, "")
// 	}

// 	pdf.Ln(10)

// 	for _, entry := range *data {
// 		birthInfo := fmt.Sprintf("%s, %s", entry.TempatLahir, entry.TanggalLahir)
// 		addRow(entry.User.Nama, entry.NIS, entry.JK, birthInfo, entry.Alamat)
// 	}

// 	if err := pdf.OutputFileAndClose(filePath); err != nil {
// 		return &response.Error{
// 			Code:    500,
// 			Message: "Terjadi kesalahan, silahkan hubungi developper",
// 		}
// 	}

// 	return nil
// }

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

func (s *StudentService) GetClassBy(column string, value interface{}) (*domain.Class, error) {
	class, err := s.Repo.FindClassBy(column, value)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Kelas dengan uuid %s tidak ditemukan", value),
			}
		}
		return nil, INTERNAL_ERROR
	}

	return class, nil

}
