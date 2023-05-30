package handler

import (
	"mygram/comment"
	"mygram/helper"
	"mygram/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService comment.Service
}

func NewCommentHandler(service comment.Service) *commentHandler {
	return &commentHandler{service}
}

func (h *commentHandler) GetComments(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	photoID, _ := strconv.Atoi(c.Query("photo_id"))

	newComment, err := h.commentService.GetComment(userID, photoID)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to get Comments")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := comment.FormatterGetComment(newComment)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *commentHandler) CreateComment(c *gin.Context) {
	var input comment.CommentInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to create Comments")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	// userId := currentUser.ID
	input.User = currentUser

	newComment, err := h.commentService.CreateComment(input)

	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Eror to create Comments")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, comment.FormatterComment(newComment))
	c.JSON(http.StatusOK, response)

}

func (h *commentHandler) UpdateComment(c *gin.Context) {
	var inputID comment.GetCommentInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData comment.UpdateCommentInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedComment, err := h.commentService.UpdateComment(inputID, inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, comment.FormatterUpdated(updatedComment))
	c.JSON(http.StatusOK, response)
}

func (h *commentHandler) DeletedComment(c *gin.Context) {
	var inputID comment.GetCommentInput

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deletedComment, err := h.commentService.DeleteComment(inputID.ID)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, deletedComment)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, "Your comment has been succesfully deleted")
	c.JSON(http.StatusOK, response)
}
