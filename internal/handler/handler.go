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

func (h *Handler) GetFourOptionQuizzes(c *gin.Context) {
	// QuizManagerから4択問題を取得
	quizzes, err := h.quizManager.ChooseQuizzes(
		quiz.FourOptionQuiz,
		1,
		10,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("failed to get quizzes: %v", err)})
		return
	}
	// 取得した問題をJSON形式で返す
	c.JSON(200, gin.H{"quizzes": quizzes})
}

func (h *Handler) OK(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
