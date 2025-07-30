package quiz

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/yuhori/go-quizionary-api/internal/utils"
)

type Quiz struct {
	Question    string   `json:"question"`
	Choices     []string `json:"choices"`
	Answer      string   `json:"answer"`
	Explanation string   `json:"explanation"`
	Sources     []string `json:"sources"`
	Tags        []string `json:"tags"`
}

type Quizzes []Quiz

type QuizType string

const (
	FourOptionQuiz QuizType = "four_option"
)

type QuizManager struct {
	fourOptionQuizzes []Quizzes
}

func New(dir string) (*QuizManager, error) {
	// ディレクトリ内のファイルを取得
	var files []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// 数字順にソート
	utils.NumericFileSort(files)

	// ファイルを順に読み込む
	fourOptionQuizzes := make([]Quizzes, 0)
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		var quizzes Quizzes
		if err := json.Unmarshal(content, &quizzes); err != nil {
			return nil, fmt.Errorf("failed to parse JSON in file %s: %w", file, err)
		}
		fourOptionQuizzes = append(fourOptionQuizzes, quizzes)
	}

	// QuizManagerの初期化
	qm := &QuizManager{
		fourOptionQuizzes: fourOptionQuizzes,
	}
	return qm, nil
}

func (qm *QuizManager) ChooseQuizzes(
	quizType QuizType,
	quizNum int,
	targetTags []string,
) ([]Quiz, error) {
	if quizType != FourOptionQuiz {
		return nil, fmt.Errorf("unsupported quiz type: %s", quizType)
	}
	// TODO: 実装
	return qm.fourOptionQuizzes[0], nil
}
