package application

import (
	"errors"

	"github.com/shuxbot/shux-api/domain"
)

type ServerApp struct {
	ServerRepo domain.ServerRepository
}

func (app *ServerApp) List() ([]string, error) {
	idArr, err := app.ServerRepo.List()

	return idArr, err
}

func (app *ServerApp) GetLeaderboard(ServerId string) ([]map[string]interface{}, error) {
	serverRanking, err := app.ServerRepo.GetRanking(ServerId)
	if err != nil {
		return nil, err
	}
	if len(serverRanking) < 5 {
		return nil, errors.New("Not enough users")
	}
	return serverRanking[:5], err
}

func (app *ServerApp) GetUserRank(ServerId string, UserId string) (map[string]interface{}, error) {
	serverRanking, err := app.ServerRepo.GetRanking(ServerId)
	for _, item := range serverRanking {
		if item["id"] == UserId {
			return item, err
		}
	}
	return nil, errors.New("User not found")
}

func (app *ServerApp) GetTickets(ServerId string) (map[string]interface{}, error) {
	tickets, err := app.ServerRepo.GetTickets(ServerId)

	return tickets, err
}

func (app *ServerApp) EditTickets(ServerId string, tickets map[string]interface{}) error {
	err := app.ServerRepo.EditTickets(ServerId, tickets)

	return err
}

func NewServerApp(serverRepo domain.ServerRepository) *ServerApp {
	return &ServerApp{ServerRepo: serverRepo}
}
