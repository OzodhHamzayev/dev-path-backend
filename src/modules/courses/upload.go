package courses

import (
	"crypto/rand"
	"encoding/hex"
	"image"
	"io"
	"net/http"
	"os"
	"path/filepath"

	_ "image/jpeg"
	_ "image/png"
	_ "golang.org/x/image/webp"
	"github.com/gin-gonic/gin"
)



func UploadImage(ctx *gin.Context) { 

	const maxSize int64 = 5 * 1024 * 1024

	ctx.Request.Body = http.MaxBytesReader(
		ctx.Writer,
		ctx.Request.Body,
		maxSize+1024*1024,
	)

	if err := ctx.Request.ParseMultipartForm(1024 * 1024); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not read the uploaded form data",
		})
		return
	}

	if ctx.Request.MultipartForm != nil {
		defer ctx.Request.MultipartForm.RemoveAll()
	}

	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "send an image to image field",
		})
		return
	}
	defer file.Close()

	if header.Size == 0 || header.Size > maxSize {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "The image file is empty or the image must not exceed 5 mb",
		})
		return
	}

	buffer := make([]byte, 512)
	n, err := io.ReadFull(file, buffer)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not read the image",
		})
		return
	}
	extension := ""
	switch http.DetectContentType(buffer[:n]) {
	case "image/jpeg":
		extension = ".jpg"
	case "image/png":
		extension = ".png"
	case "image/webp":
		extension = ".webp"
	default:
		ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error": "Undefined content type",
		})
		return
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not read the image",
		})
		return
	}

	config, _, err := image.DecodeConfig(file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not read the image dimensions",
		})
		return
	}

	const maxDimension = 6000
	const maxPixels int64 = 20_000_000

	if config.Width <= 0 ||
		config.Height <= 0 ||
		config.Width > maxDimension ||
		config.Height > maxDimension ||
		int64(config.Width)*int64(config.Height) > maxPixels {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Image dimensions exceed the allowed limits.",
			})
			return
		}
	
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not rewind the image.",
		})
		return
	}
	if _, _, err := image.Decode(file); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
		"error": "The image is invalid or corrupted.",
		})
		return
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not rewind the image.",
		})
		return
	}

	err = os.MkdirAll("uploads", 0755)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create the uploads directory.",
		})
		return
	}

	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
     	   "error": "Could not generate an image filename.",
   		})
    	return
	}

	fileName := hex.EncodeToString(randomBytes) + extension


	path := filepath.Join("uploads", fileName)

	dst, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_CREATE|os.O_EXCL,
		0644,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
		     "error": "Could not create the image file.",
		})
		return
	}

	_, copyErr := io.Copy(dst, file)
	closeErr := dst.Close()
	if copyErr != nil || closeErr != nil {
		_ = os.Remove(path)
		ctx.JSON(http.StatusInternalServerError, gin.H{
        "error": "Could not save the image.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
    	"image_url": "/uploads/" + fileName,
	})
}

