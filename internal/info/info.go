package info

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yuhori/go-quizionary-api/internal/request"
)

type Chapter struct {
	Chapter int `json:"chapter"`
	Verses  int `json:"verses"`
}

type Book struct {
	Name     string    `json:"name"`
	Chapters []Chapter `json:"chapters"`
}

type Info struct {
	Old []Book `json:"old"`
	New []Book `json:"new"`
}

type InfoManager struct {
	info Info
}

func New(path string) (*InfoManager, error) {
	// ファイルを読み込む
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// JSONをパース
	var info Info
	if err := json.Unmarshal(content, &info); err != nil {
		return nil, fmt.Errorf("failed to parse JSON in file %s: %w", path, err)
	}

	// InfoManagerの初期化
	im := &InfoManager{
		info: info,
	}
	return im, nil
}

// GetTitles
func (im *InfoManager) GetTitles(bookType request.TestamentType) []string {
	titles := make([]string, 0)

	switch bookType {
	case request.OldTestamentType:
		for _, book := range im.info.Old {
			titles = append(titles, book.Name)
		}
		return titles
	case request.NewTestamentType:
		for _, book := range im.info.New {
			titles = append(titles, book.Name)
		}
		return titles
	case request.AllTestamentType:
		// 古約と新約の両方のタイトルを追加
		for _, book := range im.info.Old {
			titles = append(titles, book.Name)
		}
		for _, book := range im.info.New {
			titles = append(titles, book.Name)
		}
		return titles
	default:
		// デフォルトは全書のタイトルを返す
		// 古約と新約の両方のタイトルを追加
		for _, book := range im.info.Old {
			titles = append(titles, book.Name)
		}
		for _, book := range im.info.New {
			titles = append(titles, book.Name)
		}
		return titles
	}
}
