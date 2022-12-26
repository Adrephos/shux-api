package domain

type (
	RoleRepository interface {
		ListRoles() ([]Role, error)
		GetRoleById(id string) error
		DeleteRole(id string) error
		UpdateRole(r *Role) error
		ReplaceRole(r *Role) error
	}
	Role struct {
		RoleId		string		`json:"role_id" firestore:"-"`
		Name		string		`json:"name" firestore:"name"`
		Flags		int64		`json:"flags" firestore:"flags"`
	}
)
