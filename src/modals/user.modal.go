package modals

import (
	"time"
)

type DateTime struct {
	createdAt time.Time
	updatedAt time.Time
}

type User struct {
	id          string `gorm:"primaryKey"`
	username    string
	password    string
	firstName   string
	lastName    *string
	dateOfBirth string
	DateTime
}
