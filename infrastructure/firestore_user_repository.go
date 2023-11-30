package infrastructure

import (
	firestore "cloud.google.com/go/firestore"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/shuxbot/shux-api/domain"
	"github.com/shuxbot/shux-api/infrastructure/persistance"
)

type FirestoreUserRepository struct {
	Client *firestore.Client
}

func (t *FirestoreUserRepository) Get(UserId string, ServerId string) (domain.User, error) {
	path := fmt.Sprintf("servers/%s/users/%s", ServerId, UserId)
	userMap, err := persistance.Get(path)
	if err != nil {
		return domain.User{}, err
	}

	jsonUser, err := json.Marshal(userMap)
	if err != nil {
		fmt.Println(err)
		return domain.User{}, err
	}

	var u domain.User
	err = json.Unmarshal(jsonUser, &u)
	if err != nil {
		fmt.Println(err)
		return domain.User{}, err
	}

	return u, err
}

func (t *FirestoreUserRepository) Delete(UserId string, ServerId string) error {
	path := fmt.Sprintf("servers/%s/users", ServerId)
	err := persistance.Delete(path, UserId)

	return err
}

func (t *FirestoreUserRepository) Update(u *domain.User, ServerId string) error {
	path := fmt.Sprintf("servers/%s/users", ServerId)
	id := u.UserId
	u.UserId = ""
	err := persistance.Update(path, *u, id)

	return err
}

func (t *FirestoreUserRepository) Replace(u *domain.User, ServerId string) error {
	path := fmt.Sprintf("servers/%s/users", ServerId)
	err := persistance.Create(path, *u, u.UserId)

	return err

}

func (t *FirestoreUserRepository) Create(u *domain.User, ServerId string) error {
	path := fmt.Sprintf("servers/%s/users", ServerId)
	err := persistance.Create(path, *u, u.UserId)

	return err
}

// Instantiates a new user repository
func NewFirestoreUserRepo(client *firestore.Client) *FirestoreUserRepository {
	return &FirestoreUserRepository{Client: client}
}
