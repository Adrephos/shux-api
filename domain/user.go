package domain

import (
	"time"
)

type (
	UserRepository interface {
		GetById(UserId string, ServerId string) (User, error)
		DeleteUser(UserId string, ServerId string) error
		UpdateUser(u *User, ServerId string) error
		ReplaceUser(u *User, ServerId string) error
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
