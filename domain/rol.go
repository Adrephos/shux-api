package domain

type (
	RolRepository interface {
		ListRoles() ([]Rol, error)
		GetRoleById(id string) error
		DeleteRol(id string) error
		UpdateRol(r *Rol) error
		ReplaceRol(r *Rol) error
	}
	Rol struct {
		RolId		string		`json:"id"`
		Name		string		`json:"name"`
		Flags		int64		`json:"flags"`
	}
)
