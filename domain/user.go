package domain

import (
	"time"
)

type (
	UserRepository interface {
		FindUser(string) (User, error)
		DeleteUser(string) (User, error)
		UpdateUser(User) (User, error)
		ReplaceUser(User) (User, error)
	}
	Warning struct{
		Reason		   string		`json:"reason"`
		Date		   time.Time	`json:"date"`
	}
	User struct {
		UserId         string		`json:"id"`
		Description    string		`json:"description"`
		Points         int64		`json:"points"`
		ShowLevel      bool			`json:"show_level"`
		Warnings       int64		`json:"warnings"`
		WarningsRecord []Warning	`json:"warnings_record"`
		Beta           bool			`json:"beta"`
		Tickets        int64		`json:"tickets"`
	}
)
