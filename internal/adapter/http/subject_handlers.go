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

type SubjectHandler struct {
	Service        *application.SubjectService
	TeacherService *application.TeacherService
}

func NewSubjectHandler(service *application.SubjectService, teacher *application.TeacherService) *SubjectHandler {
	return &SubjectHandler{
		Service:        service,
		TeacherService: teacher,
	}
}

func (h *SubjectHandler) CreateSubject(c *gin.Context) {
	var body request.CreateSubject
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
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

	if err := h.Service.CreateSubject(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Berhasil menambah data mata pelajaran",
	})

}

func (h *SubjectHandler) GetSubjectPagination(c *gin.Context) {

	urlPath := c.Request.URL.Path

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	pagination := domain.Pagination{
		Limit: limit,
		Page:  page,
	}

	result, err := h.Service.SubjectPagination(urlPath, &pagination)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan data mata pelajaran",
		Data:    result,
	})

}

func (h *SubjectHandler) GetAllSubjects(c *gin.Context) {

	res, err := h.Service.GetAllSubjects()

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan data mata pelajaran",
		Data:    res,
	})

}

func (h *SubjectHandler) GetSubject(c *gin.Context) {

	uuid := c.Param("uuid")
	subject, err := h.Service.GetSubject(uuid)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	res := response.SubjectResponse{
		Uuid: subject.Uuid,
		Name: subject.Name,
		// Teacher: &response.Teacher{
		// 	Uuid:          subject.Teacher.Uuid,
		// 	JK:            subject.Teacher.JK,
		// 	Nip:           subject.Teacher.Nip,
		// 	Nuptk:         subject.Teacher.Nuptk,
		// 	StatusPegawai: subject.Teacher.StatusPegawai,
		// 	TempatLahir:   subject.Teacher.TempatLahir,
		// 	TanggalLahir:  subject.Teacher.TanggalLahir,
		// 	NoHp:          subject.Teacher.NoHp,
		// 	Jabatan:       subject.Teacher.Jabatan,
		// 	TotalJtm:      subject.Teacher.TotalJtm,
		// 	Alamat:        subject.Teacher.Alamat,
		// 	User: &response.UserData{
		// 		Nama:      subject.Teacher.User.Nama,
		// 		Username:  subject.Teacher.User.Username,
		// 		Avatar:    subject.Teacher.User.Avatar,
		// 		CreatedAt: subject.Teacher.User.CreatedAt,
		// 		UpdatedAt: subject.Teacher.User.UpdatedAt,
		// 	},
		// 	CreatedAt: subject.Teacher.CreatedAt,
		// 	UpdatedAt: subject.Teacher.UpdatedAt,
		// },
		CreatedAt: subject.CreatedAt,
		UpdatedAt: subject.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan data mata pelajaran",
		Data:    res,
	})

}

func (h *SubjectHandler) UpdateSubject(c *gin.Context) {

	var body request.UpdateSubject
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
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

	if err := h.Service.UpdateSubject(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil memperbarui data mata pelajaran",
	})

}

func (h *SubjectHandler) DeleteSubject(c *gin.Context) {

	uuid := c.Param("uuid")
	res, err := h.Service.GetSubject(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	model := domain.Subject{
		Uuid: res.Uuid,
	}

	if err := h.Service.DeleteSubject(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mengahapus data mata pelajaran",
	})

}
