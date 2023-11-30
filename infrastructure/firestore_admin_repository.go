package infrastructure

import (
	"context"
	"encoding/json"
	"errors"

	firestore "cloud.google.com/go/firestore"
	"github.com/shuxbot/shux-api/domain"
)

type FirestoreAdminRepository struct {
	Client *firestore.Client
}

func (t *FirestoreAdminRepository) Register(username string, password string) error {
	query := t.Client.Collection("admins").Where("username", "==", username)

	documents, err := query.Documents(context.Background()).GetAll()

	if err != nil {
		return err
	}

	if len(documents) > 0 {
		return errors.New("Username already exists")
	}

	data := map[string]interface{}{
		"username": username,
		"password": password,
	}

	_, _, err = t.Client.Collection("admins").Add(context.Background(), data)

	if err != nil {
		return err
	}

	return err
}

func (t *FirestoreAdminRepository) Login(username string) (domain.Admin, error) {
	query := t.Client.Collection("admins").Where("username", "==", username)

	documents, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		return domain.Admin{}, err
	} else if len(documents) == 0 {
		return domain.Admin{}, errors.New("Username does not exist")
	}

	adminMap := documents[0].Data()

	jsonUser, err := json.Marshal(adminMap)
	if err != nil {
		return domain.Admin{}, err
	}

	var admin domain.Admin
	err = json.Unmarshal(jsonUser, &admin)
	if err != nil {
		return domain.Admin{}, err
	}

	return admin, err
}

// Instantiates a new admin repository
func NewFirestoreAdminRepo(client *firestore.Client) *FirestoreAdminRepository {
	return &FirestoreAdminRepository{Client: client}
}
