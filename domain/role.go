package domain

type (
	RoleRepository interface {
		List() ([]Role, error)
		Get(id string) error
		Delete(id string) error
		Update(r *Role) error
		Replace(r *Role) error
		Create(r *Role) error
	}
	Role struct {
		RoleId		string		`json:"role_id,omitempty" firestore:"-"`
		Name		string		`json:"name,omitempty" firestore:"name"`
		Flags		int64		`json:"flags,omitempty" firestore:"flags"`
	}
)
