package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type ClassRepoImplementation struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &ClassRepoImplementation{
		db: db,
	}
}

func (r *ClassRepoImplementation) FindClassPagination(pagination *domain.Pagination) (*domain.Pagination, error) {
	var classes []domain.Class
	var totalRows int64 = 0

	if err := r.db.Model(&domain.Class{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if pagination.Limit == 0 {
		pagination.Limit = int(totalRows)
	}

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("Teacher.User").Find(&classes).Error; err != nil {
		return nil, err
	}

	var res = []response.ClassResponse{}
	for _, class := range classes {
		res = append(res, response.ClassResponse{
			Uuid: class.Uuid,
			Name: class.Name,
			Teacher: &response.Teacher{
				Uuid:          class.Teacher.Uuid,
				JK:            class.Teacher.JK,
				Nip:           class.Teacher.Nip,
				Nuptk:         class.Teacher.Nuptk,
				StatusPegawai: class.Teacher.StatusPegawai,
				TempatLahir:   class.Teacher.TempatLahir,
				TanggalLahir:  class.Teacher.TanggalLahir,
				NoHp:          class.Teacher.NoHp,
				Jabatan:       class.Teacher.Jabatan,
				TotalJtm:      class.Teacher.TotalJtm,
				Alamat:        class.Teacher.Alamat,
				User: &response.UserData{
					Nama:      class.Teacher.User.Nama,
					Username:  class.Teacher.User.Username,
					Avatar:    class.Teacher.User.Avatar,
					CreatedAt: class.Teacher.User.CreatedAt,
					UpdatedAt: class.Teacher.User.UpdatedAt,
				},
				CreatedAt: class.Teacher.CreatedAt,
				UpdatedAt: class.Teacher.UpdatedAt,
			},
			CreatedAt: class.CreatedAt,
			UpdatedAt: class.UpdatedAt,
		})
	}

	pagination.Rows = res

	pagination.TotalRows = int(totalRows)

	return pagination, nil
}

func (r *ClassRepoImplementation) CreateClass(class *domain.Class) error {
	return r.db.Create(class).Error
}

func (r *ClassRepoImplementation) UpdateClass(class *domain.Class) error {
	return r.db.Model(class).Where("uuid = ?", class.Uuid).Updates(class).Error
}

func (r *ClassRepoImplementation) FindClasses() (*[]domain.Class, error) {
	var classes []domain.Class
	if err := r.db.Preload("Teacher.User").Find(&classes).Error; err != nil {
		return nil, err
	}

	return &classes, nil
}

func (r *ClassRepoImplementation) FindClassByUuid(uuid string) (*domain.Class, error) {
	var class domain.Class
	if err := r.db.Preload("Teacher.User").First(&class, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (r *ClassRepoImplementation) DeleteClass(class *domain.Class) error {
	return r.db.Delete(&class, "uuid = ?", class.Uuid).Error
}

func (r *ClassRepoImplementation) GetClassPDF(data []*request.ClassPDFData) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var API_URL = os.Getenv("LARAVEL_API")
	if API_URL == "" {
		return nil, err
	}

	url := fmt.Sprintf("%s/generate-pdf/Daftar_Seluruh_Kelas", API_URL)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return pdfData, nil
}

func (r *ClassRepoImplementation) FindClassBy(column string, value interface{}) (*domain.Class, error) {
	var class domain.Class
	if err := r.db.Preload("Students").First(&class, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (r *ClassRepoImplementation) FindTeacherClassesByUserID(userID uint) (*domain.Teacher, error) {
	var teacher domain.Teacher
	if err := r.db.Preload("Classes.Students").First(&teacher, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *ClassRepoImplementation) FindTeacherClass(userID uint, classUuid string) (*domain.Class, error) {
	
	var teacher domain.Teacher
	if err := r.db.First(&teacher, "user_id = ?", userID).Error; err != nil {
		log.Println("Teacher with user id not found")
		return nil, err
	}

	var class domain.Class
	if err := r.db.Preload("Students").Preload("Teacher").First(&class, "teacher_id = ? AND uuid = ?", teacher.ID, classUuid).Error; err != nil{
		return nil, err
	}

	return &class, nil
}


