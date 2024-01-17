package participantQuestion

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

type Repository interface {
	Save(questions []ParticipantQuestion, collectionName string) ([]ParticipantQuestion, error)
	GetByParticipantId(participantId string, collectionName string) ([]ParticipantQuestion, error)
}

type repository struct {
	db *firestore.Client
}

func NewRepository(db *firestore.Client) *repository {
	return &repository{db}
}

func (r *repository) Save(questions []ParticipantQuestion, collectionName string) ([]ParticipantQuestion, error) {
	var createdParticipantQuestions []ParticipantQuestion
	// Use a batch to perform multiple writes in a single transaction
	batch := r.db.Batch()

	// Add each question to the batch as a set operation
	for _, q := range questions {

		// Create a new document reference with the specified ID
		docRef := r.db.Collection(collectionName).NewDoc()

		// Add a set operation to the batch
		batch.Set(docRef, q)

		q.ID = docRef.ID

		createdParticipantQuestions = append(createdParticipantQuestions, q)

	}

	// Commit the batch to Firestore
	_, err := batch.Commit(context.Background())
	if err != nil {
		log.Printf("Failed to save participant questions: %v", err)
		return questions, err
	}

	// Return the saved participant questions with their IDs
	return createdParticipantQuestions, nil
}

func (r *repository) GetByParticipantId(participantId string, collectionName string) ([]ParticipantQuestion, error) {
	var participantQuestions []ParticipantQuestion

	// Create a query to get documents where "participant_id" matches the provided participantId
	query := r.db.Collection(collectionName).Where("participantId", "==", participantId)

	// Retrieve the documents that match the query
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		log.Printf("Failed to get participant questions by participant ID: %v", err)
		return participantQuestions, err
	}

	// Iterate through the documents and unmarshal them into ParticipantQuestion structs
	for _, doc := range docs {
		var participantQuestion ParticipantQuestion
		if err := doc.DataTo(&participantQuestion); err != nil {
			log.Printf("Failed to unmarshal participant question document: %v", err)
			return participantQuestions, err
		}

		participantQuestion.ID = doc.Ref.ID
		participantQuestions = append(participantQuestions, participantQuestion)
	}

	return participantQuestions, nil
}
