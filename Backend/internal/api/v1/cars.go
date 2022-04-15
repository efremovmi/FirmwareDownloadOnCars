package v1

import (
	"Clever_City/internal/usecase"
	"github.com/gin-gonic/gin"
)

//
//func UploadFile(uc *usecase.Handler) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		file, err := c.FormFile("file")
//		if file == nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"message": "Отсутствует файл",
//			})
//			return
//		}
//
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{
//				"message": err.Error(),
//			})
//			return
//		}
//
//		err = uc.TransferFile(file, c.Request.FormValue("name"))
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{
//				"message": err.Error(),
//			})
//			return
//		}
//
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.JSON(http.StatusOK, gin.H{
//			"message": fmt.Sprintf("'%s' загружен!", file.Filename),
//		})
//	}
//}
//
//func Kill(uc *usecase.Handler) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		err := uc.KillProcess(c.Request.FormValue("name"))
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{
//				"message": err.Error(),
//			})
//			return
//		}
//
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.JSON(http.StatusOK, gin.H{
//			"message": "Скрипт остановлен!",
//		})
//	}
//}

func UploadLogs(uc *usecase.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
