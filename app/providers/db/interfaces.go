package db

type IRepo interface {
	Create(model any) error

	Get(model any, filter map[string]interface{}) error

	FindOne(model any, filter map[string]interface{}) error

	Update(model any) error
}
