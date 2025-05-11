package models

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	Notes       string    `json:"notes"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
