package models

type Contact struct {
	Id             int `gorm:"primary_key;auto_increment"`
	PhoneNumber    string
	Email          string
	LinkedId       int
	LinkPreference string
	CreatedAt      int `gorm:"autoCreateTime:milli"`
	UpdatedAt      int `gorm:"autoUpdateTime:milli"`
	DeletedAt      int
}
