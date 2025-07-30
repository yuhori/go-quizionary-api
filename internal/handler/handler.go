package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) OK(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
