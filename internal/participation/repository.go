package participation

import (
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Repository interface {
	Save(participation Participation, collectionName string) (Participation, error)
	GetByQuizId(quizId string, collectionName string) (Participation, error)
	GetByUserId(userId string, collectionName string) (Participation, error)
}

type repository struct {
	db *firestore.Client
}

func NewRepository(db *firestore.Client) *repository {
	return &repository{db}
}

func (r *repository) Save(participation Participation, collectionName string) (Participation, error) {
	// Use the Firestore client to add a new document to the collection
	doc, _, err := r.db.Collection(collectionName).Add(context.Background(), participation)
	if err != nil {
		log.Printf("Failed to create participation: %v", err)
		return participation, err
	}

	participation.ID = doc.ID

	return participation, nil
}

func (r *repository) GetByQuizId(quizId string, collectionName string) (Participation, error) {
	var participation Participation

	// Assuming you have a collection named "participations"
	iter := r.db.Collection(collectionName).Where("quizId", "==", quizId).Where("status", "==", "creator").Limit(1).Documents(context.Background())

	// Iterate over the result set
	for {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return participation, err
		}

		// Populate participation struct with retrieved data and ID
		if err := docSnap.DataTo(&participation); err != nil {
			return participation, err
		}

		// Set the ID field in the participation struct
		participation.ID = docSnap.Ref.ID

		// Fetch user data based on userId
		userRef := r.db.Collection("users").Doc(participation.UserId)
		userSnapshot, err := userRef.Get(context.Background())
		if err != nil {
			log.Printf("Failed to get user for participation: %v", err)
			return participation, err
		}

		// Unmarshal user data into the Participation struct
		if err := userSnapshot.DataTo(&participation.User); err != nil {
			log.Printf("Failed to unmarshal user document: %v", err)
			return participation, err
		}
	}

	return participation, nil
}

func (r *repository) GetByUserId(userId string, collectionName string) (Participation, error) {
	var participation Participation
	// Assuming you have a collection named "participations"
	iter := r.db.Collection(collectionName).Where("userId", "==", userId).Limit(1).Documents(context.Background())

	// Iterate over the result set
	for {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return participation, err
		}

		// Populate participation struct with retrieved data and ID
		if err := docSnap.DataTo(&participation); err != nil {
			return participation, err
		}

		// Set the ID field in the participation struct
		participation.ID = docSnap.Ref.ID
	}

	return participation, nil
}
