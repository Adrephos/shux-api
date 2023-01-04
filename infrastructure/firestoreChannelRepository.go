package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	firestore "cloud.google.com/go/firestore"
	"github.com/shuxbot/shux-api/domain"
	"github.com/shuxbot/shux-api/infrastructure/persistance"
	"google.golang.org/api/iterator"
)

type FirestoreChannelRepository struct {
	Client *firestore.Client
}

func (t *FirestoreChannelRepository) List(ServerId string) ([]domain.Channel, error) {
	ctx := context.Background()
	var channelArr []domain.Channel
	path := fmt.Sprintf("servers/%s/channels", ServerId)

	collRef := t.Client.Collection(path)

	iter := collRef.Documents(ctx)

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		docPath := doc.Ref.Path
		parts := strings.Split(docPath, "/")

		// Get the portion of the path after the "servers" collection
		subpath := strings.Join(parts[5:], "/")


		channelMap, err := persistance.Get(subpath)
		jsonChannel, err := json.Marshal(channelMap)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		var c domain.Channel
		err = json.Unmarshal(jsonChannel, &c)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		channelArr = append(channelArr, c)
	}

	return channelArr, nil

}

func (t *FirestoreChannelRepository) Get(ChannelId string, ServerId string) (domain.Channel, error) {
	var c domain.Channel
	return c, nil
}

func (t *FirestoreChannelRepository) Delete(ChannelId string, ServerId string) error {
	return nil
}

func (t *FirestoreChannelRepository) Update(c *domain.Channel, ServerId string) error {
	return nil
}

func (t *FirestoreChannelRepository) Replace(c *domain.Channel, ServerId string) error {
	return nil
}

func (t *FirestoreChannelRepository) Create(c *domain.Channel, ServerId string) error {
	return nil
}

func NewFirestoreChannelRepo(client *firestore.Client) *FirestoreChannelRepository {
	return &FirestoreChannelRepository{Client: client}
}
