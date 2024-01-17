package otp

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Repository interface {
	Save(otp Otp, collectionName string) (Otp, error)
	FindByOtpCode(otpCode string, collectionName string) (Otp, error)
	Update(otp Otp, collectionName string) (Otp, error)
}

type repository struct {
	db *firestore.Client
}

func NewRepository(db *firestore.Client) *repository {
	return &repository{db}
}

func (r *repository) Save(otp Otp, collectionName string) (Otp, error) {
	// Use the Firestore client to add a new document to the collection
	doc, _, err := r.db.Collection(collectionName).Add(context.Background(), otp)
	if err != nil {
		log.Printf("Failed to create OTP: %v", err)
		return otp, err
	}

	otp.ID = doc.ID

	return otp, nil
}

func (r *repository) FindByOtpCode(otpCode string, collectionName string) (Otp, error) {
	var otp Otp
	// Assuming you have a collection named "otps"
	iter := r.db.Collection(collectionName).Where("otpCode", "==", otpCode).Limit(1).Documents(context.Background())

	// Iterate over the result set
	for {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return otp, err
		}

		// Populate otp struct with retrieved data and ID
		if err := docSnap.DataTo(&otp); err != nil {
			return otp, err
		}

		// Set the ID field in the otp struct
		otp.ID = docSnap.Ref.ID
	}

	return otp, nil
}

func (r *repository) Update(otp Otp, collectionName string) (Otp, error) {
	// Create a document reference with the specified ID
	docRef := r.db.Collection(collectionName).Doc(otp.ID)

	// Use Set with Merge option to update the existing document
	_, err := docRef.Set(context.Background(), otp, firestore.MergeAll)
	if err != nil {
		log.Printf("Failed to update OTP: %v", err)
		return otp, err
	}

	otp.ID = docRef.ID

	return otp, nil
}
