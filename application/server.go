package application

import (
	"github.com/shuxbot/shux-api/domain"
)

type ServerApp struct {
	ServerRepo domain.ServerRepository
}

func (app *ServerApp) List() ([]string, error) {
	idArr, err := app.ServerRepo.List()

	return idArr, err
}

func NewServerApp(serverRepo domain.ServerRepository) *ServerApp {
	return &ServerApp{ServerRepo: serverRepo}
}
