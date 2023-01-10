package infrastructure

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/shuxbot/shux-api/domain"
	"google.golang.org/api/iterator"
)

type FirestoreServerRepository struct {
	Client *firestore.Client
}

func (t *FirestoreServerRepository) List() ([]string, error) {
	client := t.Client
	ctx := context.Background()
	var idArr []string

	collRef := client.Collection("servers")

	iter := collRef.Documents(ctx)

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		idArr = append(idArr, doc.Ref.ID)
	}

	return idArr, nil

}

func (t *FirestoreServerRepository) GetRanking(ServerId string) ([]domain.User, error){
	var userArr []domain.User
	return userArr, nil
}

func NewFirestoreServerRepo(client *firestore.Client) *FirestoreServerRepository {
	return &FirestoreServerRepository{Client: client}
}
