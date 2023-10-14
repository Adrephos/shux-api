package infrastructure

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
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

func (t *FirestoreServerRepository) GetRanking(ServerId string) ([]map[string]interface{}, error) {
	client := t.Client
	ctx := context.Background()
	usersRef := client.Collection("servers").Doc(ServerId).Collection("users")
	rankRef := usersRef.OrderBy("points", firestore.Desc)
	var userArr []map[string]interface{}

	iter := rankRef.Documents(ctx)
	pos := 1

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		docMap := doc.Data()
		userRank := make(map[string]interface{})

		userRank["points"] = docMap["points"]
		userRank["rank"] = pos
		userRank["id"] = doc.Ref.ID

		userArr = append(userArr, userRank)

		pos++
	}

	return userArr, nil
}

func NewFirestoreServerRepo(client *firestore.Client) *FirestoreServerRepository {
	return &FirestoreServerRepository{Client: client}
}
