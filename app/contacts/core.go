package contacts

import (
	"context"
	"interview/app/models"
	"interview/app/providers/db"
)

var Core core

type core struct {
	repo db.IRepo
}

func NewCore(repo db.IRepo) ICore {
	Core = core{
		repo: repo,
	}

	return Core
}

func (c core) GetContact(ctx context.Context, filter map[string]interface{}) (*models.Contact, error) {
	var contact []*models.Contact
	err := c.repo.Get(contact, filter)
	if err != nil {
		return nil, err
	}

	if len(contact) == 0 {
		return nil, nil
	}

	return contact[0], nil
}
