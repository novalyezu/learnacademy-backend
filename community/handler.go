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
		helper.ErrorHandler(c, err)
		return
	}

	currUser := c.MustGet("currentUser").(user.UserOutput)
	input.UserID = currUser.ID
	thumbnailFH, errBindFile := c.FormFile("thumbnail")
	if errBindFile != nil {
		helper.ErrorHandler(c, helper.NewBadRequestError("thumbnail required"))
		return
	}
	errValidateFile := h.fileService.ValidateImage(c, thumbnailFH, "thumbnail", 1_000_000)
	if errValidateFile != nil {
		helper.ErrorHandler(c, errValidateFile)
		return
	}
	thumbnailFile, errOpenFile := h.fileService.OpenFormFile(c, thumbnailFH)
	if errOpenFile != nil {
		helper.ErrorHandler(c, errOpenFile)
		return
	}

	thumbnailUrl, errUpload := h.fileService.Upload(helper.UploadParams{
		File: thumbnailFile,
		Dest: "/community",
	})
	if errUpload != nil {
		helper.ErrorHandler(c, errUpload)
		return
	}

	input.Thumbnail = thumbnailUrl
	newCommunity, errCreate := h.communityService.Create(input)
	if errCreate != nil {
		helper.ErrorHandler(c, errCreate)
		return
	}

	newCommunity.User = currUser

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, "OK", "Create community success", newCommunity))
}

func (h *CommunityHandler) GetCommunities(c *gin.Context) {
	var input GetCommunitiesInput

	err := c.ShouldBindQuery(&input)
	if err != nil {
		helper.ErrorHandler(c, err)
		return
	}

	communities, errGet := h.communityService.GetAll(input)
	if errGet != nil {
		helper.ErrorHandler(c, errGet)
		return
	}

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, "OK", "Get communities success", communities))
}

func (h *CommunityHandler) GetCommunityByID(c *gin.Context) {
	var input GetCommunityByIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorHandler(c, err)
		return
	}

	community, errGet := h.communityService.GetByID(input.ID)
	if errGet != nil {
		helper.ErrorHandler(c, errGet)
		return
	}

	c.JSON(http.StatusOK, helper.WrapperResponse(http.StatusOK, "OK", "Get community success", community))
}
