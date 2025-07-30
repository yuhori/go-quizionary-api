package quiz

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"math/rand"

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
	quizChapter int,
	quizNum int,
) ([]Quiz, error) {
	switch quizType {
	case FourOptionQuiz:
		return pickRandomQuizzes(qm.fourOptionQuizzes[quizChapter], quizNum), nil
	default:
		return nil, fmt.Errorf("unsupported quiz type: %s", quizType)
	}
}

// func (qm *QuizManager) FilterQuizzes(
// 	quizType QuizType,
// 	targetTags []string,
// ) ([]Quiz, error) {
// 	switch quizType {
// 	case FourOptionQuiz:
// 		return qm.filterFourOptionQuizzes(targetTags), nil
// 	default:
// 		return nil, fmt.Errorf("unsupported quiz type: %s", quizType)
// 	}
// }

func pickRandomQuizzes(quizzes Quizzes, count int) Quizzes {
	if len(quizzes) < count {
		count = len(quizzes)
	}

	// シャッフル
	shuffled := make(Quizzes, len(quizzes))
	copy(shuffled, quizzes)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	// 先頭から count 個取得
	return shuffled[:count]
}
