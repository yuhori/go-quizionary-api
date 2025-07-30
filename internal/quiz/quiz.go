package quiz

type QuizType string

const (
	FourOptionQuiz QuizType = "four_option"
)

type QuizManager struct {
}

func New(path string) *QuizManager {
	return &QuizManager{}
}

func (qm *QuizManager) ChooseQuizzes(
	quizType QuizType,
	quizNum int,
	targetTags []string,
) ([]string, error) {
	// ここに4択問題の取得ロジックを実装
	return []string{"Question 1", "Question 2", "Question 3", "Question 4"}, nil
}
