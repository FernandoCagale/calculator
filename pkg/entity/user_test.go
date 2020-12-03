package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_IsDateOfBirth_OK(t *testing.T) {
	user := &User{
		Id:          "001",
		FirstName:   "Rami",
		LastName:    "Marriott",
		DateOfBirth: time.Date(1990, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
	}

	assert.Equal(t, true, user.IsDateOfBirth(time.Now()))
}

func Test_IsDateOfBirth_NOK(t *testing.T) {
	user := &User{
		Id:          "001",
		FirstName:   "Rami",
		LastName:    "Marriott",
		DateOfBirth: time.Date(1990, 12,14, 0, 0, 0, 0, time.UTC),
	}

	assert.Equal(t, false, user.IsDateOfBirth(time.Now()))
}