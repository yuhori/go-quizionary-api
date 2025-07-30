package request

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TestamentType string

const (
	OldTestamentType TestamentType = "old"
	NewTestamentType TestamentType = "new"
	AllTestamentType TestamentType = "all"
)

type GetFourOptionQuizzesRequest struct {
	TestamentType TestamentType
	Index         int
	Num           int
}

type GetTitlesRequest struct {
	TestamentType TestamentType
}

func ParseGetFourOptionQuizzesRequest(c *gin.Context) (*GetFourOptionQuizzesRequest, error) {
	// クエリパラメータから値を取得
	testamentType := c.DefaultQuery("type", string(AllTestamentType))
	indexStr := c.Query("index")
	numStr := c.Query("num")

	// int に変換
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return nil, fmt.Errorf("invalid index: %w", err)
	}
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return nil, fmt.Errorf("invalid num: %w", err)
	}

	// バリデーション
	if testamentType != string(OldTestamentType) && testamentType != string(NewTestamentType) && testamentType != string(AllTestamentType) {
		return nil, fmt.Errorf("invalid type: %s, must be 'old', 'new', or 'all'", testamentType)
	}

	// index の調整
	if testamentType == string(NewTestamentType) {
		index += 39
	}
	if index < 1 || index > 66 {
		return nil, fmt.Errorf("index must be between 1 and 66")
	}
	if num < 1 {
		return nil, fmt.Errorf("num must be greater than 0")
	}
	return &GetFourOptionQuizzesRequest{
		Index: index,
		Num:   num,
	}, nil
}

func ParseGetTitlesRequest(c *gin.Context) (*GetTitlesRequest, error) {
	// クエリパラメータから値を取得
	testamentType := c.DefaultQuery("type", string(AllTestamentType))

	// バリデーション
	if testamentType != string(OldTestamentType) && testamentType != string(NewTestamentType) && testamentType != string(AllTestamentType) {
		return nil, fmt.Errorf("invalid type: %s, must be 'old', 'new', or 'all'", testamentType)
	}

	return &GetTitlesRequest{
		TestamentType: TestamentType(testamentType),
	}, nil
}
