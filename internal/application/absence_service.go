package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type AbsenceService struct {
	Repo repository.AbsenceRepository
}

func NewAbsenceService(repo repository.AbsenceRepository) *AbsenceService {
	return &AbsenceService{
		Repo: repo,
	}
}

func (s *AbsenceService) CheckStudentIsAbsence(studentID, scheduleID uint) error {
	if result := s.Repo.CheckStudentIsAbsence(studentID, scheduleID); result != 0 {
		return &response.Error{
			Code:    403,
			Message: "Anda Sudah Melakukan Absensi Untuk Jadwal Tersebut",
		}
	}
	return nil
}

func (s *AbsenceService) CheckSchedule(scheduleID uint) (string, error) {
	schedule, err := s.Repo.FindScheduleByID(scheduleID)
	if err != nil {
		return "", &response.Error{
			Code:    404,
			Message: "Jadwal tidak ditemukan",
		}
	}

	if ok := utils.IsTodayEqualTo(schedule.Day); !ok {
		return "", &response.Error{
			Code:    400,
			Message: "Jadwal pelajaran tidak berada di hari ini",
		}
	}

	if ok := utils.IsBeforeTime(schedule.Start); ok {
		return "", &response.Error{
			Code:    400,
			Message: "Jadwal pelajaran belum dimulai",
		}
	}

	status := "HADIR"

	if ok := utils.IsAfterTime(schedule.End); ok {
		status = "TERLAMBAT"
	}

	return status, nil
}

func (s *AbsenceService) CreateAbsence(absence *domain.Absence, faceImage string) error {
	user, err := s.Repo.FindUserByID(absence.Student.UserID)
	if err != nil {
		return &response.Error{
			Code:    404,
			Message: "User tidak ditemukan",
		}
	}

	if user.Avatar == nil || *user.Avatar == "default-avatar.jpg" {
		return &response.Error{
			Code:    404,
			Message: "User tidak memiliki avatar, silahkan upload avatar terlebih dahulu",
		}
	}

	avatarPath := fmt.Sprintf("internal/assets/avatar/%s", *user.Avatar)

	formAbsence, err := s.CreateFormAbsence(avatarPath, faceImage)

	if err != nil {
		return &response.Error{
			Code:    500,
			Message: "Kesalahan dalam sistem, silahkan hubungi developer",
		}
	}

	var FLASK = os.Getenv("FLASK_API")
	if FLASK == "" {
		FLASK = "http://127.0.0.1:5000"
	}

	url := fmt.Sprintf("%s/compare", FLASK)

	res, err := s.CreatePostRequest(url, formAbsence)
	if err != nil {
		return &response.Error{
			Code:    500,
			Message: "Kesalahan dalam sistem, silahkan hubungi developer",
		}
	}

	if !res["matching"] {
		return &response.Error{
			Code:    400,
			Message: "Wajah anda tidak sama dengan yang tersimpan di database",
		}
	}

	if err := s.Repo.CreateAbsence(absence); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Kesalahan dalam sistem, silahkan hubungi developer",
		}
	}

	return nil
}

func (s *AbsenceService) CreateFormAbsence(imageOne, imageTwo string) (*response.FormAbsence, error) {
	// Buat buffer untuk menampung data form
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	maps := []map[string]string{
		{
			"key":   "image1",
			"value": imageOne,
		},
		{
			"key":   "image2",
			"value": imageTwo,
		},
	}

	for _, m := range maps {
		file, err := os.Open(m["value"])
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Tambahkan file ke form
		fileWriter, err := writer.CreateFormFile(m["key"], m["value"])
		if err != nil {
			return nil, err
		}

		// Salin isi file ke writer form
		_, err = io.Copy(fileWriter, file)
		if err != nil {
			return nil, err
		}
	}

	// Selesaikan form
	writer.Close()

	return &response.FormAbsence{
		RequestBody: &requestBody,
		Writer:      writer,
	}, nil
}

func (s *AbsenceService) CreatePostRequest(url string, formAbsence *response.FormAbsence) (map[string]bool, error) {

	// Buat permintaan POST
	request, err := http.NewRequest("POST", url, formAbsence.RequestBody)
	if err != nil {
		return nil, err
	}

	// Tentukan tipe konten form
	request.Header.Set("Content-Type", formAbsence.Writer.FormDataContentType())

	// Kirim permintaan ke server Flask
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Baca dan tampilkan respons
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]bool

	if err := json.Unmarshal(responseBody, &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return data, nil
}

func (s *AbsenceService) GetAllAbsences(urlPath string, pagination *domain.Pagination) (*domain.Pagination, error) {
	result, err := s.Repo.FindAbsencePagination(pagination)
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	page := GeneratePages(urlPath, result)

	return page, nil
}
