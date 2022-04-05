package api

import (
	v1 "Clever_City/internal/api/v1"
	"Clever_City/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Router struct {
	uc     *usecase.Handler
	router *gin.Engine
}

func NewRouter(uc *usecase.Handler) *Router {
	return &Router{uc: uc}
}

func (r *Router) Start() {
	r.router = gin.Default()
	_ = r.router.SetTrustedProxies([]string{"192.168.0.102"})

	v1.RegisterRoutes(r.router, r.uc)

	r.router.Run()
}
