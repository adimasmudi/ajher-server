package quiz

import (
	"ajher-server/internal/question"
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Repository interface {
	Save(quiz Quiz, collectionName string) (Quiz, error)
	GetById(id string, collectionName string) (Quiz, error)
	GetByCode(code string, collectionName string) (Quiz, error)
}

type repository struct {
	db *firestore.Client
}

func NewRepository(db *firestore.Client) *repository {
	return &repository{db}
}

func (r *repository) Save(quiz Quiz, collectionName string) (Quiz, error) {
	// Use the Firestore client to add a new document to the collection
	doc, _, err := r.db.Collection(collectionName).Add(context.Background(), quiz)
	if err != nil {
		log.Printf("Failed to create quiz: %v", err)
		return quiz, err
	}

	quiz.ID = doc.ID

	return quiz, nil
}

func (r *repository) GetById(id string, collectionName string) (Quiz, error) {
	var quiz Quiz

	// Create a reference to the quiz document using the ID
	docRef := r.db.Collection(collectionName).Doc(id)

	// Get the quiz document snapshot
	docSnapshot, err := docRef.Get(context.Background())
	if err != nil {
		log.Printf("Failed to get quiz by ID: %v", err)
		return quiz, err
	}

	// Unmarshal the quiz document snapshot into the Quiz struct
	if err := docSnapshot.DataTo(&quiz); err != nil {
		log.Printf("Failed to unmarshal quiz document: %v", err)
		return quiz, err
	}

	quiz.ID = docRef.ID

	log.Printf("Quiz ID before fetching questions: %s", quiz.ID)

	// Fetch questions from the subcollection
	questionsRef := r.db.Collection("questions").Where("quizId", "==", quiz.ID)
	questionsSnapshot, err := questionsRef.Documents(context.Background()).GetAll()
	if err != nil {
		log.Printf("Failed to get questions for quiz: %v", err)
		return quiz, err
	}

	log.Printf("Fetched %d questions for quiz: %s", len(questionsSnapshot), quiz.ID)

	// Loop through the questions and add them to the quiz
	for _, questionDoc := range questionsSnapshot {
		var question question.Question
		if err := questionDoc.DataTo(&question); err != nil {
			log.Printf("Failed to unmarshal question document: %v", err)
			return quiz, err
		}

		question.ID = questionDoc.Ref.ID
		// Assuming Quiz struct has a Questions field to store the questions
		quiz.Question = append(quiz.Question, question)
	}

	// Fetch category data based on quizCategoryId
	categoryRef := r.db.Collection("quizCategories").Doc(quiz.QuizCategoryId)
	categorySnapshot, err := categoryRef.Get(context.Background())
	if err != nil {
		log.Printf("Failed to get category for quiz: %v", err)
		return quiz, err
	}

	// Unmarshal category data into the Quiz struct
	if err := categorySnapshot.DataTo(&quiz.QuizCategory); err != nil {
		log.Printf("Failed to unmarshal category document: %v", err)
		return quiz, err
	}

	return quiz, nil
}

func (r *repository) GetByCode(code string, collectionName string) (Quiz, error) {
	var quiz Quiz

	// Assuming you have a collection named "quizzes"
	iter := r.db.Collection(collectionName).Where("code", "==", code).Limit(1).Documents(context.Background())

	// Iterate over the result set
	for {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return quiz, err
		}

		// Populate quiz struct with retrieved data and ID
		if err := docSnap.DataTo(&quiz); err != nil {
			return quiz, err
		}

		// Set the ID field in the quiz struct
		quiz.ID = docSnap.Ref.ID

		// Fetch questions from the subcollection
		questionsRef := docSnap.Ref.Collection("questions").Where("quizId", "==", quiz.ID)
		questionsSnapshot, err := questionsRef.Documents(context.Background()).GetAll()
		if err != nil {
			log.Printf("Failed to get questions for quiz: %v", err)
			return quiz, err
		}

		// Loop through the questions and add them to the quiz
		for _, questionDoc := range questionsSnapshot {
			var question question.Question
			if err := questionDoc.DataTo(&question); err != nil {
				log.Printf("Failed to unmarshal question document: %v", err)
				return quiz, err
			}
			question.ID = questionDoc.Ref.ID
			// Assuming Quiz struct has a Questions field to store the questions
			quiz.Question = append(quiz.Question, question)
		}
	}

	return quiz, nil
}
