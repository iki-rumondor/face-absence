package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type StudentRepoImplementation struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &StudentRepoImplementation{
		db: db,
	}
}

func (r *StudentRepoImplementation) CreateStudent(student *domain.Student) error {
	return r.db.Create(student).Error
}

func (r *StudentRepoImplementation) PaginationStudents(pagination *domain.Pagination) (*domain.Pagination, error) {
	var students []domain.Student

	var totalPages, fromRow, toRow = 0, 0, 0
	var totalRows int64 = 0

	if err := r.db.Model(&domain.Student{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if pagination.Limit == 0 {
		pagination.Limit = int(totalRows)
	}

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("Class").Find(&students).Error; err != nil {
		return nil, err
	}

	var res = []response.StudentResponse{}
	for _, student := range students {
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

	pagination.Rows = res

	pagination.TotalRows = int(totalRows)

	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit)) - 1)
	pagination.TotalPages = totalPages

	if pagination.Page == 0 {
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = pagination.Page + 1*pagination.Limit
		}
	}

	if toRow > int(totalRows) {
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return pagination, nil
}

func (r *StudentRepoImplementation) FindAllStudents() (*[]domain.Student, error) {
	var students []domain.Student
	if err := r.db.Preload("Class").Find(&students).Error; err != nil {
		return nil, err
	}
	return &students, nil
}

func (r *StudentRepoImplementation) FindStudent(uuid string) (*domain.Student, error) {
	var student domain.Student
	if err := r.db.Preload("Class").First(&student, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepoImplementation) FindStudentByUserID(userID uint) (*domain.Student, error) {
	var student domain.Student
	if err := r.db.Preload("Class").First(&student, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepoImplementation) UpdateStudent(student *domain.Student) error {
	return r.db.Model(student).Updates(student).Error
}

func (r *StudentRepoImplementation) UpdateStudentImage(student *domain.Student, imagePath string) error {
	// var studentFace domain.StudentFace
	// result := r.db.First(&studentFace, "student_id = ?", student.ID).RowsAffected

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(student).Update("image", imagePath).Error; err != nil {
			return err
		}

		// if result > 0 {
		// 	if err := tx.Model(&studentFace).Update("face_encode", faceString).Error; err != nil {
		// 		return err
		// 	}
		// }

		// model := domain.StudentFace{
		// 	FaceEncode: faceString,
		// 	StudentID:  student.ID,
		// }

		// if err := tx.Create(&model).Error; err != nil {
		// 	return err
		// }

		return nil
	})

}

func (r *StudentRepoImplementation) DeleteStudent(student *domain.Student) error {
	return r.db.Delete(student).Error
}

func (r *StudentRepoImplementation) FindLatestHistory() (*domain.PdfDownloadHistory, error) {
	var history domain.PdfDownloadHistory
	if err := r.db.Last(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func (r *StudentRepoImplementation) FindClassBy(column string, value interface{}) (*domain.Class, error) {
	var class domain.Class
	if err := r.db.First(&class, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (r *StudentRepoImplementation) CreatePdfHistory(history *domain.PdfDownloadHistory) error {
	return r.db.Create(history).Error
}

func (r *StudentRepoImplementation) GetStudentsPDF(data []*request.StudentPDFData) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var API_URL = os.Getenv("LARAVEL_API")
	if API_URL == "" {
		API_URL = "http://127.0.0.1:8000/api"
	}

	url := fmt.Sprintf("%s/generate-pdf/Daftar_Santri", API_URL)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *StudentRepoImplementation) CreateBatchStudents(students *[]domain.Student, classUuid string) error {
	var class domain.Class
	if err := r.db.First(&class, "uuid = ?", class.Uuid).Error; err != nil {
		return err
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range *students {
			item.ClassID = class.ID
			if err := tx.Create(item).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *StudentRepoImplementation) GetFaceEncode(pathFile string) (map[string]interface{}, error) {
	var API_URL = os.Getenv("FLASK_API")
	if API_URL == "" {
		API_URL = "http://localhost:5000"
	}

	endpoint := fmt.Sprintf("%s/encode_face", API_URL)

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	file, err := os.Open(filepath.Join("internal/assets/avatar", pathFile))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileWriter, err := writer.CreateFormFile("image", pathFile)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	writer.Close()

	request, err := http.NewRequest("POST", endpoint, &requestBody)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	if err := json.Unmarshal(responseBody, &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return data, nil
}

func (r *StudentRepoImplementation) CheckIsFace(pathFile string) (map[string]interface{}, error) {
	var API_URL = os.Getenv("FLASK_API")
	if API_URL == "" {
		API_URL = "http://localhost:5000"
	}

	endpoint := fmt.Sprintf("%s/check_face", API_URL)

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	file, err := os.Open(filepath.Join("internal/assets/avatar", pathFile))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileWriter, err := writer.CreateFormFile("image", pathFile)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	writer.Close()

	request, err := http.NewRequest("POST", endpoint, &requestBody)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(responseBody))
	var data map[string]interface{}
	if err := json.Unmarshal(responseBody, &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return data, nil
}
