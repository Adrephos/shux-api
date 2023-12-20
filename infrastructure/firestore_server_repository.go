package infrastructure

import (
	"context"
	"errors"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type FirestoreServerRepository struct {
	Client *firestore.Client
}

func (t *FirestoreServerRepository) List() ([]string, error) {
	client := t.Client
	ctx := context.Background()
	idArr := make([]string, 0)

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

func (t *FirestoreServerRepository) GetTickets(ServerId string) (map[string]interface{}, error) {
	client := t.Client
	ctx := context.Background()
	serverRef := client.Collection("servers").Doc(ServerId)
	doc, err := serverRef.Get(ctx)
	ticketsMap := map[string]interface{}{
		"tickets": make(map[string]interface{}),
	}

	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return nil, errors.New("Tickets not found")
		}
		return ticketsMap, nil
	}

	tickets, ok := doc.Data()["tickets"].(map[string]interface{})
	if !ok {
		return ticketsMap, nil
	}
	ticketsMap["tickets"] = tickets

	return ticketsMap, nil
}

func (t *FirestoreServerRepository) EditTickets(ServerId string, tickets map[string]interface{}) error {
	client := t.Client
	ctx := context.Background()
	serverRef := client.Collection("servers").Doc(ServerId)
	_, err := serverRef.Set(ctx, map[string]interface{}{
		"tickets": tickets,
	})

	if err != nil {
		return err
	}

	return nil
}

func NewFirestoreServerRepo(client *firestore.Client) *FirestoreServerRepository {
	return &FirestoreServerRepository{Client: client}
}
