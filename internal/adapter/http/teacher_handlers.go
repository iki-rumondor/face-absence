package customHTTP

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type TeacherHandlers struct {
	Service *application.TeacherService
}

func NewTeacherHandler(service *application.TeacherService) *TeacherHandlers {
	return &TeacherHandlers{
		Service: service,
	}
}

func (h *TeacherHandlers) CreateTeacher(c *gin.Context) {
	var body request.CreateTeacher
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

	if err := h.Service.CreateTeacher(body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Berhasil menambah data guru",
	})

}

func (h *TeacherHandlers) GetTeachersPagination(c *gin.Context) {

	urlPath := c.Request.URL.Path

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	pagination := domain.Pagination{
		Limit: limit,
		Page:  page,
	}

	result, err := h.Service.TeachersPagination(urlPath, &pagination)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan data guru",
		Data:    result,
	})

}

func (h *TeacherHandlers) GetTeachers(c *gin.Context) {

	teachers, err := h.Service.GetTeachers()

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var res []*response.Teacher

	for _, teacher := range *teachers {
		res = append(res, &response.Teacher{
			Uuid:          teacher.Uuid,
			JK:            teacher.JK,
			Nip:           teacher.Nip,
			Nuptk:         teacher.Nuptk,
			StatusPegawai: teacher.StatusPegawai,
			TempatLahir:   teacher.TempatLahir,
			TanggalLahir:  teacher.TanggalLahir,
			NoHp:          teacher.NoHp,
			Jabatan:       teacher.Jabatan,
			TotalJtm:      teacher.TotalJtm,
			Alamat:        teacher.Alamat,
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

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan data guru",
		Data:    res,
	})

}

func (h *TeacherHandlers) GetTeacher(c *gin.Context) {

	uuid := c.Param("uuid")
	teacher, err := h.Service.GetTeacher(uuid)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	res := &response.Teacher{
		Uuid:          teacher.Uuid,
		JK:            teacher.JK,
		Nip:           teacher.Nip,
		Nuptk:         teacher.Nuptk,
		StatusPegawai: teacher.StatusPegawai,
		TempatLahir:   teacher.TempatLahir,
		TanggalLahir:  teacher.TanggalLahir,
		NoHp:          teacher.NoHp,
		Jabatan:       teacher.Jabatan,
		TotalJtm:      teacher.TotalJtm,
		Alamat:        teacher.Alamat,
		User: &response.UserData{
			Nama:      teacher.User.Nama,
			Username:  teacher.User.Username,
			Avatar:    teacher.User.Avatar,
			CreatedAt: teacher.User.CreatedAt,
			UpdatedAt: teacher.User.UpdatedAt,
		},
		CreatedAt: teacher.CreatedAt,
		UpdatedAt: teacher.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan data guru",
		Data:    res,
	})

}

func (h *TeacherHandlers) UpdateTeacher(c *gin.Context) {

	var body request.UpdateTeacher
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
	body.Uuid = uuid

	if err := h.Service.UpdateTeacher(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Berhasil memperbarui data guru",
	})

}

func (h *TeacherHandlers) DeleteTeacher(c *gin.Context) {

	uuid := c.Param("uuid")

	if err := h.Service.DeleteTeacher(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Berhasil menghapus data guru",
	})
}

func (h *TeacherHandlers) GetTeachersPDF(c *gin.Context) {

	dataPDF, err := h.Service.CreateTeachersPDF()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=teachers.pdf")
	c.Data(http.StatusOK, "application/pdf", dataPDF)
}
