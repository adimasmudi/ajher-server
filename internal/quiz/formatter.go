package quiz

import (
	"ajher-server/internal/participation"
	"ajher-server/internal/question"
	"fmt"
	"math/rand"
	"time"
)

type QuizFormatter struct {
	ID             string  `json:"id"`
	Title          string  `json:"title"`
	Code           string  `json:"code"`
	Description    string  `json:"description"`
	CategoryName   string  `json:"category_name"`
	Creator        string  `json:"creator"`
	QuestionAmount int     `json:"question_amount"`
	TotalTime      string  `json:"total_time"`
	TotalPoint     float64 `json:"total_point"`
}

func FormatQuiz(quiz Quiz, participation participation.Participation) QuizFormatter {
	formatter := QuizFormatter{
		ID:             quiz.ID,
		Title:          quiz.Title,
		Code:           quiz.Code,
		Description:    quiz.Description,
		CategoryName:   quiz.QuizCategory.CategoryName,
		Creator:        participation.User.FullName,
		QuestionAmount: len(quiz.Question),
		TotalPoint:     getTotalPoint(quiz.Question),
		TotalTime:      getTotalTime(quiz.Question),
	}

	return formatter
}

func getTotalPoint(questions []question.Question) float64 {
	points := 0.0
	for _, question := range questions {
		points += question.Point
	}

	return points
}

func getTotalTime(questions []question.Question) string {

	seconds := 0

	for _, question := range questions {
		seconds += question.Duration
	}

	duration := time.Second * time.Duration(seconds)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	remainingSeconds := seconds % 60

	var formattedTime string

	if hours > 0 {
		formattedTime += fmt.Sprintf("%d hour", hours)
		if hours > 1 {
			formattedTime += "s"
		}
	}

	if minutes > 0 {
		if formattedTime != "" {
			formattedTime += " "
		}
		formattedTime += fmt.Sprintf("%d minute", minutes)
		if minutes > 1 {
			formattedTime += "s"
		}
	}

	if remainingSeconds > 0 {
		if formattedTime != "" {
			formattedTime += " "
		}
		formattedTime += fmt.Sprintf("%d second", remainingSeconds)
		if remainingSeconds > 1 {
			formattedTime += "s"
		}
	}

	return formattedTime
}

func ShuffleArray(arr []question.Question) []question.Question {

	rand.Seed(time.Now().UnixNano())
	n := len(arr)

	fmt.Println("arr above", arr)

	// Fisher-Yates shuffle algorithm
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Println("arr", arr)

	return arr

}
