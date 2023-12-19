package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/obadoraibu/wordle_plus_backend.git/internal/domain"
)

type Service interface {
	NewWord(c *gin.Context, r *domain.NewWordRequest) (*domain.NewWordResponse, error)
	CheckWord(c *gin.Context, r *domain.CheckWordRequest) (string, error)
	DailyWord(c *gin.Context) (*domain.DailyWordResponse, error)
}

type Handler struct {
	service Service
}

type Dependencies struct {
	Service Service
}

func NewHandler(deps Dependencies) *Handler {
	return &Handler{
		service: deps.Service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/new_word", h.NewWord)
	router.GET("/check_word/:word", h.CheckWord)
	router.GET("/daily_word", h.DailyWord)

	return router
}
