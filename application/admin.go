package application

import (
	"github.com/shuxbot/shux-api/domain"
)

type AdminApp struct {
	AdminRepo domain.AdminRepository
}

func (app *AdminApp) Register(username string, password string) error {
	err := app.AdminRepo.Register(username, password)

	return err
}

func (app *AdminApp) Login(username string) (domain.Admin, error) {
	admin, err := app.AdminRepo.Login(username)

	return admin, err
}

func NewAdminApp(adminRepository domain.AdminRepository) *AdminApp {
	return &AdminApp{
		AdminRepo: adminRepository,
	}
}
