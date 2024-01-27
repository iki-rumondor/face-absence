package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type TeacherRepoImplementation struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &TeacherRepoImplementation{
		db: db,
	}
}

func (r *TeacherRepoImplementation) FindTeachersPagination(pagination *domain.Pagination) (*domain.Pagination, error) {
	var teachers []domain.Teacher
	var totalRows int64 = 0

	if err := r.db.Model(&domain.Teacher{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if pagination.Limit == 0 {
		pagination.Limit = int(totalRows)
	}

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("User").Find(&teachers).Error; err != nil {
		return nil, err
	}

	var res = []response.Teacher{}
	for _, teacher := range teachers {
		res = append(res, response.Teacher{
			Uuid:          teacher.Uuid,
			JK:            teacher.JK,
			TempatLahir:   teacher.TempatLahir,
			TanggalLahir:  teacher.TanggalLahir,
			NoHp:          teacher.NoHp,
			Alamat:        teacher.Alamat,
			Nip:           teacher.Nip,
			Nuptk:         teacher.Nuptk,
			StatusPegawai: teacher.StatusPegawai,
			Jabatan:       teacher.Jabatan,
			TotalJtm:      teacher.TotalJtm,
			User: &response.UserData{
				Nama:      teacher.User.Nama,
				Username:  teacher.User.Username,
				Avatar:    teacher.User.Avatar,
				CreatedAt: teacher.User.CreatedAt,
				UpdatedAt: teacher.User.UpdatedAt,
			},
			CreatedAt: teacher.CreatedAt,
			UpdatedAt: teacher.UpdatedAt,
		})
	}

	pagination.Rows = res

	pagination.TotalRows = int(totalRows)

	return pagination, nil
}

func (r *TeacherRepoImplementation) CreateTeacherUser(teacher *domain.Teacher, user *domain.User) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		teacher.UserID = user.ID
		if err := tx.Create(teacher).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *TeacherRepoImplementation) FindTeachers() (*[]domain.Teacher, error) {
	var teachers []domain.Teacher
	if err := r.db.Preload("User").Find(&teachers).Error; err != nil {
		return nil, err
	}

	return &teachers, nil
}

func (r *TeacherRepoImplementation) FindTeacherByUuid(uuid string) (*domain.Teacher, error) {
	var teacher domain.Teacher
	if err := r.db.Preload("User").First(&teacher, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &teacher, nil
}

func (r *TeacherRepoImplementation) FindTeacherByColumn(column, data string) (*domain.Teacher, error) {
	var teacher domain.Teacher
	if err := r.db.First(&teacher, fmt.Sprintf("%s = ?", column), data).Error; err != nil {
		return nil, err
	}

	return &teacher, nil
}

func (r *TeacherRepoImplementation) UpdateTeacherUser(teacher *domain.Teacher, user *domain.User) error {

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(user).Where("id = ?", teacher.UserID).Updates(user).Error; err != nil {
			return err
		}

		if err := tx.Model(teacher).Where("uuid = ?", teacher.Uuid).Updates(teacher).Error; err != nil {
			return err
		}

		return nil
	})

}

func (r *TeacherRepoImplementation) DeleteTeacher(teacherUuid string) error {
	var teacher domain.Teacher
	if err := r.db.Preload("User").First(&teacher, "uuid = ?", teacherUuid).Error; err != nil {
		return err
	}

	return r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&teacher).Association("Subjects").Clear(); err != nil {
			return err
		}

		if err := tx.Delete(&teacher).Error; err != nil {
			return err
		}

		if err := tx.Delete(&teacher.User).Error; err != nil {
			return err
		}

		return nil
	})

}

func (r *TeacherRepoImplementation) FindUserByUsername(username string) (*domain.User, error) {

	var user domain.User
	if err := r.db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *TeacherRepoImplementation) GetTeachersPDF(data []*request.TeacherPDFData) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var API_URL = os.Getenv("LARAVEL_API")
	if API_URL == "" {
		return nil, err
	}

	url := fmt.Sprintf("%s/generate-pdf/Daftar_Guru", API_URL)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
