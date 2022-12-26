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
		UserId         string		`json:"user_id" firestore:"-"`
		Description    string		`json:"description" firestore:"description"`
		Points         int64		`json:"points" firestore:"points"`
		Warnings       int64		`json:"warnings" firestore:"warnings"`
		WarningsRecord []Warning	`json:"warnings_record" firestore:"warnings_record"`
		Tickets        int64		`json:"tickets" firestore:"tickets"`
		Flags		   int64		`json:"flags" firestore:"flags"`
	}
)
