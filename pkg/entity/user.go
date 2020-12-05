package entity

import (
	"time"
)

type User struct {
	Id          string    `json:"id" bson:"_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

func (e *User) IsDateOfBirth(date time.Time) bool {
	if date.Month() == e.DateOfBirth.Month() && date.Day() == e.DateOfBirth.Day() {
		return true
	}
	return false
}