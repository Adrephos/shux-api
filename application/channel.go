package application

import (
	"github.com/shuxbot/shux-api/domain"
)

type ChannelApp struct {
	ChannelRepo domain.ChannelRepository
}

func (app *ChannelApp) List(ServerId string) ([]domain.Channel, error) {
	ChannelArr, err := app.ChannelRepo.List(ServerId)

	return ChannelArr, err
}

func (app *ChannelApp) Get(ChannelId string, ServerId string) (domain.Channel, error) {
	channel, err := app.ChannelRepo.Get(ChannelId, ServerId)

	return channel, err
}

func (app *ChannelApp) Delete(ChannelId string, ServerId string) error {
	err := app.ChannelRepo.Delete(ChannelId, ServerId)

	return err
}

func (app *ChannelApp) Update(c *domain.Channel, ServerId string) error{
	err := app.ChannelRepo.Update(c, ServerId)

	return err
}

func (app *ChannelApp) Replace(c *domain.Channel, ServerId string) error{
	err := app.ChannelRepo.Replace(c, ServerId)

	return err
}

func (app *ChannelApp) Create(c *domain.Channel, ServerId string) error{
	err := app.ChannelRepo.Create(c, ServerId)

	return err
}

func NewChannelApp(channelRepo domain.ChannelRepository) *ChannelApp {
	return &ChannelApp{ChannelRepo: channelRepo}
}
