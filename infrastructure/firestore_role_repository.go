package infrastructure

import (
	"fmt"
	"github.com/goccy/go-json"
	firestore "cloud.google.com/go/firestore"
	"github.com/shuxbot/shux-api/domain"
	"github.com/shuxbot/shux-api/infrastructure/persistance"
)

type FirestoreRoleRepository struct {
	Client *firestore.Client
}

func (t *FirestoreRoleRepository) List(ServerId string) ([]domain.Role, error) {
	var roleArr []domain.Role
	path := fmt.Sprintf("servers/%s/roles", ServerId)

	roleMapArr, err := persistance.List(path)

	if err != nil{
		return roleArr, err
	}

	for _, roleMap := range roleMapArr {
		jsonChannel, err := json.Marshal(roleMap)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		var r domain.Role
		err = json.Unmarshal(jsonChannel, &r)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		roleArr = append(roleArr, r)
	}

	return roleArr, nil

}

func (t *FirestoreRoleRepository) Get(RoleId string, ServerId string) (domain.Role, error) {
	path := fmt.Sprintf("servers/%s/roles/%s", ServerId, RoleId)
	roleMap, err := persistance.Get(path)
	if err != nil {
		return domain.Role{}, err
	}

	jsonChannel, err := json.Marshal(roleMap)
    if err != nil {
        fmt.Println(err)
		return domain.Role{}, err
    }

	var r domain.Role
	err = json.Unmarshal(jsonChannel, &r)
	if err != nil {
        fmt.Println(err)
		return domain.Role{}, err
    }

	return r, err
}

func (t *FirestoreRoleRepository) Delete(RoleId string, ServerId string) error {
	path := fmt.Sprintf("servers/%s/roles", ServerId)
	err := persistance.Delete(path, RoleId)

	return err
}

func (t *FirestoreRoleRepository) Update(r *domain.Role, ServerId string) error {
	path := fmt.Sprintf("servers/%s/roles", ServerId)
	id := r.RoleId
	r.RoleId = ""
	err := persistance.Update(path, *r, id)

	return err
}

func (t *FirestoreRoleRepository) Replace(r *domain.Role, ServerId string) error {
	path := fmt.Sprintf("servers/%s/roles", ServerId)
	err := persistance.Create(path, *r, r.RoleId)
	
	return err
}

func (t *FirestoreRoleRepository) Create(r *domain.Role, ServerId string) error {
	path := fmt.Sprintf("servers/%s/roles", ServerId)
	err := persistance.Create(path, *r, r.RoleId)

	return err
}

func NewFirestoreRoleRepo(client *firestore.Client) *FirestoreRoleRepository {
	return &FirestoreRoleRepository{Client: client}
}
