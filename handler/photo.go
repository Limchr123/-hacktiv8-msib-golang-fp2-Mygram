package handler

import (
	"mygram/campaign"
	"mygram/helper"
	"mygram/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	campaignService campaign.Service
}

func NewPhotoHandler(service campaign.Service) *photoHandler {
	return &photoHandler{service}
}

// pengambilan campaign yang bancampaign
func (h *photoHandler) GetCampaigns(c *gin.Context) {
	//karena kita butuh userID dan dia di inisiasi integer maka kita convert menjadi string
	//c.query apa?
	userID, _ := strconv.Atoi(c.Query("user_id"))

	//seperti biasa untuk berikan respon dan gunakan userID yang uda di inisiasi
	photo, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get campaigns")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatterGetCampaign(photo)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)

}

func (h *photoHandler) CreateImage(c *gin.Context) {
	var input campaign.PhotoInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	// userId := currentUser.ID
	input.User.ID = currentUser.ID

	newImg, err := h.campaignService.CreateImage(input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, newImg)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, campaign.FormatterCreateCampaign(newImg))
	c.JSON(http.StatusOK, response)

}

func (h *photoHandler) UpdatedCampaign(c *gin.Context) {
	var inputID campaign.GetPhotoDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData campaign.PhotoInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedCampaign, err := h.campaignService.UpdateCampaigns(inputID, inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, campaign.FormatterUpdatedCampaign(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *photoHandler) DeletePhoto(c *gin.Context) {
	var inputID campaign.GetPhotoDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// currentUser := c.MustGet("currentUser").(user.User)
	// // inputID := currentUser.ID
	// inputID.ID := currentUser.ID

	updatedCampaign, err := h.campaignService.DeletePhoto(inputID.ID)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, updatedCampaign)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, "Your photo has been successfully deleted")
	c.JSON(http.StatusOK, response)

}

func (h *photoHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetPhotoDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// currentUser := c.MustGet("currentUser").(user.User)
	// input.ID = currentUser

	campaignDetail, err := h.campaignService.GetCampaignById(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, campaignDetail)
	c.JSON(http.StatusOK, response)
}
