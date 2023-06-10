package db

import "gorm.io/gorm"

var repo Repo

type Db *gorm.DB

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IRepo {
	repo = Repo{
		db: db,
	}
	return repo
}

func (d Repo) Create(model any) error {
	err := d.db.Create(model).Error
	if err != nil {
		return nil
	}

	return err
}

func (d Repo) Get(model any, id int) error {
	err := d.db.Where("id = ?", id).First(model).Error
	if err != nil {
		return err
	}
	return err
}

func (d Repo) Update(model any) error {
	err := d.db.Save(model).Error
	if err != nil {
		return err
	}

	return nil
}
