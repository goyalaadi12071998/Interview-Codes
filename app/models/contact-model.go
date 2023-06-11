package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Id             int
	PhoneNumber    string
	Email          string
	LinkedId       int
	LinkPreference string
	CreatedAt      int
	UpdatedAt      int
	DeletedAt      int
}
