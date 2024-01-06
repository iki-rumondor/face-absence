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

func (s *AbsenceService) CheckSchedule(scheduleID uint) (string, error) {
	schedule, err := s.Repo.FindScheduleByID(scheduleID)
	if err != nil{
		return "", &response.Error{
			Code: 404,
			Message: "Schedule not found",
		}
	}

	if ok := utils.IsTodayEqualTo(schedule.Day); !ok{
		return "", &response.Error{
			Code: 400,
			Message: "Now is not schedule day",
		}
	}

	if ok := utils.IsBeforeTime(schedule.Start); ok{
		return "", &response.Error{
			Code: 400,
			Message: "Now is not schedule time",
		}
	}

	status := "HADIR"

	if ok := utils.IsAfterTime(schedule.End); ok{
		status = "TERLAMBAT"
	}

	return status, nil
}


func (s *AbsenceService) CreateAbsence(absence *domain.Absence, faceImage string) error {
	user, err := s.Repo.FindUserByID(absence.StudentID)
	if err != nil {
		return &response.Error{
			Code:    404,
			Message: "User is not found",
		}
	}

	if user.Avatar == nil {
		return &response.Error{
			Code:    404,
			Message: "User Avatar is not found, please update your avatar",
		}
	}

	avatarPath := fmt.Sprintf("internal/assets/avatar/%s", *user.Avatar)

	formAbsence, err := s.CreateFormAbsence(avatarPath, faceImage)

	if err != nil {
		return &response.Error{
			Code:    500,
			Message: "Failed to create form absence: " + err.Error(),
		}
	}

	url := "http://127.0.0.1:8082/compare"

	res, err := s.CreatePostRequest(url, formAbsence)
	if err != nil {
		return &response.Error{
			Code:    500,
			Message: "Failed to request on face compare service: " + err.Error(),
		}
	}
	
	if !res["matching"]{
		return &response.Error{
			Code:    400,
			Message: "Your face is not matching with face in database",
		}
	}
	
	if err := s.Repo.CreateAbsence(absence); err != nil{
		return &response.Error{
			Code:    500,
			Message: "Failed to create absence: " + err.Error(),
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
			"value": imageOne,
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
	
	fmt.Println(responseBody)

	var data map[string]bool

	if err := json.Unmarshal(responseBody, &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return data, nil
}
