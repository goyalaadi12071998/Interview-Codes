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

func (c core) GetContact(ctx context.Context, filter map[string]interface{}) ([]models.Contact, error) {
	var contacts []models.Contact

	data, err := c.repo.Get(contacts, filter)
	if err != nil {
		return nil, err
	}

	contacts = data.([]models.Contact)
	return contacts, nil
}

func (c core) CreateContact(ctx context.Context, data *models.Contact) (*models.Contact, error) {
	var contact *models.Contact
	contact = data

	err := c.repo.Create(contact)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (c core) UpdateContact(ctx context.Context, data *models.Contact, filter map[string]interface{}) (*models.Contact, error) {
	var contact *models.Contact
	contact = data

	err := c.repo.Update(contact, filter, contact.Id)
	if err != nil {
		return nil, err
	}

	return contact, nil
}
