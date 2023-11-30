package domain

type (
	AdminRepository interface {
		Register(username string, password string) error
		Login(username string) (Admin, error)
	}
	Admin struct {
		AdminId  string `json:"id,omitempty" firestore:"-"`
		Username string `json:"username,omitempty" firestore:"username"`
		Password string `json:"password,omitempty" firestore:"password"`
	}
)
