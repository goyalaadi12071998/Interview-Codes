package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	ID             int
	PhoneNumber    string
	email          string
	linkedId       int
	linkPrecedence string
	createdAt      int
	updatedAt      int
	deletedAt      int
}
