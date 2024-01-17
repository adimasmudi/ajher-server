package answer

import (
	"ajher-server/internal/question"
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

type Repository interface {
	Save(answer Answer, collectionName string) (Answer, error)
	GetUserAnswers(quizID string, userID string) ([]AnswerWithQuestion, error)
	Update(answers []Answer, collectionName string) ([]Answer, error)
}

type repository struct {
	db *firestore.Client
}

func NewRepository(db *firestore.Client) *repository {
	return &repository{db}
}

func (r *repository) Save(answer Answer, collectionName string) (Answer, error) {
	// Use the Firestore client to add a new document to the collection
	doc, _, err := r.db.Collection(collectionName).Add(context.Background(), answer)
	if err != nil {
		log.Printf("Failed to create answer: %v", err)
		return answer, err
	}

	answer.ID = doc.ID

	return answer, nil
}

func (r *repository) GetUserAnswers(quizID string, userID string) ([]AnswerWithQuestion, error) {
	var answersWithQuestions []AnswerWithQuestion

	// Example Firestore collection names
	participationsCollection := "participations"
	quizzesCollection := "quizzes"
	questionsCollection := "questions"
	answersCollection := "answers"

	// Retrieve the user's participation document
	participationDoc, err := r.db.Collection(participationsCollection).
		Where("userId", "==", userID).
		Where("quizId", "==", quizID).
		Documents(context.Background()).
		Next()

	if err != nil {
		log.Printf("Error retrieving participation document: %v", err)
		return answersWithQuestions, err
	}

	participationData := participationDoc.Data()

	// Retrieve the quiz document
	quizDoc, err := r.db.Collection(quizzesCollection).
		Doc(participationData["quizId"].(string)).
		Get(context.Background())

	if err != nil {
		log.Printf("Error retrieving quiz document: %v", err)
		return answersWithQuestions, err
	}

	_ = quizDoc.Data()

	// Retrieve all questions for the quiz
	questionsQuery := r.db.Collection(questionsCollection).
		Where("quizId", "==", quizID).
		Documents(context.Background())

	questionDocs, err := questionsQuery.GetAll()
	if err != nil {
		log.Printf("Error retrieving question documents: %v", err)
		return answersWithQuestions, err
	}

	// Iterate over each question and retrieve answers for the user
	for _, questionDoc := range questionDocs {
		questionData := questionDoc.Data()

		// Retrieve the answer document for the user
		answerDoc, err := r.db.Collection(answersCollection).
			Where("userId", "==", userID).
			Where("questionId", "==", questionDoc.Ref.ID).
			Documents(context.Background()).
			Next()

		if err != nil {
			log.Printf("Error retrieving answer document: %v", err)
			continue
		}

		answerData := answerDoc.Data()

		// Create an AnswerWithQuestion struct based on the retrieved data
		answerWithQuestion := AnswerWithQuestion{
			Answer: Answer{
				ID:                  answerDoc.Ref.ID,
				UserID:              userID,
				QuestionID:          questionDoc.Ref.ID,
				Grade:               answerData["grade"].(float64),
				Label:               answerData["label"].(string),
				AnswerText:          answerData["answerText"].(string),
				AnswerDuration:      answerData["answerDuration"].(int64),
				Status:              answerData["status"].(string),
				GeneratedSuggestion: answerData["generatedSuggestion"].(string),
				CreatedAt:           answerData["createdAt"].(time.Time),
				UpdatedAt:           answerData["updatedAt"].(time.Time),
			},
			Question: question.Question{
				ID:              questionDoc.Ref.ID,
				QuizId:          questionData["quizId"].(string),
				Question:        questionData["question"].(string),
				ReferenceAnswer: questionData["referenceAnswer"].(string),
				GradePercentage: questionData["gradePercentage"].(float64),
				Status:          questionData["status"].(string),
				Duration:        questionData["duration"].(int64),
				Point:           questionData["point"].(float64),
				CreatedAt:       questionData["createdAt"].(time.Time),
				UpdatedAt:       questionData["updatedAt"].(time.Time),
			},
		}

		// Add the answerWithQuestion to the answersWithQuestions slice
		answersWithQuestions = append(answersWithQuestions, answerWithQuestion)
	}

	return answersWithQuestions, nil
}

func (r *repository) Update(answers []Answer, collectionName string) ([]Answer, error) {
	ctx := context.Background()
	// Assuming "answers" is a collection reference in Firestore
	answersCollection := r.db.Collection(collectionName)

	for _, answer := range answers {
		// Set the updatedAt field to the current time
		answer.UpdatedAt = time.Now()

		// Create a new document reference with an automatically generated ID
		docRef := answersCollection.Doc(answer.ID)

		// Update the document in Firestore
		_, err := docRef.Set(ctx, answer)
		if err != nil {
			log.Printf("Error updating answer document: %v", err)
			return nil, err
		}

		// Set the ID field of the Answer struct with the generated ID
		answer.ID = docRef.ID
	}

	// Return the updated answers with their IDs
	return answers, nil
}
