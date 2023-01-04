package application

import (
	"github.com/shuxbot/shux-api/domain"
)

type ChannelApp struct{
	ChannelRepo domain.ChannelRepository
}

func (app *ChannelApp) List(ServerId string) ([]domain.Channel, error){
	ChannelArr, err := app.ChannelRepo.List(ServerId)

	return ChannelArr, err
}

func NewChannelApp(channelRepo domain.ChannelRepository) *ChannelApp {
	return &ChannelApp{ChannelRepo: channelRepo}
}
