package domain

type (
	ChannelRepository interface {
		List() ([]Channel, error)
		Get(id string) (Channel, error)
		Delete(id string) (Channel, error)
		Update(c *Channel) error
		Replace(c *Channel) error
		Create(c *Channel) error
	}
	Channel struct {
		ChannelId	   string		`json:"channel_id,omitempty" firestore:"-"`
		Flags		   int64		`json:"flags,omitempty" firestore:"flags,omitempty"`
	}
)
