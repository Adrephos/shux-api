package domain

type (
	ChannelRepository interface {
		ListChannels() ([]Channel, error)
		GetChannelById(id string) (Channel, error)
		DeleteChannel(id string) (Channel, error)
		UpdateChannel(c *Channel) error
		ReplaceChannel(c *Channel) error
	}
	Channel struct {
		ChannelId	   string		`json:"id"`
		Flags		   int64		`json:"flags"`
	}
)
