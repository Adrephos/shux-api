package domain

type (
	RoleRepository interface {
	List(ServerId string) ([]Role, error)
		Get(RoleId string, ServerId string) (Role, error)
		Delete(RoleId string, ServerId string) error
		Update(r *Role, ServerId string) error
		Replace(r *Role, ServerId string) error
		Create(r *Role, ServerId string) error
	}
	Role struct {
		RoleId		string		`json:"role_id,omitempty" firestore:"-"`
		Name		string		`json:"name,omitempty" firestore:"name"`
		Flags		int64		`json:"flags,omitempty" firestore:"flags"`
	}
)
