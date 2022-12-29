package domain

type (
	ChannelRepository interface {
		ListChannels() ([]Channel, error)
		GetChannelById(id string) (Channel, error)
		DeleteChannel(id string) (Channel, error)
		UpdateChannel(c *Channel) error
		ReplaceChannel(c *Channel) error
		CreateChannel(c *Channel) error
	}
	Channel struct {
		ChannelId	   string		`json:"channel_id,omitempty" firestore:"-"`
		Flags		   int64		`json:"flags,omitempty" firestore:"flags,omitempty"`
	}
)
