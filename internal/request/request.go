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
	QuizChapter int
	QuizNum     int
}

type GetChaptersRequest struct {
	BookType BookType
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

func ParseGetChaptersRequest(c *gin.Context) (*GetChaptersRequest, error) {
	// クエリパラメータから値を取得
	bookType := c.DefaultQuery("type", string(AllBookType))

	// バリデーション
	if bookType != string(OldBookType) && bookType != string(NewBookType) && bookType != string(AllBookType) {
		return nil, fmt.Errorf("invalid book type: %s", bookType)
	}

	return &GetChaptersRequest{
		BookType: BookType(bookType),
	}, nil
}
