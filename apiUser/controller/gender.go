package controller

import (
	"apiUser/repository"
	"apiUser/services"
	"apiUser/transport"
	"log"
	"net/http"
	"strings"

	"apiUser/request"
	"apiUser/response"

	"github.com/gin-gonic/gin"
)

type ControllerGender interface {
	Create(c *gin.Context)
	Select(c *gin.Context)
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
	var req request.RequestGender

	// decode request
	err := transport.DecodeRequestJson(c, req)
	if err != nil {
		transport.ReturnBadRequest(c, nil)
		return
	}

	// validate data
	if req.ID == 0 || strings.TrimSpace(req.Name) == "" {
		transport.ReturnBadRequest(c, nil)
		return
	}

	// create data
	res, err := h.serviceGender.Create(req.Name)
	if err != nil {
		log.Printf("error: handler_gendere_create_service -> %v\n", err)
		transport.ReturnInternalServerError(c, nil)
		return
	}

	transport.ReturnOK(c, response.ResponseGenderMessageAndError{
		Message: res,
		Error:   err,
	})
}

func (h *handlerGender) Select(c *gin.Context) {
	var req request.RequestGender

	// decode request
	err := transport.DecodeRequestJson(c, req)
	if err != nil {
		log.Printf("error: controller_gender_select decode_request -> %v\n", err)
		transport.ReturnBadRequest(c, nil)
	}

	// validate request
	var res repository.Gender

	// Response data

	if req.ID > 0 {
		res = h.serviceGender.GetRow(req.ID)
		transport.ReturnOK(c, res)
		return
	}

	if req.Name != "" {
		res = h.serviceGender.GetRowByName(req.Name)
		transport.ReturnOK(c, res)
		return
	}

	if req.ID == 0 && strings.TrimSpace(req.Name) == "" {
		genderList := h.serviceGender.GetRows()
		transport.ReturnOK(c, genderList)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *handlerGender) Update(c *gin.Context) {
	var req request.RequestGender

	// decode request
	err := transport.DecodeRequestJson(c, &req)
	if err != nil {
		/*
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bad request",
			})
		*/
		transport.ReturnBadRequest(c, nil)
		return
	}

	// validate request data
	if req.ID == 0 || req.Name == "" {
		transport.ReturnBadRequest(c, nil)
		return
	}

	// Update data
	res, err := h.serviceGender.Update(req.ID, req.Name)
	if err != nil {
		log.Println("update gender error: ", err)
		transport.ReturnInternalServerError(c, nil)
		return
	}

	// return response
	transport.ReturnOK(c, res)
}

func (h *handlerGender) Delete(c *gin.Context) {
	var req request.RequestGender

	err := transport.DecodeRequestJson(c, req)
	if err != nil {
		log.Printf("error: controller_gender_delete decode request -> %v\n", err)
		transport.ReturnBadRequest(c, nil)
		return
	}

	err = h.serviceGender.Delete(req.ID)
	if err != nil {
		log.Printf("error: controller_gender_delete delete data -> %v\n", err)
		transport.ReturnInternalServerError(c, nil)
		return
	}

	transport.ReturnOK(c, err)
}
