package quizCategory

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Repository interface {
	Save(quizCategory QuizCategory, collectionName string) (QuizCategory, error)
	Update(quizCategory QuizCategory, collectionName string) (QuizCategory, error)
	Delete(ID string, collectionName string) (QuizCategory, error)
	GetById(ID string, collectionName string) (QuizCategory, error)
	GetAll(collectionName string) ([]QuizCategory, error)
}

type repository struct {
	db *firestore.Client
}

func NewRepository(db *firestore.Client) *repository {
	return &repository{db}
}

func (r *repository) Save(quizCategory QuizCategory, collectionName string) (QuizCategory, error) {
	// Use the Firestore client to add a new document to the collection
	doc, _, err := r.db.Collection(collectionName).Add(context.Background(), quizCategory)
	if err != nil {
		log.Printf("Failed to create quiz category: %v", err)
		return quizCategory, err
	}

	quizCategory.ID = doc.ID
	return quizCategory, nil
}

func (r *repository) Update(quizCategory QuizCategory, collectionName string) (QuizCategory, error) {
	// Create a reference to the document using the quiz category's ID
	docRef := r.db.Collection(collectionName).Doc(quizCategory.ID)

	// Update the document with the new data
	_, err := docRef.Set(context.Background(), quizCategory)
	if err != nil {
		log.Printf("Failed to update quiz category: %v", err)
		return quizCategory, err
	}

	quizCategory.ID = docRef.ID

	return quizCategory, nil
}

func (r *repository) Delete(ID string, collectionName string) (QuizCategory, error) {
	var quizCategory QuizCategory

	// Create a reference to the document using the quiz category's ID
	docRef := r.db.Collection(collectionName).Doc(ID)

	// Get the document snapshot to retrieve the existing data before deletion
	docSnapshot, err := docRef.Get(context.Background())
	if err != nil {
		log.Printf("Failed to get quiz category before deletion: %v", err)
		return quizCategory, err
	}

	// Unmarshal the document snapshot into the QuizCategory struct
	if err := docSnapshot.DataTo(&quizCategory); err != nil {
		log.Printf("Failed to unmarshal quiz category document: %v", err)
		return quizCategory, err
	}

	// Delete the document
	_, err = docRef.Delete(context.Background())
	if err != nil {
		log.Printf("Failed to delete quiz category: %v", err)
		return quizCategory, err
	}

	quizCategory.ID = docRef.ID

	return quizCategory, nil
}

func (r *repository) GetById(ID string, collectionName string) (QuizCategory, error) {
	var quizCategory QuizCategory

	// Create a reference to the document using the quiz category's ID
	docRef := r.db.Collection(collectionName).Doc(ID)

	// Get the document snapshot
	docSnapshot, err := docRef.Get(context.Background())
	if err != nil {
		log.Printf("Failed to get quiz category by ID: %v", err)
		return quizCategory, err
	}

	// Unmarshal the document snapshot into the QuizCategory struct
	if err := docSnapshot.DataTo(&quizCategory); err != nil {
		log.Printf("Failed to unmarshal quiz category document: %v", err)
		return quizCategory, err
	}

	quizCategory.ID = docRef.ID

	return quizCategory, nil
}

func (r *repository) GetAll(collectionName string) ([]QuizCategory, error) {
	var quizCategories []QuizCategory

	// Create a query to get all documents from the collection
	iter := r.db.Collection(collectionName).Documents(context.Background())
	defer iter.Stop()

	// Iterate through the documents and unmarshal them into QuizCategory structs
	for {
		var quizCategory QuizCategory
		docSnapshot, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Printf("Failed to fetch quiz category: %v", err)
			return quizCategories, err
		}

		if err := docSnapshot.DataTo(&quizCategory); err != nil {
			log.Printf("Failed to unmarshal quiz category document: %v", err)
			return quizCategories, err
		}

		quizCategory.ID = docSnapshot.Ref.ID

		quizCategories = append(quizCategories, quizCategory)
	}

	return quizCategories, nil
}
