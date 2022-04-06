package usecase

import (
	"Clever_City/internal/repository"
)

type Handler struct {
	repo *repository.Repo
}

func NewHandler(repo *repository.Repo) *Handler {
	return &Handler{repo: repo}
}
