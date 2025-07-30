package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetFourOptionQuestions(c *gin.Context) {
	// ここに4択問題の取得ロジックを実装
	c.JSON(200, gin.H{"message": "Get four option questions"})
}
