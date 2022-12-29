package domain

type (
	RoleRepository interface {
		ListRoles() ([]Role, error)
		GetRoleById(id string) error
		DeleteRole(id string) error
		UpdateRole(r *Role) error
		ReplaceRole(r *Role) error
		CreateRole(r *Role) error
	}
	Role struct {
		RoleId		string		`json:"role_id,omitempty" firestore:"-"`
		Name		string		`json:"name,omitempty" firestore:"name,omitempty"`
		Flags		int64		`json:"flags,omitempty" firestore:"flags,omitempty"`
	}
)
