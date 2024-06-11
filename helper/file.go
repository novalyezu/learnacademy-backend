package helper

import (
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type UploadParams struct {
	File io.Reader
	Dest string
}

type FileService interface {
	Upload(params UploadParams) (string, error)
	OpenFormFile(c *gin.Context, fh *multipart.FileHeader) multipart.File
	ValidateImage(c *gin.Context, fh *multipart.FileHeader, field string, maxSize int64) error
}

type fileService struct {
	cloudinary *cloudinary.Cloudinary
	bucket     string
	imageUrl   string
}

func NewFileService() FileService {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}
	bucket := os.Getenv("CLOUDINARY_BUCKET")
	cloudName := os.Getenv("CLOUDINARY_NAME")
	return &fileService{
		cloudinary: cld,
		bucket:     bucket,
		imageUrl:   "https://res.cloudinary.com/" + cloudName + "/image/upload/",
	}
}

func (s *fileService) Upload(params UploadParams) (string, error) {
	ctx := context.Background()
	uniqueFilename := false
	overwrite := true
	result, errUpload := s.cloudinary.Upload.Upload(ctx, params.File, uploader.UploadParams{
		Folder:         s.bucket + params.Dest,
		UniqueFilename: &uniqueFilename,
		Overwrite:      &overwrite,
	})

	if errUpload != nil {
		return "", errUpload
	}

	url := s.imageUrl + result.PublicID + "." + result.Format

	return url, nil
}

func (s *fileService) OpenFormFile(c *gin.Context, fh *multipart.FileHeader) multipart.File {
	file, errOpen := fh.Open()
	if errOpen != nil {
		c.JSON(http.StatusInternalServerError, WrapperResponse(http.StatusInternalServerError, "InternalServerError", errOpen.Error(), nil))
		return nil
	}
	defer file.Close()
	return file
}

func (s *fileService) ValidateImage(c *gin.Context, fh *multipart.FileHeader, field string, maxSize int64) error {
	allowedType := []string{"image/png", "image/jpg", "image/jpeg"}
	fileType := fh.Header.Get("Content-Type")
	fileSize := fh.Size

	errorTypeMsg := field + " type allowed: " + strings.Join(allowedType, ",")
	if !slices.Contains(allowedType, fileType) {
		return errors.New(errorTypeMsg)
	}

	errorMaxSizeMsg := field + " max size is " + strconv.Itoa(int(maxSize)) + " bytes"
	if fileSize > maxSize {
		return errors.New(errorMaxSizeMsg)
	}

	return nil
}
