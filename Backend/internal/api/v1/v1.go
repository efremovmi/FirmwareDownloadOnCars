package v1

import (
	"Clever_City/internal/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, uc *usecase.Handler) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/new", UploadFile(uc))
	}

}
