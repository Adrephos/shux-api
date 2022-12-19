package domain

type (
	RolRepository interface {
		ListRoles() ([]Rol, error)
		FindRoles(string) (Rol, error)
		DeleteRol(string) (Rol, error)
		UpdateRol(Rol) (Rol, error)
		ReplaceRol(Rol) (Rol, error)
	}
	Rol struct {
		RolId		string		`json:"id"`
		Name		string		`json:"name"`
		Flags		int64		`json:"flags"`
	}
)
