package boot

import (
	"context"
	"interview/app/common"
	"interview/app/router"
)

func initProviders() error {
	return nil
}

func initRoutes() error {
	configs := common.GetConfig()
	err := router.InitializeRouter(router.CoreConfigs{
		Name: configs.Core.Name,
		Host: configs.Core.Host,
		Port: configs.Core.Port,
	})
	return err
}

func Init(ctx context.Context, env string) error {
	err := common.InitConfig(env)
	if err != nil {
		return err
	}

	err = initProviders()
	if err != nil {
		return err
	}

	err = initRoutes()
	if err != nil {
		return err
	}

	return nil
}
