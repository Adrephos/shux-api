package domain

type (
	ChannelRepository interface {
		ListChannels() ([]Channel, error)
		FindChannel(string) (Channel, error)
		DeleteChannel(string) (Channel, error)
		UpdateChannel(Channel) (Channel, error)
		ReplaceChannel(Channel) (Channel, error)
	}
	Channel struct {
		ChannelId	   string		`json:"id"`
		Flags		   int64		`json:"flags"`
	}
)
