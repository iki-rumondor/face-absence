package customHTTP

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type StudentHandlers struct {
	Service      *application.StudentService
	ClassService *application.ClassService
}

func NewStudentHandler(service *application.StudentService, class *application.ClassService) *StudentHandlers {
	return &StudentHandlers{
		Service:      service,
		ClassService: class,
	}
}

func (h *StudentHandlers) CreateStudent(c *gin.Context) {
	var body request.CreateStudent
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: err.Error(),
		})
		return
	}

	class, err := h.ClassService.GetClass(body.ClassUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	student := domain.Student{
		Nama:         body.Nama,
		Uuid:         uuid.NewString(),
		NIS:          body.NIS,
		JK:           body.JK,
		TempatLahir:  body.TempatLahir,
		TanggalLahir: body.TanggalLahir,
		TanggalMasuk: body.TanggalMasuk,
		Alamat:       body.Alamat,
		ClassID:      class.ID,
	}

	if err := h.Service.CreateStudent(&student); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Santri berhasil ditambahkan",
	})
}

func (h *StudentHandlers) GetAllStudentsData(c *gin.Context) {

	urlPath := c.Request.URL.Path

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	pagination := domain.Pagination{
		Limit: limit,
		Page:  page,
	}

	result, err := h.Service.StudentsPagination(urlPath, &pagination)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan data santri",
		Data:    result,
	})
}

func (h *StudentHandlers) GetStudentData(c *gin.Context) {

	student, err := h.Service.GetStudent(c.Param("uuid"))

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var school_fees []response.SchoolFee

	for _, item := range *student.SchoolFees {
		school_fees = append(school_fees, response.SchoolFee{
			Uuid:    item.Uuid,
			Date:    item.Date.Format("02-01-2006"),
			Nominal: item.Nominal,
			Month:   item.Month,
			Status:  item.Status,
			SchoolYear: &response.SchoolYearResponse{
				Uuid: item.SchoolYear.Uuid,
				Name: item.SchoolYear.Name,
			},
		})
	}

	res := response.StudentResponse{
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
		SchoolFees: &school_fees,
		CreatedAt:  student.CreatedAt,
		UpdatedAt:  student.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: fmt.Sprintf("Santri dengan uuid %s tidak ditemukan", student.Uuid),
		Data:    res,
	})
}

func (h *StudentHandlers) UpdateStudentData(c *gin.Context) {
	var body request.UpdateStudent
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: err.Error(),
		})
		return
	}

	uuid := c.Param("uuid")

	if err := h.Service.UpdateStudent(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data santri berhasil diperbarui",
	})
}

func (h *StudentHandlers) DeleteStudent(c *gin.Context) {

	uuid := c.Param("uuid")

	if err := h.Service.DeleteStudent(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Data santri berhasil dihapus",
	})
}

func (h *StudentHandlers) UpdateStudentImage(c *gin.Context) {

	uuid := c.Param("uuid")

	imagePath, err := utils.SaveUploadedImage(c)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	if err := h.Service.UpdateStudentImage(uuid, imagePath); err != nil {
		if err := os.Remove("internal/assets/avatar/" + imagePath); err != nil {
			log.Println(err.Error())
		}
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Image santri berhasil diperbarui",
	})
}

func (h *StudentHandlers) GetStudentsPDF(c *gin.Context) {
	dataPDF, err := h.Service.CreateStudentsPDF()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=students.pdf")
	c.Data(http.StatusOK, "application/pdf", dataPDF)
}

func (h *StudentHandlers) ImportStudentsData(c *gin.Context) {
	var body request.ImportStudents
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	file, err := c.FormFile("students")
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "File students tidak ditemukan",
		})
		return
	}

	if err := utils.IsExcelFile(file); err != nil {
		utils.HandleError(c, err)
		return
	}

	tempFolder := "internal/assets/temp"
	pathFile := filepath.Join(tempFolder, file.Filename)

	if err := c.SaveUploadedFile(file, pathFile); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan saat menyimpan file",
		})
	}

	defer func() {
		if err := os.Remove(pathFile); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if err := h.Service.ImportStudents(pathFile, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Santri berhasil ditambahkan",
	})
}

// func (h *StudentHandlers) CreateReport(c *gin.Context) {
// 	randomName := fmt.Sprintf("%s.pdf", uuid.NewString())
// 	filePath := fmt.Sprintf("internal/assets/temp/%s", randomName)

// 	if err := h.Service.CreateStudentPDF(filePath); err != nil {
// 		utils.HandleError(c, err)
// 		return
// 	}

// 	history := domain.PdfDownloadHistory{
// 		Name: randomName,
// 	}

// 	if err := h.Service.CreatePdfHistory(&history); err != nil {
// 		utils.HandleError(c, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, response.SuccessResponse{
// 		Success: true,
// 		Message: fmt.Sprintf("/public/file/%s", randomName),
// 	})

// }
