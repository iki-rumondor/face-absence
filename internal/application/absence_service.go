package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/gorm"
)

type AbsenceService struct {
	Repo repository.AbsenceRepository
}

func NewAbsenceService(repo repository.AbsenceRepository) *AbsenceService {
	return &AbsenceService{
		Repo: repo,
	}
}

func (s *AbsenceService) CheckSchedule(schedule *domain.Schedule) (string, error) {

	if ok := utils.IsDayEqualTo(schedule.Day); !ok {
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

func (s *AbsenceService) CreateAbsence(req *request.CreateAbsence, faceImage string) error {

	student, err := s.Repo.FindStudentByUuid(req.StudentUuid)
	if err != nil {
		return &response.Error{
			Code:    404,
			Message: "Santri tidak ditemukan",
		}
	}

	schedule, err := s.Repo.FindScheduleByUuid(req.ScheduleUuid)
	if err != nil {
		return &response.Error{
			Code:    404,
			Message: "Jadwal tidak ditemukan",
		}
	}

	if result := s.Repo.CheckStudentIsAbsence(student.ID, schedule.ID); result > 0 {
		return &response.Error{
			Code:    403,
			Message: "Anda Sudah Melakukan Absensi Untuk Jadwal Tersebut",
		}
	}

	status, err := s.CheckSchedule(schedule)
	if err != nil {
		return err
	}

	absence := domain.Absence{
		Uuid:       uuid.NewString(),
		StudentID:  student.ID,
		ScheduleID: schedule.ID,
		Student:    student,
		Status:     status,
	}

	if student.Image == "default-avatar.jpg" {
		return &response.Error{
			Code:    404,
			Message: "Santri masih menggunakan default avatar",
		}
	}

	imagePath := fmt.Sprintf("internal/assets/avatar/%s", student.Image)

	formAbsence, err := s.CreateFormAbsence(imagePath, faceImage)

	if err != nil {
		return &response.Error{
			Code:    500,
			Message: "Gagal Terhubung Dengan Face Recognition",
		}
	}

	var FLASK = os.Getenv("FLASK_API")
	if FLASK == "" {
		FLASK = "http://localhost:5000"
	}

	url := fmt.Sprintf("%s/compare", FLASK)

	res, err := s.CreatePostRequest(url, formAbsence)
	if err != nil {
		return &response.Error{
			Code:    500,
			Message: "Gagal Terhubung Dengan Face Recognition",
		}
	}

	log.Println(res)

	if res["success"] == false {
		return &response.Error{
			Code:    400,
			Message: res["message"].(string),
		}
	}

	if !res["matching"].(bool) {
		return &response.Error{
			Code:    400,
			Message: "Wajah anda tidak sama dengan yang tersimpan di database",
		}
	}

	if err := s.Repo.CreateAbsence(&absence); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Gagal Menyimpan Absensi Dalam Database",
		}
	}

	return nil
}

func (s *AbsenceService) CreateFormAbsence(imageOne, imageTwo string) (*response.FormAbsence, error) {
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

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			return nil, err
		}
	}

	writer.Close()

	return &response.FormAbsence{
		RequestBody: &requestBody,
		Writer:      writer,
	}, nil
}

func (s *AbsenceService) CreatePostRequest(url string, formAbsence *response.FormAbsence) (map[string]interface{}, error) {

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

	var data map[string]interface{}

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

func (s *AbsenceService) GetAbsencesUser(userID uint) (*[]domain.Absence, error) {

	student, err := s.Repo.FindStudentByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: "User tidak ditemukan",
			}
		}
		return nil, INTERNAL_ERROR
	}

	result, err := s.Repo.FindAbsencesStudent(student.ID)
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return result, nil

}

func (s *AbsenceService) CreateAbsencesPDF(scheduleUuid, schoolYearUuid string, month int) ([]byte, error) {
	schedule, err := s.Repo.FindScheduleByUuid(scheduleUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: "Jadwal Tidak Ditemukan",
			}
		}
		return nil, INTERNAL_ERROR
	}

	year, err := s.Repo.FindSchoolYear(schoolYearUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: "Tahun Ajaran Tidak Ditemukan",
			}
		}
		return nil, INTERNAL_ERROR
	}

	var semester = "Ganjil"
	if month >= 1 && month <= 6 {
		semester = "Genap"
	}

	parts := strings.Split(year.Name, "/")
	if len(parts) != 2 {
		return nil, INTERNAL_ERROR
	}

	yearName := parts[0]
	if semester == "Genap" {
		yearName = parts[1]
	}

	yearInt, err := strconv.Atoi(yearName)
	if err != nil {
		return nil, INTERNAL_ERROR
	}

	absences, err := s.Repo.FindAbsenceByYearMonth(schedule.ID, yearInt, month)
	if err != nil {
		return nil, INTERNAL_ERROR
	}

	jumlahHari := utils.JumlahHariPadaBulan(month, yearInt)

	var students []request.StudentsAbsence
	for _, item := range *schedule.Class.Students {
		var studentAbs []request.Absence
		var JumlahAlpha, JumlahIzin, JumlahSakit int
		for i := 1; i <= jumlahHari; i++ {

			status := "TANPA KETERANGAN"
			for _, abs := range *absences {
				log.Println("day :", abs.CreatedAt.Day())
				if item.ID == abs.StudentID && i == abs.CreatedAt.Day() {
					if abs.Status == "HADIR" {
						status = abs.Status
						break
					}
					if abs.Status == "SAKIT" {
						status = abs.Status
						break
					}
					if abs.Status == "IZIN" {
						status = abs.Status
						break
					}
					if abs.Status == "ALPA" {
						status = abs.Status
						break
					}
				}
			}
			if status == "ALPA" {
				JumlahAlpha++
			}

			if status == "SAKIT" {
				JumlahSakit++
			}

			if status == "IZIN" {
				JumlahIzin++
			}

			studentAbs = append(studentAbs, request.Absence{
				Tanggal: i,
				Status:  status,
			})
		}
		students = append(students, request.StudentsAbsence{
			Nama:        item.Nama,
			Absences:    studentAbs,
			JumlahSakit: JumlahSakit,
			JumlahIzin:  JumlahIzin,
			JumlahAlpha: JumlahAlpha,
		})
	}

	var data = request.AbsencePDFData{
		Subject:         schedule.Subject.Name,
		Semester:        semester,
		JumlahHari:      jumlahHari,
		Month:           utils.GetBulanIndonesia(fmt.Sprintf("%02d", month)),
		Class:           schedule.Class.Name,
		SchoolYear:      year.Name,
		StudentsAbsence: students,
	}

	resp, err := s.Repo.GetAbsencesPDF(&data)
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
