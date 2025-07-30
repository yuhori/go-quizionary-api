package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuhori/go-quizionary-api/internal/quiz"
)

type Handler struct {
	quizManager *quiz.QuizManager
}

func New(dir string) (*Handler, error) {
	manager, err := quiz.New(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to create quiz manager: %w", err)
	}

	h := &Handler{
		quizManager: manager,
	}
	return h, nil
}

func (h *Handler) GetFourOptionQuestions(c *gin.Context) {
	// ここに4択問題の取得ロジックを実装
	c.JSON(200, gin.H{"message": "Get four option questions"})
}

func (h *Handler) OK(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
