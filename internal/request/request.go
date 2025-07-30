package request

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookType string

const (
	OldBookType BookType = "old"
	NewBookType BookType = "new"
	AllBookType BookType = "all"
)

type GetFourOptionQuizzesRequest struct {
	QuizIndex int
	QuizNum   int
}

type GetTitlesRequest struct {
	BookType BookType
}

func ParseGetFourOptionQuizzesRequest(c *gin.Context) (*GetFourOptionQuizzesRequest, error) {
	// クエリパラメータから値を取得
	index := c.Query("index")
	num := c.Query("num")

	// int に変換
	quizIndex, err := strconv.Atoi(index)
	if err != nil {
		return nil, fmt.Errorf("invalid index: %w", err)
	}
	quizNum, err := strconv.Atoi(num)
	if err != nil {
		return nil, fmt.Errorf("invalid num: %w", err)
	}

	// バリデーション
	if quizIndex < 1 || quizIndex > 66 {
		return nil, fmt.Errorf("index must be between 1 and 66")
	}
	if quizNum < 1 {
		return nil, fmt.Errorf("num must be greater than 0")
	}
	return &GetFourOptionQuizzesRequest{
		QuizIndex: quizIndex,
		QuizNum:   quizNum,
	}, nil
}

func ParseGetTitlesRequest(c *gin.Context) (*GetTitlesRequest, error) {
	// クエリパラメータから値を取得
	bookType := c.DefaultQuery("type", string(AllBookType))

	// バリデーション
	if bookType != string(OldBookType) && bookType != string(NewBookType) && bookType != string(AllBookType) {
		return nil, fmt.Errorf("invalid book type: %s", bookType)
	}

	return &GetTitlesRequest{
		BookType: BookType(bookType),
	}, nil
}
