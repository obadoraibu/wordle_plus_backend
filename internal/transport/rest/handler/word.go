package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/obadoraibu/wordle_plus_backend.git/internal/domain"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) NewWord(c *gin.Context) {
	length := c.DefaultQuery("length", "5")

	intLength, err := strconv.Atoi(length)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	req := &domain.NewWordRequest{Length: intLength}

	resp, err := h.service.NewWord(c, req)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, domain.ErrGeneratingNewWord.Error())
		return
	}
	logrus.Print(resp.Word)
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CheckWord(c *gin.Context) {
	word := c.Param("word")

	req := &domain.CheckWordRequest{
		Word: word,
	}

	_, err := h.service.CheckWord(c, req)
	if err != nil {
		if errors.Is(err, domain.ErrWordDoesntExist) {
			sendErrorResponse(c, http.StatusNotFound, domain.ErrWordDoesntExist.Error())
			return
		} else {
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	logrus.Print(word)
	c.Status(200)
}

func (h *Handler) DailyWord(c *gin.Context) {
	resp, err := h.service.DailyWord(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, resp)
}
