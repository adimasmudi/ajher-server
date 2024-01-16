package answer

type FinishAnswerFormatter struct {
	Point                float64 `json:"point"`
	CorrectAnswer        int     `json:"correct_answer"`
	CompletionPercentage float64 `json:"completion_percentage"`
	Skipped              int     `json:"skipped"`
	IncorrectAnswer      int     `json:"incorrect_answer"`
}

func FormatFinishAnswer(answers []Answer) FinishAnswerFormatter {
	points := 0.0
	correctAnswers := 0
	skipped := 0
	incorrectAnswers := 0
	for _, data := range answers {
		if data.Label == "right" {
			points += data.Question.Point
			correctAnswers += 1
		} else {
			incorrectAnswers += 1
		}

		if data.Answer == "" {
			skipped++
		}
	}

	completionPercentage := float64((1 - (skipped / len(answers))) * 100)

	formatter := FinishAnswerFormatter{
		Point:                points,
		CorrectAnswer:        correctAnswers,
		IncorrectAnswer:      incorrectAnswers,
		CompletionPercentage: completionPercentage,
		Skipped:              skipped,
	}

	return formatter
}
