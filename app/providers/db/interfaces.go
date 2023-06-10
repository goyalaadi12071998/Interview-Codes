package db

type IRepo interface {
	Create(model any) error
	Get(model any, id int) error
	Update(model any) error
}
