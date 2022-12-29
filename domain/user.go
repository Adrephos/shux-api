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
		CreateUser(u *User, ServerId string) error
	}
	Warning struct{
		Reason		   string		`json:"reason"`
		Date		   time.Time	`json:"date"`
	}
	User struct {
		UserId         string		`json:"user_id,omitempty" firestore:"-"`
		Description    string		`json:"description,omitempty" firestore:"description,omitempty"`
		Points         int64		`json:"points,omitempty" firestore:"points,omitempty"`
		Warnings       int64		`json:"warnings,omitempty" firestore:"warnings,omitempty"`
		WarningsRecord []Warning	`json:"warnings_record,omitempty" firestore:"warnings_record,omitempty"`
		Tickets        int64		`json:"tickets,omitempty" firestore:"tickets,omitempty"`
		Flags		   int64		`json:"flags,omitempty" firestore:"flags,omitempty"`
	}
)
