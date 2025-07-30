package quiz

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type Quiz struct {
	Question    string   `json:"question"`
	Choices     []string `json:"choices"`
	Answer      string   `json:"answer"`
	Explanation string   `json:"explanation"`
	Sources     []string `json:"sources"`
	Tags        []string `json:"tags"`
}

type QuizType string

const (
	FourOptionQuiz QuizType = "four_option"
)

type QuizManager struct {
	fourOptionQuizzes map[string]Quiz
}

func New(dir string) (*QuizManager, error) {
	manager := &QuizManager{
		fourOptionQuizzes: make(map[string]Quiz),
	}

	// dir 以下のファイルを全て読み込む
	if err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to access path %s: %w", path, err)
		}

		if !d.IsDir() {
			// ファイル内容を読み取る
			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", path, err)
			}

			var quiz Quiz
			if err := json.Unmarshal(content, &quiz); err != nil {
				return fmt.Errorf("failed to parse JSON in file %s: %w", path, err)
			}

			manager.fourOptionQuizzes[path] = quiz
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to walk directory %s: %w", dir, err)
	}

	// QuizManagerの初期化
	return manager, nil
}

func (qm *QuizManager) ChooseQuizzes(
	quizType QuizType,
	quizNum int,
	targetTags []string,
) ([]string, error) {
	// ここに4択問題の取得ロジックを実装
	return []string{"Question 1", "Question 2", "Question 3", "Question 4"}, nil
}
