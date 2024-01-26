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

	var resp []response.SubjectResponse

	for _, res := range *res {
		var teachers []response.Teacher
		for _, item := range res.Teachers {
			teachers = append(teachers, response.Teacher{
				Uuid:          item.Uuid,
				JK:            item.JK,
				Nip:           item.Nip,
				Nuptk:         item.Nuptk,
				StatusPegawai: item.StatusPegawai,
				TempatLahir:   item.TempatLahir,
				TanggalLahir:  item.TanggalLahir,
				NoHp:          item.NoHp,
				Jabatan:       item.Jabatan,
				TotalJtm:      item.TotalJtm,
				Alamat:        item.Alamat,
				User: &response.UserData{
					Nama:      item.User.Nama,
					Username:  item.User.Username,
					Avatar:    item.User.Avatar,
					CreatedAt: item.User.CreatedAt,
					UpdatedAt: item.User.UpdatedAt,
				},
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			})
		}

		resp = append(resp, response.SubjectResponse{
			Uuid:      res.Uuid,
			Name:      res.Name,
			Teachers:  &teachers,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan data mata pelajaran",
		Data:    resp,
	})

}

func (h *SubjectHandler) GetSubject(c *gin.Context) {

	uuid := c.Param("uuid")

	subject, err := h.Service.GetSubject(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var teachers []response.Teacher
	for _, item := range subject.Teachers {
		teachers = append(teachers, response.Teacher{
			Uuid:          item.Uuid,
			JK:            item.JK,
			Nip:           item.Nip,
			Nuptk:         item.Nuptk,
			StatusPegawai: item.StatusPegawai,
			TempatLahir:   item.TempatLahir,
			TanggalLahir:  item.TanggalLahir,
			NoHp:          item.NoHp,
			Jabatan:       item.Jabatan,
			TotalJtm:      item.TotalJtm,
			Alamat:        item.Alamat,
			User: &response.UserData{
				Nama:      item.User.Nama,
				Username:  item.User.Username,
				Avatar:    item.User.Avatar,
				CreatedAt: item.User.CreatedAt,
				UpdatedAt: item.User.UpdatedAt,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	res := response.SubjectResponse{
		Uuid:      subject.Uuid,
		Name:      subject.Name,
		Teachers:  &teachers,
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
