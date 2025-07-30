package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuhori/go-quizionary-api/internal/quiz"
	"github.com/yuhori/go-quizionary-api/internal/request"
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
	// リクエストをパース
	req, err := request.ParseGetFourOptionQuizzesRequest(c)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("invalid request: %v", err)})
		return
	}

	// QuizManagerから4択問題を取得
	quizzes, err := h.quizManager.ChooseQuizzes(
		quiz.FourOptionQuiz,
		req.QuizChapter,
		req.QuizNum,
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
