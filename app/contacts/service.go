package contacts

var Service service

type service struct {
	core ICore
}

func NewService(core ICore) IService {
	Service = service{
		core: core,
	}
	return Service
}
