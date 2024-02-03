package customHTTP

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type SchoolFeeHandler struct {
	Service *application.SchoolFeeService
}

func NewSchoolFeeHandler(service *application.SchoolFeeService) *SchoolFeeHandler {
	return &SchoolFeeHandler{
		Service: service,
	}
}

func (h *SchoolFeeHandler) CreateSchoolFee(c *gin.Context) {
	var body request.SchoolFee
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	if err := h.Service.CreateSchoolFee(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "SPP berhasil ditambahkan",
	})
}

func (h *SchoolFeeHandler) GetSchoolFee(c *gin.Context) {
	uuid := c.Param("uuid")
	schoolFee, err := h.Service.GetSchoolFeeByUuid(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	res := response.SchoolFee{
		Uuid:    schoolFee.Uuid,
		Date:    schoolFee.Date,
		Nominal: schoolFee.Nominal,
		Student: &response.StudentResponse{
			Nama:         schoolFee.Student.Nama,
			Uuid:         schoolFee.Student.Uuid,
			JK:           schoolFee.Student.JK,
			NIS:          schoolFee.Student.NIS,
			TempatLahir:  schoolFee.Student.TempatLahir,
			TanggalLahir: schoolFee.Student.TanggalLahir,
			Alamat:       schoolFee.Student.Alamat,
			TanggalMasuk: schoolFee.Student.TanggalMasuk,
			Image:        schoolFee.Student.Image,
			Class: &response.ClassData{
				Uuid:      schoolFee.Student.Class.Uuid,
				Name:      schoolFee.Student.Class.Name,
				CreatedAt: schoolFee.Student.Class.CreatedAt,
				UpdatedAt: schoolFee.Student.Class.UpdatedAt,
			},
			CreatedAt: schoolFee.Student.CreatedAt,
			UpdatedAt: schoolFee.Student.UpdatedAt,
		},
		CreatedAt: schoolFee.CreatedAt,
		UpdatedAt: schoolFee.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan SPP",
		Data:    &res,
	})
}

func (h *SchoolFeeHandler) GetStudentSchoolFee(c *gin.Context) {
	uuid := c.Param("uuid")
	schoolFee, err := h.Service.GetStudentSchoolFee(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	res := response.SchoolFee{
		Uuid:    schoolFee.Uuid,
		Date:    schoolFee.Date,
		Nominal: schoolFee.Nominal,
		Student: &response.StudentResponse{
			Nama:         schoolFee.Student.Nama,
			Uuid:         schoolFee.Student.Uuid,
			JK:           schoolFee.Student.JK,
			NIS:          schoolFee.Student.NIS,
			TempatLahir:  schoolFee.Student.TempatLahir,
			TanggalLahir: schoolFee.Student.TanggalLahir,
			Alamat:       schoolFee.Student.Alamat,
			TanggalMasuk: schoolFee.Student.TanggalMasuk,
			Image:        schoolFee.Student.Image,
			Class: &response.ClassData{
				Uuid:      schoolFee.Student.Class.Uuid,
				Name:      schoolFee.Student.Class.Name,
				CreatedAt: schoolFee.Student.Class.CreatedAt,
				UpdatedAt: schoolFee.Student.Class.UpdatedAt,
			},
			CreatedAt: schoolFee.Student.CreatedAt,
			UpdatedAt: schoolFee.Student.UpdatedAt,
		},
		CreatedAt: schoolFee.CreatedAt,
		UpdatedAt: schoolFee.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan SPP",
		Data:    &res,
	})
}

func (h *SchoolFeeHandler) GetAllSchoolFees(c *gin.Context) {
	urlPath := c.Request.URL.Path

	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))

	schoolFees, pagination, err := h.Service.GetAllSchoolFees(urlPath, page, limit)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var rows []response.SchoolFee

	for _, item := range *schoolFees {
		rows = append(rows, response.SchoolFee{
			Uuid:    item.Uuid,
			Date:    item.Date,
			Nominal: item.Nominal,
			Student: &response.StudentResponse{
				Nama:         item.Student.Nama,
				Uuid:         item.Student.Uuid,
				JK:           item.Student.JK,
				NIS:          item.Student.NIS,
				TempatLahir:  item.Student.TempatLahir,
				TanggalLahir: item.Student.TanggalLahir,
				Alamat:       item.Student.Alamat,
				TanggalMasuk: item.Student.TanggalMasuk,
				Image:        item.Student.Image,
				Class: &response.ClassData{
					Uuid:      item.Student.Class.Uuid,
					Name:      item.Student.Class.Name,
					CreatedAt: item.Student.Class.CreatedAt,
					UpdatedAt: item.Student.Class.UpdatedAt,
				},
				CreatedAt: item.Student.CreatedAt,
				UpdatedAt: item.Student.UpdatedAt,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	pagination["rows"] = rows

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan seluruh kelas",
		Data:    pagination,
	})
}

func (h *SchoolFeeHandler) UpdateSchoolFee(c *gin.Context) {
	var body request.SchoolFee
	if err := c.BindJSON(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	uuid := c.Param("uuid")

	if err := h.Service.UpdateSchoolFee(uuid, &body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mengupdate data SPP",
	})
}

func (h *SchoolFeeHandler) DeleteSchoolFee(c *gin.Context) {
	uuid := c.Param("uuid")

	if err := h.Service.DeleteSchoolFee(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil menghapus data SPP",
	})
}
