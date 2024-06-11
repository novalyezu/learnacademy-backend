package community

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/novalyezu/learnacademy-backend/helper"
	"github.com/novalyezu/learnacademy-backend/user"
)

type CommunityHandler struct {
	communityService CommunityService
	fileService      helper.FileService
}

func NewCommunityHandler(communityService CommunityService, fileService helper.FileService) *CommunityHandler {
	return &CommunityHandler{
		communityService: communityService,
		fileService:      fileService,
	}
}

func (h *CommunityHandler) CreateCommunity(c *gin.Context) {
	var input CreateCommunityInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.ValidationErrorResponse(err)
		c.JSON(http.StatusBadRequest, helper.WrapperResponse(http.StatusBadRequest, "BadRequest", "input validation errors", gin.H{"errors": errors}))
		return
	}

	currUser := c.MustGet("currentUser").(user.User)
	input.UserID = currUser.ID
	thumbnailFH, errBindFile := c.FormFile("thumbnail")
	if errBindFile != nil {
		c.JSON(http.StatusBadRequest, helper.WrapperResponse(http.StatusBadRequest, "BadRequest", "thumbnail required", nil))
		return
	}
	errValidateFile := h.fileService.ValidateImage(c, thumbnailFH, "thumbnail", 1_000_000)
	if errValidateFile != nil {
		c.JSON(http.StatusBadRequest, helper.WrapperResponse(http.StatusBadRequest, "BadRequest", errValidateFile.Error(), nil))
		return
	}
	thumbnailFile := h.fileService.OpenFormFile(c, thumbnailFH)

	thumbnailUrl, errUpload := h.fileService.Upload(helper.UploadParams{
		File: thumbnailFile,
		Dest: "/community",
	})
	if errUpload != nil {
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, "InternalServerError", errUpload.Error(), nil))
		return
	}

	input.Thumbnail = thumbnailUrl
	newCommunity, errCreate := h.communityService.Create(input)
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, helper.WrapperResponse(http.StatusInternalServerError, "InternalServerError", errCreate.Error(), nil))
		return
	}

	newCommunity.User = currUser
	output := FormatToCommunityOutput(newCommunity)

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, "OK", "Create community success", output))
}
