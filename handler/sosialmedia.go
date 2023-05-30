package handler

import (
	"mygram/helper"
	"mygram/sosialMedia"
	"mygram/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type sosmedHandler struct {
	sosmedService sosialMedia.Service
}

func NewSosmedHandler(service sosialMedia.Service) *sosmedHandler {
	return &sosmedHandler{service}
}

func (h *sosmedHandler) DeletedSosmed(c *gin.Context) {
	var inputID sosialMedia.GetSosmedInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedSosmed, err := h.sosmedService.DeletedSosmed(inputID.ID)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, deletedSosmed)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, "Your social media has been successfuly deleted")
	c.JSON(http.StatusOK, response)

}

func (h *sosmedHandler) UpdateSosmed(c *gin.Context) {
	var inputID sosialMedia.GetSosmedInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData sosialMedia.SosmedInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User.ID = currentUser.ID

	newSosmed, err := h.sosmedService.UpdateSosmed(inputID, inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, sosialMedia.FormatterUpdate(newSosmed))
	c.JSON(http.StatusOK, response)

}

func (h *sosmedHandler) GetSosmed(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	sosmed, err := h.sosmedService.GetSosmed(userID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get sosmed")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, sosialMedia.FormatterGetSosmed(sosmed))
	c.JSON(http.StatusOK, response)
}

func (h *sosmedHandler) CreateSosmed(c *gin.Context) {
	var input sosialMedia.SosmedInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to create Sosial Media")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User.ID = currentUser.ID

	newSosmed, err := h.sosmedService.CreateSosmed(input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to create Sosial Media")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, sosialMedia.FormatterSosmed(newSosmed))
	c.JSON(http.StatusOK, response)
}
