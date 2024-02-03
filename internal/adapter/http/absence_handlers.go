package customHTTP

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type AbsenceHandler struct {
	Service         *application.AbsenceService
	ScheduleService *application.ScheduleService
	StudentService  *application.StudentService
}

func NewAbsenceHandler(service *application.AbsenceService, schedule *application.ScheduleService, student *application.StudentService) *AbsenceHandler {
	return &AbsenceHandler{
		Service:         service,
		ScheduleService: schedule,
		StudentService:  student,
	}
}

func (h *AbsenceHandler) CreateAbsence(c *gin.Context) {
	var body request.CreateAbsence
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: err.Error(),
		})
		return
	}

	file, err := c.FormFile("face_image")
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "Field face_image tidak ditemukan",
		})
	}

	if ok := utils.IsValidImageExtension(file.Filename); !ok {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "File yang diupload bukan sebuah gambar",
		})
		return
	}

	if ok := utils.IsValidImageSize(file.Size); !ok {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "File maksimal 5MB",
		})
		return
	}

	folder := "internal/assets/temp"
	filename := utils.GenerateRandomFileName(file.Filename)
	pathFile := filepath.Join(folder, filename)

	if err := c.SaveUploadedFile(file, pathFile); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan ketika menyimpan file",
		})
	}

	defer func() {
		if err := os.Remove(pathFile); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if err := h.Service.CreateAbsence(&body, pathFile); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Absensi berhasil disimpan",
	})
}

func (h *AbsenceHandler) GetAllAbsences(c *gin.Context) {
	urlPath := c.Request.URL.Path

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	pagination := domain.Pagination{
		Limit: limit,
		Page:  page,
	}

	result, err := h.Service.GetAllAbsences(urlPath, &pagination)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan data tahun pelajaran",
		Data:    result,
	})
}

func (h *AbsenceHandler) GetStudentAbsences(c *gin.Context) {

	id := c.GetUint("user_id")
	if id == 0 {
		utils.HandleError(c, INTERNAL_ERROR)
		return
	}

	absences, err := h.Service.GetAbsencesUser(id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var res []*response.AbsenceResponse

	for _, item := range *absences {
		res = append(res, &response.AbsenceResponse{
			Uuid:   item.Uuid,
			Status: item.Status,
			Student: &response.StudentResponse{
				Uuid:         item.Student.Uuid,
				JK:           item.Student.JK,
				NIS:          item.Student.NIS,
				TempatLahir:  item.Student.TempatLahir,
				TanggalLahir: item.Student.TanggalLahir,
				Alamat:       item.Student.Alamat,
				TanggalMasuk: item.Student.TanggalMasuk,
				Image:        item.Student.Image,
			},
			Schedule: &response.ScheduleResponse{
				Uuid:  item.Schedule.Uuid,
				Day:   item.Schedule.Day,
				Start: item.Schedule.Start,
				End:   item.Schedule.End,
				Subject: &response.SubjectResponse{
					Uuid:      item.Schedule.Subject.Uuid,
					Name:      item.Schedule.Subject.Name,
					CreatedAt: item.Schedule.Subject.CreatedAt,
					UpdatedAt: item.Schedule.Subject.UpdatedAt,
				},
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Absensi berhasil ditemukan",
		Data:    res,
	})

}

func (h *AbsenceHandler) GetAbsencesPDF(c *gin.Context) {

	dataPDF, err := h.Service.CreateAbsencesPDF()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=absences.pdf")
	c.Data(http.StatusOK, "application/pdf", dataPDF)
}
