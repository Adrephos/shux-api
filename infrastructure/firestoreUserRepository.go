package infrastructure

import (
	"fmt"
	"encoding/json"
	firestore "cloud.google.com/go/firestore"
	"github.com/shuxbot/shux-api/domain"
	"github.com/shuxbot/shux-api/infrastructure/persistance"
)

type FirestoreUserRepository struct {
	Client *firestore.Client
}

func (t *FirestoreUserRepository) GetById(UserId string, ServerId string) (domain.User, error){
	path := fmt.Sprintf("servers/%s/users/%s", ServerId, UserId)
	userMap, err := persistance.Get(path)

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

func (t *FirestoreUserRepository) DeleteUser(UserId string, ServerId string) error {
	path := fmt.Sprintf("servers/%s/users", ServerId)
	err := persistance.Delete(path, UserId)
	if err != nil{
		return err
	}
	return nil
}
func (t *FirestoreUserRepository) UpdateUser(u *domain.User, ServerId string) error {
	path := fmt.Sprintf("servers/%s/users", ServerId)
	err := persistance.Update(path, *u, *&u.UserId)
	return err

}
//func (t *FirestoreUserRepository) GetById(UserId string, ServerId string) (User, error){}



