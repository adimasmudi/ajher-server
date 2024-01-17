package answer

type FinishAnswerFormatter struct {
	Point                float64 `json:"point"`
	CorrectAnswer        int64   `json:"correct_answer"`
	CompletionPercentage float64 `json:"completion_percentage"`
	Skipped              int64   `json:"skipped"`
	IncorrectAnswer      int64   `json:"incorrect_answer"`
}

func FormatFinishAnswer(answers []AnswerWithQuestion) FinishAnswerFormatter {
	points := 0.0
	var correctAnswers int64 = 0
	var skipped int64 = 0
	var incorrectAnswers int64 = 0
	for _, data := range answers {
		if data.Label == "right" {
			points += data.Question.Point
			correctAnswers += 1
		} else {
			incorrectAnswers += 1
		}

		if data.AnswerText == "" {
			skipped++
		}
	}

	completionPercentage := float64((1 - (skipped / int64(len(answers)))) * 100)

	formatter := FinishAnswerFormatter{
		Point:                points,
		CorrectAnswer:        correctAnswers,
		IncorrectAnswer:      incorrectAnswers,
		CompletionPercentage: completionPercentage,
		Skipped:              skipped,
	}

	return formatter
}
