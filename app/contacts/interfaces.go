package contacts

import (
	"context"
	errorclass "interview/app/error"
	"interview/app/models"
	"interview/app/structs"
)

type ICore interface {
	GetContact(ctx context.Context, filter map[string]interface{}) (*models.Contact, error)
	CreateContact(ctx context.Context, data *models.Contact) (*models.Contact, error)
}

type IService interface {
	CreateContact(ctx context.Context, data *structs.RequestIdentify) (*structs.ResponseIdentify, *errorclass.Error)
}
