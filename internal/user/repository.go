package user

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Repository interface {
	FindByEmail(email string, collectionName string) (User, error)
	GetById(ID string, collectionName string) (User, error)
	Save(user User, collectionName string) (User, error)
	Update(user User, collectionName string) (User, error)
}

type repository struct {
	db *firestore.Client
}

func NewRepository(db *firestore.Client) *repository {
	return &repository{db}
}

func (r *repository) FindByEmail(email string, collectionName string) (User, error) {
	var user User
	// Assuming you have a collection named "users"
	iter := r.db.Collection(collectionName).Where("email", "==", email).Limit(1).Documents(context.Background())

	// Iterate over the result set
	for {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return user, err
		}

		// Populate user struct with retrieved data and ID
		if err := docSnap.DataTo(&user); err != nil {
			return user, err
		}

		// Set the ID field in the user struct
		user.ID = docSnap.Ref.ID
	}

	return user, nil
}

func (r *repository) GetById(ID string, collectionName string) (User, error) {
	var user User

	// Create a reference to the document using the user's ID
	docRef := r.db.Collection(collectionName).Doc(ID)

	// Retrieve the document snapshot
	docSnapshot, err := docRef.Get(context.Background())
	if err != nil {
		log.Printf("Failed to get user by ID: %v", err)
		return user, err
	}

	// Unmarshal the document snapshot into the User struct
	if err := docSnapshot.DataTo(&user); err != nil {
		log.Printf("Failed to unmarshal user document: %v", err)
		return user, err
	}

	user.ID = docRef.ID

	return user, nil
}

func (r *repository) Save(user User, collectionName string) (User, error) {
	doc, _, err := r.db.Collection(collectionName).Add(context.Background(), user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return user, err
	}

	user.ID = doc.ID

	return user, nil
}

func (r *repository) Update(user User, collectionName string) (User, error) {
	// Check if the user has a valid ID
	if user.ID == "" {
		return user, errors.New("ID not specified")
	}

	// Create a reference to the document using the user's ID
	docRef := r.db.Collection(collectionName).Doc(user.ID)

	// Update the document with the new data
	_, err := docRef.Set(context.Background(), user)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return user, err
	}

	user.ID = docRef.ID

	return user, nil
}
