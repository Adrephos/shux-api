package application

import (
	"github.com/shuxbot/shux-api/domain"
)

type UserApp struct{
	UserRepo domain.UserRepository
}

func (app *UserApp) Get(UserId string, ServerId string) (domain.User, error){
	user, err := app.UserRepo.Get(UserId, ServerId)

	return user, err
}

func (app *UserApp) Delete(UserId string, ServerId string) error {
	err := app.UserRepo.Delete(UserId, ServerId)

	return err
}

func (app *UserApp) Update(u *domain.User, ServerId string) error{
	err := app.UserRepo.Update(u, ServerId)

	return err
}

func (app *UserApp) Create(u *domain.User, ServerId string) error{
	err := app.UserRepo.Create(u, ServerId)

	return err
}



func NewUserApp(userRepo domain.UserRepository) *UserApp {
	return &UserApp{UserRepo: userRepo}
}
