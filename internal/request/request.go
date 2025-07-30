package request

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetFourOptionQuizzesRequest struct {
	QuizChapter int
	QuizNum     int
}

func ParseGetFourOptionQuizzesRequest(c *gin.Context) (*GetFourOptionQuizzesRequest, error) {
	// クエリパラメータから値を取得
	chapter := c.Query("chapter")
	num := c.Query("num")

	// int に変換
	quizChapter, err := strconv.Atoi(chapter)
	if err != nil {
		return nil, fmt.Errorf("invalid chapter: %w", err)
	}
	quizNum, err := strconv.Atoi(num)
	if err != nil {
		return nil, fmt.Errorf("invalid num: %w", err)
	}

	// バリデーション
	if quizChapter < 1 || quizChapter > 66 {
		return nil, fmt.Errorf("chapter must be between 1 and 66")
	}
	if quizNum < 1 {
		return nil, fmt.Errorf("num must be greater than 0")
	}
	return &GetFourOptionQuizzesRequest{
		QuizChapter: quizChapter,
		QuizNum:     quizNum,
	}, nil
}
