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
		RoleId		string		`json:"id"`
		Name		string		`json:"name"`
		Flags		int64		`json:"flags"`
	}
)
