package contacts

import "interview/app/providers/db"

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
