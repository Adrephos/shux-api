package domain

import (
	"time"
)

type (
	UserRepository interface {
		Get(UserId string, ServerId string) (User, error)
		Delete(UserId string, ServerId string) error
		Update(u *User, ServerId string) error
		Replace(u *User, ServerId string) error
		Create(u *User, ServerId string) error
	}
	Warning struct {
		Reason string    `json:"reason"`
		Date   time.Time `json:"date"`
	}
	User struct {
		UserId         string    `json:"id,omitempty" firestore:"-"`
		Description    string    `json:"description,omitempty" firestore:"description,omitempty"`
		Points         int64     `json:"points,omitempty" firestore:"points,omitempty"`
		Warnings       int64     `json:"warnings,omitempty" firestore:"warnings,omitempty"`
		WarningsRecord []Warning `json:"warnings_record,omitempty" firestore:"warnings_record,omitempty"`
		Tickets        int64     `json:"tickets,omitempty" firestore:"tickets,omitempty"`
		Flags          int64     `json:"flags,omitempty" firestore:"flags,omitempty"`
	}
)
