package v1

import (
	"Clever_City/internal/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(uc *usecase.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if file == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Отсутствует файл",
			})
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		err = uc.TransferFile(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' загружен!", file.Filename),
		})
	}
}
