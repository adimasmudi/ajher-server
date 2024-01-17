package participantQuestion

type QuestionFormatter struct {
	ID              string  `json:"id"`
	QuizId          string  `json:"quiz_id"`
	Number          int64   `json:"number"`
	Question        string  `json:"question"`
	GradePercentage float64 `json:"grade_percentage"`
	Point           float64 `json:"point"`
	Status          string  `json:"status"`
	Duration        int64   `json:"duration"`
}

func FormatQuestion(participantQuestion []ParticipantQuestion) []QuestionFormatter {
	var formatResult []QuestionFormatter

	for _, data := range participantQuestion {
		formatter := QuestionFormatter{
			ID:              data.Question.ID,
			QuizId:          data.Question.QuizId,
			Number:          data.Number,
			Question:        data.Question.Question,
			GradePercentage: data.Question.GradePercentage,
			Point:           data.Question.Point,
			Status:          data.Question.Status,
			Duration:        data.Question.Duration,
		}

		formatResult = append(formatResult, formatter)
	}

	return formatResult
}
