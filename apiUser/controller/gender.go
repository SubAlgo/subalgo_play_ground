package controller

import (
	"apiUser/repository"
	"apiUser/services"
	"errors"
	"log"
	"net/http"

	reqPackage "apiUser/request"
	resPackage "apiUser/response"

	"github.com/gin-gonic/gin"
)

type ControllerGender interface {
	Create(c *gin.Context)
	GetRow(c *gin.Context)
	GetRows(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func NewControllerGender() ControllerGender {
	return &handlerGender{
		serviceGender: services.NewServiceGender(),
	}
}

type handlerGender struct {
	serviceGender services.ServiceGender
}

func (h *handlerGender) Create(c *gin.Context) {
	var req reqPackage.RequestGender
	if err := c.BindJSON(&req); err != nil {
		log.Printf("error: handler_gendere_create_bind_request -> %v\n", err)
		c.JSON(http.StatusBadRequest, resPackage.ResponseGenderMessageAndError{
			Error: errors.New("bad request"),
		})
		return
	}

	result, err := h.serviceGender.Create(req.Name)
	if err != nil {
		log.Printf("error: handler_gendere_create_service -> %v\n", err)
		c.JSON(http.StatusInternalServerError, resPackage.ResponseGenderMessageAndError{
			Error: errors.New("internal server error"),
		})
		return
	}

	c.JSON(http.StatusOK, resPackage.ResponseGenderMessageAndError{
		Message: result,
		Error:   err,
	})

}

func (h *handlerGender) GetRow(c *gin.Context) {
	var req reqPackage.RequestGender

	if err := c.BindJSON(&req); err != nil {
		log.Printf("error: handler_gendere_GetRow_binding_request -> %v\n", err)
		c.JSON(http.StatusBadRequest, resPackage.ResponseGenderMessageAndError{
			Error: errors.New("bad request"),
		})
		return
	}

	var res repository.Gender
	if req.ID > 0 {
		res = h.serviceGender.GetRow(req.ID)
	} else if req.Name != "" {
		res = h.serviceGender.GetRowByName(req.Name)
	} else {
		genders := h.serviceGender.GetRows()
		c.JSON(http.StatusOK, genders)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handlerGender) GetRows(c *gin.Context) {
	res := h.serviceGender.GetRows()
	c.JSON(http.StatusOK, res)
}

func (h *handlerGender) Update(c *gin.Context) {
	var req reqPackage.RequestGender

	if err := c.BindJSON(&req); err != nil {
		log.Printf("error: handler_gendere_Update_binding_request -> %v\n", err)
		c.JSON(http.StatusBadRequest, resPackage.ResponseGenderMessageAndError{
			Error: errors.New("bad request"),
		})
		return
	}

	res, _ := h.serviceGender.Update(req.ID, req.Name)

	c.JSON(http.StatusOK, res)
}

func (h *handlerGender) Delete(c *gin.Context) {
	var req reqPackage.RequestGender

	if err := c.BindJSON(&req); err != nil {
		log.Printf("error: handler_gendere_Update_binding_request -> %v\n", err)
		c.JSON(http.StatusBadRequest, resPackage.ResponseGenderMessageAndError{
			Error: errors.New("bad request"),
		})
		return
	}

	err := h.serviceGender.Delete(req.ID)
	c.JSON(http.StatusOK, err)
}
