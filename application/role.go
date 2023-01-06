package application

import (
	"github.com/shuxbot/shux-api/domain"
)

type RoleApp struct {
	RoleRepo domain.RoleRepository
}

func (app *RoleApp) List(ServerId string) ([]domain.Role, error) {
	RoleArr, err := app.RoleRepo.List(ServerId)

	return RoleArr, err
}

func (app *RoleApp) Get(RoleId string, ServerId string) (domain.Role, error) {
	role, err := app.RoleRepo.Get(RoleId, ServerId)

	return role, err
}

func (app *RoleApp) Delete(RoleId string, ServerId string) error {
	err := app.RoleRepo.Delete(RoleId, ServerId)

	return err
}

func (app *RoleApp) Update(r *domain.Role, ServerId string) error{
	err := app.RoleRepo.Update(r, ServerId)

	return err
}

func (app *RoleApp) Replace(r *domain.Role, ServerId string) error{
	err := app.RoleRepo.Replace(r, ServerId)

	return err
}

func (app *RoleApp) Create(r *domain.Role, ServerId string) error{
	err := app.RoleRepo.Create(r, ServerId)

	return err
}

func NewRoleApp(roleRepo domain.RoleRepository) *RoleApp {
	return &RoleApp{RoleRepo: roleRepo}
}
