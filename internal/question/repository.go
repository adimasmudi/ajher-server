package question

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Repository interface {
	Save(questions []Question, collectionName string) ([]Question, error)
	GetAll(collectionName string) ([]Question, error)
	GetAllByQuizId(quizId string, collectionName string) ([]Question, error)
	GetById(id string, collectionName string) (Question, error)
}

type repository struct {
	db *firestore.Client
}

func NewRepository(db *firestore.Client) *repository {
	return &repository{db}
}

func (r *repository) Save(questions []Question, collectionName string) ([]Question, error) {
	// Initialize an empty slice to store the IDs of the created questions
	var createdQuestions []Question

	// Use a batch to perform multiple writes in a single transaction
	batch := r.db.Batch()

	// Add each question to the batch as a set operation
	for _, q := range questions {
		// Omit setting the ID to let Firestore generate one
		docRef := r.db.Collection(collectionName).NewDoc()

		// Add a set operation to the batch
		batch.Set(docRef, q)

		q.ID = docRef.ID

		// Append the auto-generated ID to the createdQuestionIDs slice
		createdQuestions = append(createdQuestions, q)
	}

	// Commit the batch to Firestore
	_, err := batch.Commit(context.Background())
	if err != nil {
		log.Printf("Failed to save questions: %v", err)
		return nil, err
	}

	// Return the auto-generated question IDs
	return createdQuestions, nil
}

func (r *repository) GetAll(collectionName string) ([]Question, error) {
	var questions []Question

	// Create a query to get all documents from the collection
	iter := r.db.Collection(collectionName).Documents(context.Background())
	defer iter.Stop()

	// Iterate through the documents and unmarshal them into Question structs
	for {
		var question Question
		docSnapshot, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Printf("Failed to fetch question: %v", err)
			return questions, err
		}

		if err := docSnapshot.DataTo(&question); err != nil {
			log.Printf("Failed to unmarshal question document: %v", err)
			return questions, err
		}

		question.ID = docSnapshot.Ref.ID

		questions = append(questions, question)
	}

	return questions, nil
}

func (r *repository) GetAllByQuizId(quizId string, collectionName string) ([]Question, error) {
	var questions []Question

	// Create a query to get documents where "quiz_id" matches the provided quizId
	query := r.db.Collection(collectionName).Where("quizId", "==", quizId)

	// Retrieve the documents that match the query
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		log.Printf("Failed to get questions by quiz ID: %v", err)
		return questions, err
	}

	// Iterate through the documents and unmarshal them into Question structs
	for _, doc := range docs {
		var question Question
		if err := doc.DataTo(&question); err != nil {
			log.Printf("Failed to unmarshal question document: %v", err)
			return questions, err
		}

		question.ID = doc.Ref.ID
		questions = append(questions, question)
	}

	return questions, nil
}

func (r *repository) GetById(id string, collectionName string) (Question, error) {
	var question Question

	// Create a reference to the document using the question's ID
	docRef := r.db.Collection(collectionName).Doc(id)

	// Get the document snapshot
	docSnapshot, err := docRef.Get(context.Background())
	if err != nil {
		log.Printf("Failed to get question by ID: %v", err)
		return question, err
	}

	// Unmarshal the document snapshot into the Question struct
	if err := docSnapshot.DataTo(&question); err != nil {
		log.Printf("Failed to unmarshal question document: %v", err)
		return question, err
	}

	question.ID = docRef.ID

	return question, nil
}
