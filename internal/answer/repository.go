package answer

import "gorm.io/gorm"

type Repository interface {
	Save(answer Answer) (Answer, error)
	GetUserAnswers(quizId string, userId int) ([]Answer, error)
	Update(answers []Answer) ([]Answer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(answer Answer) (Answer, error) {
	err := r.db.Create(&answer).Error

	if err != nil {
		return answer, err
	}

	return answer, nil
}

func (r *repository) GetUserAnswers(quizId string, userId int) ([]Answer, error) {
	var answers []Answer
	err := r.db.Raw(`
		SELECT answers.*, questions.*
		FROM participations
		JOIN quizzes ON participations.quiz_id = quizzes.id
		JOIN questions ON questions.quiz_id = quizzes.id
		JOIN answers ON answers.question_id = questions.id
		WHERE participations.user_id = ? AND participations.quiz_id = ? AND answers.user_id = ?
	`, userId, quizId, userId).Preload("Question").Find(&answers).Error

	if err != nil {
		return answers, err
	}

	return answers, nil
}

func (r *repository) Update(answers []Answer) ([]Answer, error) {
	err := r.db.Save(&answers).Error

	if err != nil {
		return answers, err
	}

	return answers, nil
}
