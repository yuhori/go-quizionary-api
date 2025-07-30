package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuhori/go-quizionary-api/internal/info"
	"github.com/yuhori/go-quizionary-api/internal/quiz"
	"github.com/yuhori/go-quizionary-api/internal/request"
)

type Handler struct {
	quizManager *quiz.QuizManager
	infoManager *info.InfoManager
}

func New(dir string) (*Handler, error) {
	// QuizManagerの初期化
	quizManager, err := quiz.New(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to create quiz manager: %w", err)
	}

	// InfoManagerの初期化
	infoManager, err := info.New("files/info/testament.json")
	if err != nil {
		return nil, fmt.Errorf("failed to create info manager: %w", err)
	}

	h := &Handler{
		quizManager: quizManager,
		infoManager: infoManager,
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

func (h *Handler) GetChapters(c *gin.Context) {
	// リクエストをパース
	req, err := request.ParseGetChaptersRequest(c)
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("invalid request: %v", err)})
		return
	}

	// QuizManagerから4択問題を取得
	titles := h.infoManager.GetTitles(
		req.BookType,
	)

	// 取得した問題をJSON形式で返す
	c.JSON(200, gin.H{"titles": titles})
}

func (h *Handler) OK(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
