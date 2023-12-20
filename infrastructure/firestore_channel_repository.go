package infrastructure

import (
	"fmt"
	"github.com/goccy/go-json"
	firestore "cloud.google.com/go/firestore"
	"github.com/shuxbot/shux-api/domain"
	"github.com/shuxbot/shux-api/infrastructure/persistance"
)

type FirestoreChannelRepository struct {
	Client *firestore.Client
}

func (t *FirestoreChannelRepository) List(ServerId string) ([]domain.Channel, error) {
	channelArr := make([]domain.Channel, 0)
	path := fmt.Sprintf("servers/%s/channels", ServerId)

	channelMapArr, err := persistance.List(path)

	if err != nil{
		return channelArr, err
	}

	for _, channelMap := range channelMapArr {
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
	path := fmt.Sprintf("servers/%s/channels/%s", ServerId, ChannelId)
	channelMap, err := persistance.Get(path)
	if err != nil {
		return domain.Channel{}, err
	}

	jsonChannel, err := json.Marshal(channelMap)
    if err != nil {
        fmt.Println(err)
		return domain.Channel{}, err
    }

	var c domain.Channel
	err = json.Unmarshal(jsonChannel, &c)
	if err != nil {
        fmt.Println(err)
		return domain.Channel{}, err
    }

	return c, err
}

func (t *FirestoreChannelRepository) Delete(ChannelId string, ServerId string) error {
	path := fmt.Sprintf("servers/%s/channels", ServerId)
	err := persistance.Delete(path, ChannelId)

	return err
}

func (t *FirestoreChannelRepository) Update(c *domain.Channel, ServerId string) error {
	path := fmt.Sprintf("servers/%s/channels", ServerId)
	id := c.ChannelId
	c.ChannelId = ""
	err := persistance.Update(path, *c, id)


	return err
}

func (t *FirestoreChannelRepository) Replace(c *domain.Channel, ServerId string) error {
	path := fmt.Sprintf("servers/%s/channels", ServerId)
	err := persistance.Create(path, *c, c.ChannelId)
	
	return err
}

func (t *FirestoreChannelRepository) Create(c *domain.Channel, ServerId string) error {
	path := fmt.Sprintf("servers/%s/channels", ServerId)
	err := persistance.Create(path, *c, c.ChannelId)

	return err
}

func NewFirestoreChannelRepo(client *firestore.Client) *FirestoreChannelRepository {
	return &FirestoreChannelRepository{Client: client}
}
